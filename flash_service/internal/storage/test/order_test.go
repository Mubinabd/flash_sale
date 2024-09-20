package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/Mubinabd/flash_sale/internal/storage/repository"
)

func TestUpdateOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not mock db: %v", err)
	}
	defer db.Close()

	repo := repository.NewOrderRepo(db)
	req := &pb.UpdateOrderReq{
		Id: "order-1",
		Body: &pb.UpdateOrder{
			UserID:      "fdc7af50-c99d-420c-a74a-43be3cc11c73",
			FlashSaleID: "e8a127d1-b129-4023-85c4-0743a27dd61f",
			OrderStatus: "completed",
		},
	}

	mock.ExpectExec("UPDATE orders SET").WithArgs(req.Body.UserID, req.Body.FlashSaleID, req.Body.OrderStatus, sqlmock.AnyArg(), req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.UpdateOrder(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestGetOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not mock db: %v", err)
	}
	defer db.Close()

	repo := repository.NewOrderRepo(db)
	req := &pb.GetById{Id: "order-1"}

	rows := sqlmock.NewRows([]string{
		"id", "user_id", "username", "email", "full_name",
		"date_of_birth", "flash_sale_id", "name", "start_time",
		"end_time", "status", "status", "created_at",
	}).AddRow("order-1", "fdc7af50-c99d-420c-a74a-43be3cc11c73", "john_doe", "john@example.com", "John Doe",
		"1990-01-01", "flash-e8a127d1-b129-4023-85c4-0743a27dd61f", "Flash Sale", time.Now(), time.Now(), "active", "pending", time.Now())

	mock.ExpectQuery("SELECT").WithArgs(req.Id).WillReturnRows(rows)

	res, err := repo.GetOrder(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if res.Id != "order-1" {
		t.Errorf("expected order ID to be order-1, got %v", res.Id)
	}
}

func TestListAllOrders(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not mock db: %v", err)
	}
	defer db.Close()

	repo := repository.NewOrderRepo(db)

	rows := sqlmock.NewRows([]string{
		"id", "user_id", "username", "email", "full_name",
		"date_of_birth", "flash_sale_id", "name", "start_time",
		"end_time", "status", "status", "created_at",
	}).AddRow("order-1", "fdc7af50-c99d-420c-a74a-43be3cc11c73", "john_doe", "john@example.com", "John Doe",
		"1990-01-01", "flash-e8a127d1-b129-4023-85c4-0743a27dd61f", "Flash Sale", time.Now(), time.Now(), "active", "pending", time.Now())

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	res, err := repo.ListAllOrders(&pb.ListAllOrdersReq{})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(res.Orders) != 1 {
		t.Errorf("expected 1 order, got %v", len(res.Orders))
	}
}

func TestDeleteOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not mock db: %v", err)
	}
	defer db.Close()

	repo := repository.NewOrderRepo(db)
	req := &pb.GetById{Id: "order-1"}

	mock.ExpectExec("UPDATE orders SET").WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.DeleteOrder(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
