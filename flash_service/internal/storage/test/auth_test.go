package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/Mubinabd/flash_sale/internal/storage/repository"
	"golang.org/x/crypto/bcrypt"
)

func TestRegister(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	authRepo := repository.NewAuthRepo(db)

	req := &pb.RegisterReq{
		Username:    "testuser",
		Email:       "test@example.com",
		Password:    "password",
		FullName:    "Test User",
		DateOfBirth: "2000-01-01",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO users \(.+\) VALUES \(.+\) RETURNING id`).
		WithArgs(req.Username, req.Email, req.Password, req.FullName, req.DateOfBirth).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectExec(`INSERT INTO settings \(.+\)`).
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	_, err = authRepo.Register(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	authRepo := repository.NewAuthRepo(db)

	req := &pb.LoginReq{
		Username: "testuser",
		Password: "password",
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	mock.ExpectQuery(`SELECT id, username, email, role, password FROM users WHERE username = \$1`).
		WithArgs(req.Username).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "role", "password"}).
			AddRow("1", req.Username, "test@example.com", "user", string(passwordHash)))

	res, err := authRepo.Login(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if res.Username != req.Username {
		t.Errorf("expected username %s, got %s", req.Username, res.Username)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestForgotPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	authRepo := repository.NewAuthRepo(db)

	req := &pb.GetByEmail{
		Email: "test@example.com",
	}

	mock.ExpectQuery(`SELECT email FROM users WHERE email = \$1`).
		WithArgs(req.Email).
		WillReturnRows(sqlmock.NewRows([]string{"email"}).AddRow(req.Email))

	_, err = authRepo.ForgotPassword(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveRefreshToken(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	authRepo := repository.NewAuthRepo(db)

	req := &pb.RefToken{
		UserId: "1",
		Token:  "some-token",
	}

	mock.ExpectExec(`INSERT INTO tokens \(user_id, token\) VALUES \(\$1, \$2\)`).
		WithArgs(req.UserId, req.Token).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = authRepo.SaveRefreshToken(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	authRepo := repository.NewAuthRepo(db)

	req := &pb.ListUserReq{
		Username: "",
		FullName: "",
		Pagination: &pb.Pagination{
			Limit:  10,
			Offset: 0,
		},
	}

	mock.ExpectQuery(`SELECT id, username, full_name, email, date_of_birth, role FROM users WHERE deleted_at=0 AND role = 'user' LIMIT \$1 OFFSET \$2`).
		WithArgs(req.Pagination.Limit, req.Pagination.Offset).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "full_name", "email", "date_of_birth", "role"}).
			AddRow("1", "testuser", "Test User", "test@example.com", "2000-01-01", "user"))

	res, err := authRepo.GetAllUsers(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(res.Users) != 1 {
		t.Errorf("expected 1 user, got %d", len(res.Users))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	authRepo := repository.NewAuthRepo(db)

	req := &pb.GetById{
		Id: "1",
	}

	mock.ExpectQuery(`SELECT id, username, full_name, email, date_of_birth, role FROM users WHERE id = \$1 AND deleted_at=0`).
		WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "full_name", "email", "date_of_birth", "role"}).
			AddRow("1", "testuser", "Test User", "test@example.com", "2000-01-01", "user"))

	res, err := authRepo.GetUserById(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if res.Username != "testuser" {
		t.Errorf("expected username %s, got %s", "testuser", res.Username)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
