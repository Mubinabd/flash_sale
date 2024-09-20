package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/google/uuid"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) CreateOrder(req *pb.CreateOrderReq) (*pb.Void, error) {
	id := uuid.NewString()

	query := `INSERT INTO 
		orders 
		(id, 
		user_id, 
		flash_sale_id, 
		status) 
		VALUES 
		($1, $2, $3, $4)`

	_, err := r.db.Exec(query, id, req.UserID, req.FlashSaleID, req.OrderStatus)

	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (r *OrderRepo) UpdateOrder(req *pb.UpdateOrderReq) (*pb.Void, error) {
	var args []interface{}
	var conditions []string

	if req.Body.UserID != "" && req.Body.UserID != "string" {
		args = append(args, req.Body.UserID)
		conditions = append(conditions, fmt.Sprintf("user_id = $%d", len(args)))
	}

	if req.Body.FlashSaleID != "" && req.Body.FlashSaleID != "string" {
		args = append(args, req.Body.FlashSaleID)
		conditions = append(conditions, fmt.Sprintf("flash_sale_id = $%d", len(args)))
	}

	if req.Body.OrderStatus != "" && req.Body.OrderStatus != "string" {
		args = append(args, req.Body.OrderStatus)
		conditions = append(conditions, fmt.Sprintf("status = $%d", len(args)))
	}

	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	query := fmt.Sprintf("UPDATE orders SET %s WHERE id = $%d", strings.Join(conditions, ", "), len(args)+1)

	args = append(args, req.Id)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating orders", err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (r *OrderRepo) GetOrder(req *pb.GetById) (*pb.Order, error) {
	query := `
		SELECT 
			o.id,
			u.id,
			u.username,
			u.email,
			u.full_name,
			u.date_of_birth,
			f.id,
			f.name,
			f.start_time,
			f.end_time,
			f.status,
			o.status,
			o.created_at
		FROM 
			orders o
		LEFT JOIN 
			users u
		ON 
			o.user_id = u.id
		LEFT JOIN 
			flash_sales f
		ON 
			o.flash_sale_id = f.id
		WHERE 
			o.id = $1	
		AND
			o.deleted_at = 0
	`

	res := &pb.Order{
		User: &pb.UserRes{},
		FlashSaleID: &pb.FlashSale{},
	}

	err := r.db.QueryRow(query, req.Id).
		Scan(
			&res.Id,
			&res.User.Id,
			&res.User.Username,
			&res.User.Email,
			&res.User.FullName,
			&res.User.DateOfBirth,
			&res.FlashSaleID.Id,
			&res.FlashSaleID.Name,
			&res.FlashSaleID.StartTime,
			&res.FlashSaleID.EndTime,
			&res.FlashSaleID.Status,
			&res.OrderStatus,
			&res.CreatedAt,
		)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *OrderRepo) ListAllOrders(req *pb.ListAllOrdersReq) (*pb.ListAllOrdersRes, error) {

	query := `
		SELECT 
			o.id,
			u.id,
			u.username,
			u.email,
			u.full_name,
			u.date_of_birth,
			f.id,
			f.name,
			f.start_time,
			f.end_time,
			f.status,
			o.status,
			o.created_at
		FROM 
			orders o
		LEFT JOIN 
			users u
		ON 
			o.user_id = u.id
		LEFT JOIN 
			flash_sales f
		ON 
			o.flash_sale_id = f.id
		WHERE 
			o.deleted_at = 0		
		ORDER BY
			o.created_at DESC
		`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := &pb.ListAllOrdersRes{
		Orders: make([]*pb.Order, 0), 
	}
	for rows.Next() {
		var order pb.Order
		order.User = &pb.UserRes{}
		order.FlashSaleID = &pb.FlashSale{}
		err := rows.Scan(
			&order.Id,
			&order.User.Id,
			&order.User.Username,
			&order.User.Email,
			&order.User.FullName,
			&order.User.DateOfBirth,
			&order.FlashSaleID.Id,
			&order.FlashSaleID.Name,
			&order.FlashSaleID.StartTime,
			&order.FlashSaleID.EndTime,
			&order.FlashSaleID.Status,
			&order.OrderStatus,
			&order.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res.Orders = append(res.Orders, &order)
	}
	res.TotalCount = int64(len(res.Orders))
	return res, nil
}

func (r *OrderRepo) DeleteOrder(req *pb.GetById) (*pb.Void, error) {
	query := `
		UPDATE 
			orders 
		SET 
			deleted_at = extract(epoch from now()) 
		WHERE id = $1`
	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}


func (s *OrderRepo) GetOrderHistory(req *pb.OrderHistoryReq) (*pb.OrderHistoryRes, error) {
    rows, err := s.db.Query(`SELECT * FROM orders WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`,
        req.UserID, req.Pagination.Limit, req.Pagination.Offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var orders []*pb.Order
    for rows.Next() {
        var order pb.Order
        if err := rows.Scan(&order.Id, &order.User.Id, &order.FlashSaleID, &order.OrderStatus, &order.CreatedAt); err != nil {
            return nil, err
        }
        orders = append(orders, &order)
    }

    return &pb.OrderHistoryRes{Orders: orders, TotalCount: int64(len(orders))}, nil
}


func (s *OrderRepo) CancelOrder(req *pb.GetById) (*pb.CancelOrderRes, error) {
    _, err := s.db.Exec(`UPDATE orders SET status = 'canceled', updated_at = NOW() WHERE id = $1`, req.Id)
    if err != nil {
        return nil, err
    }

    _, err = s.db.Exec(`INSERT INTO refunds (order_id, refund_status, refund_amount, created_at) VALUES ($1, 'pending', 0, NOW())`, req.Id)
    if err != nil {
        return nil, err
    }

    return &pb.CancelOrderRes{CancellationStatus: "canceled", RefundStatus: "pending"}, nil
}
