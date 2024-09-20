package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/Mubinabd/flash_sale/internal/storage/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateFlashSale(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewFlashSaleRepo(db)

	req := &pb.CreateFlashSalesReq{
		Name:      "Summer Sale",
		StartTime: "2024-09-01T00:00:00Z",
		EndTime:   "2024-09-05T00:00:00Z",
		Status:    "ACTIVE",
	}

	query := "INSERT INTO flash_sales \\(id, name, start_time, end_time, status\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\)"
	mock.ExpectExec(query).
		WithArgs(sqlmock.AnyArg(), req.Name, req.StartTime, req.EndTime, req.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.CreateFlashSale(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateFlashSale(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewFlashSaleRepo(db)

	req := &pb.UpdateFlashSalesReq{
		Id: "1234",
		Body: &pb.UpdateFlashSale{
			Name:      "Winter Sale",
			StartTime: "2024-12-01T00:00:00Z",
			EndTime:   "2024-12-05T00:00:00Z",
			Status:    "INACTIVE",
		},
	}

	query := `UPDATE flash_sales SET name = \$1, start_time = \$2, end_time = \$3, status = \$4, updated_at = \$5 WHERE id = \$6`
	mock.ExpectExec(query).
		WithArgs(req.Body.Name, req.Body.StartTime, req.Body.EndTime, req.Body.Status, sqlmock.AnyArg(), req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.UpdateFlashSale(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetFlashSale(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewFlashSaleRepo(db)

	req := &pb.GetById{Id: "1234"}

	rows := sqlmock.NewRows([]string{"id", "name", "start_time", "end_time", "status", "created_at"}).
		AddRow("1234", "Black Friday Sale", "2024-11-25T00:00:00Z", "2024-11-29T23:59:59Z", "ACTIVE", time.Now().UTC())

	query := `SELECT id, name, start_time, end_time, status, created_at FROM flash_sales WHERE id = \$1 AND deleted_at = 0`
	mock.ExpectQuery(query).WithArgs(req.Id).WillReturnRows(rows)

	result, err := repo.GetFlashSale(req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "1234", result.Id)
	assert.Equal(t, "Black Friday Sale", result.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteFlashSale(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewFlashSaleRepo(db)

	req := &pb.GetById{Id: "1234"}

	query := `UPDATE flash_sales SET deleted_at = extract\\(epoch from now\\(\\)\\) WHERE id = \$1`
	mock.ExpectExec(query).WithArgs(req.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.DeleteFlashSale(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestListAllFlashSales(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewFlashSaleRepo(db)

	req := &pb.ListAllFlashSalesReq{
		Name:   "Holiday Sale",
		Status: "ACTIVE",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "start_time", "end_time", "status", "created_at"}).
		AddRow(uuid.NewString(), "Holiday Sale", "2024-12-01T00:00:00Z", "2024-12-05T00:00:00Z", "ACTIVE", time.Now().UTC())

	query := `SELECT id, name, start_time, end_time, status, created_at FROM flash_sales WHERE deleted_at = 0 AND name = \$1 AND status = \$2`
	mock.ExpectQuery(query).WithArgs(req.Name, req.Status).WillReturnRows(rows)

	result, err := repo.ListAllFlashSales(req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.FlashSales, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}
