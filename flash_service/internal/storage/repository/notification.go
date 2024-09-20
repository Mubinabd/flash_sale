package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Mubinabd/flash_sale/internal/pkg/config"
	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/Mubinabd/flash_sale/internal/pkg/help"
	"github.com/google/uuid"
)

type NotificationRepo struct {
	db *sql.DB
	cf *config.Config
}

func NewNotificationRepo(db *sql.DB, cf *config.Config) *NotificationRepo {
	return &NotificationRepo{db: db, cf: cf}
}
func (r *NotificationRepo) CreateNotification(req *pb.NotificationCreate) (*pb.Void, error) {
	tr, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	var (
		user_id    string
		user_email string
		user_name  string
	)

	if req.UserId == "" {
		query := `select id from users where role = 'admin' and deleted_at = 0 limit 1`
		row := tr.QueryRow(query)
		err := row.Scan(&user_id)
		if err == sql.ErrNoRows {
			tr.Rollback()
			return nil, errors.New("no admin found")
		} else if err != nil {
			tr.Rollback()
			return nil, err
		}
	} else {
		user_id = req.UserId
	}

	//geting the email
	user_query := `select email,username from users where id = $1 and deleted_at = 0`
	row := tr.QueryRow(user_query, req.UserId)
	err = row.Scan(&user_email, &user_name)
	if err == sql.ErrNoRows {
		tr.Rollback()
		return nil, errors.New("user not found")
	} else if err != nil {
		tr.Rollback()
		return nil, err
	}

	//sending email
	from := r.cf.Email
	password := r.cf.EmailPassword
	err = help.SendVerificationCode(help.Params{
		From:     from,
		Password: password,
		To:       user_email,
		Message:  fmt.Sprintf("Hi %s", user_name),
		Code:     req.Content,
	})

	if err != nil {
		tr.Rollback()
		return nil, errors.New("failed to send notification email" + err.Error())
	}
	query := `insert into notifications(id,
										type,
										user_id,
										content,
										status)
						values($1,$2,$3,$4,$5)`
	_, err = tr.Exec(query,
		uuid.NewString(),
		req.Type,
		user_id,
		req.Content,
		"pending")
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	err = tr.Commit()
	if err != nil {
		tr.Rollback()
		return nil, err
	}
	return &pb.Void{}, nil
}
func (r *NotificationRepo) DeleteNotification(id *pb.GetById) (*pb.Void, error) {
	query := `update notifications set deleted_at = EXTRACT(EPOCH FROM now()) 
				where id = $1 and deleted_at = 0`
	_, err := r.db.Exec(query, id.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (r *NotificationRepo) UpdateNotification(req *pb.NotificationUpdate) (*pb.Void, error) {
	query := "UPDATE notifications SET "
	var cons []string
	var args []interface{}

	// Dynamically build the query
	if req.Body.Content != "" && req.Body.Content != "string" {
		cons = append(cons, fmt.Sprintf("content=$%d", len(args)+1))
		args = append(args, req.Body.Content)
	}
	if req.Body.Status != "" && req.Body.Status != "string" {
		cons = append(cons, fmt.Sprintf("status=$%d", len(args)+1))
		args = append(args, req.Body.Status)
	}

	// Ensure there's at least one field to update
	if len(cons) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += strings.Join(cons, ", ")
	query += fmt.Sprintf(" WHERE deleted_at = 0 and id=$%d", len(args)+1)
	args = append(args, req.NotificationId)

	// Execute the query
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *NotificationRepo) GetNotifications(req *pb.NotifFilter) (*pb.NotificationList, error) {

	query := `SELECT id, 
					type, 
					content, 
					status, 
					created_at,
					user_id
		FROM notifications 
		WHERE deleted_at = 0`
	var cons []string
	var args []interface{}

	// Dynamically build the query
	if req.UserId != "" && req.UserId != "string" {
		cons = append(cons, fmt.Sprintf("user_id = $%d", len(args)+1))
		args = append(args, "%"+req.UserId+"%")
	}

	if req.Status != "" && req.Status != "string" {
		cons = append(cons, fmt.Sprintf("status = $%d", len(args)+1))
		args = append(args, "%"+req.Status+"%")
	}
	if req.Content != "" && req.Content != "string" {
		cons = append(cons, fmt.Sprintf("content = $%d", len(args)+1))
		args = append(args, "%"+req.Content+"%")
	}

	// Append conditions to query if any exist
	if len(cons) > 0 {
		query += " AND " + strings.Join(cons, " AND ")
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, req.Filter.Limit, req.Filter.Offset)

	// Execute the query
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	// Prepare the response
	var notificationList pb.NotificationList
	for rows.Next() {
		var notif pb.NotificationGet
		if err := rows.Scan(&notif.Id,
			&notif.UserId,
			&notif.Type,
			&notif.Status,
			&notif.CreatedAt,
			&notif.Content); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		notificationList.Notifications = append(notificationList.Notifications, &notif)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %w", err)
	}
	for _, n := range notificationList.Notifications {
		query := `update notifications set status = read 
					where deleted_at = 0 and id = $1`
		_, err := r.db.Exec(query, n.Id)
		if err != nil {
			return nil, err
		}
	}

	return &notificationList, nil
}
func (r *NotificationRepo) GetNotification(id *pb.GetById) (*pb.NotificationGet, error) {
	query := `select id,
					content,
					type,
					status,
					created_at,
					user_id
			from notifications where deleted_at = 0 and id = $1`
	row := r.db.QueryRow(query, id.Id)

	var notif pb.NotificationGet
	err := row.Scan(&notif.Id,
		&notif.UserId,
		&notif.Content,
		&notif.Status,
		&notif.CreatedAt,
		&notif.Type)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &notif, nil
}
