package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/Mubinabd/flash_sale/internal/storage/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateFlashSaleProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewFlashSaleProductsRepo(db)

	req := &genproto.CreateFlashSaleProductReq{
		FlashSaleId:       "flashSaleId",
		ProductId:         "productId",
		AvailableQuantity: 10,
		DiscountedPrice:   99.99,
	}

	mock.ExpectExec("INSERT INTO flash_sales_products").
		WithArgs(sqlmock.AnyArg(), req.FlashSaleId, req.ProductId, req.AvailableQuantity, req.DiscountedPrice).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := repo.CreateFlashSaleProduct(req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, &genproto.Void{}, res)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateFlashSaleProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewFlashSaleProductsRepo(db)

	req := &genproto.UpdateFlashSaleProductReq{
		Id: "productId",
		Body: &genproto.UpdateFlashSaleProduct{
			FlashSaleId:       "flashSaleId",
			ProductId:         "productId",
			AvailableQuantity: 20,
			DiscountedPrice:   89.99,
		},
	}

	mock.ExpectExec("UPDATE flash_sales_products").
		WithArgs(req.Body.FlashSaleId, req.Body.ProductId, req.Body.AvailableQuantity, req.Body.DiscountedPrice, sqlmock.AnyArg(), req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := repo.UpdateFlashSaleProduct(req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, &genproto.Void{}, res)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteFlashSaleProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewFlashSaleProductsRepo(db)

	req := &genproto.GetById{Id: "productId"}

	mock.ExpectExec("UPDATE flash_sales_products").
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := repo.DeleteFlashSaleProduct(req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, &genproto.Void{}, res)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetFlashSaleProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewFlashSaleProductsRepo(db)

	req := &genproto.GetById{Id: "productId"}

	mock.ExpectQuery("SELECT (.+) FROM flash_sales_products").
		WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "available_quantity", "discounted_price",
			"flash_sale_id", "flash_sale_name", "flash_sale_start_time",
			"flash_sale_end_time", "flash_sale_status",
			"product_id", "product_name", "product_price",
			"product_description", "product_image_url", "product_stock_quantity",
		}).AddRow(
			"productId", 10, 99.99,
			"flashSaleId", "Flash Sale", "2022-01-01 00:00:00",
			"2022-12-31 23:59:59", "active",
			"productId", "Product", 199.99,
			"Description", "http://image.url", 50,
		))

	res, err := repo.GetFlashSaleProduct(req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, req.Id, res.Id)
	assert.Equal(t, int64(10), res.AvailableQuantity)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
