package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/Mubinabd/flash_sale/internal/storage/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := repository.NewProductRepo(db)

	req := &pb.CreateProductReq{
		Name:          "Test Product",
		Description:   "Test Description",
		Price:         10.0,
		ImageUrl:      "http://example.com/image.jpg",
		StockQuantity: 100,
	}

	id := uuid.NewString()
	mock.ExpectExec("INSERT INTO products").WithArgs(id, req.Name, req.Description, req.Price, req.ImageUrl, req.StockQuantity).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.CreateProduct(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := repository.NewProductRepo(db)

	req := &pb.UpdateProductReq{
		Id: "some-product-id",
		Body: &pb.UpdateBody{
			Name:  "Updated Product",
			Price: 20.0,
		},
	}

	mock.ExpectExec("UPDATE products").WithArgs(req.Body.Name, req.Body.Price, sqlmock.AnyArg(), req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.UpdateProduct(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := repository.NewProductRepo(db)

	id := "some-product-id"
	expectedProduct := &pb.Products{
		Id:            id,
		Name:          "Test Product",
		Description:   "Test Description",
		Price:         10.0,
		ImageUrl:      "http://example.com/image.jpg",
		StockQuantity: 100,
	}

	mock.ExpectQuery("SELECT id, name, description, price, image_url, stock_quantity").WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "price", "image_url", "stock_quantity"}).
			AddRow(expectedProduct.Id, expectedProduct.Name, expectedProduct.Description, expectedProduct.Price, expectedProduct.ImageUrl, expectedProduct.StockQuantity))

	req := &pb.GetById{Id: id}
	product, err := repo.GetProduct(req)
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestListAllProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := repository.NewProductRepo(db)

	mock.ExpectQuery("SELECT id, name, description, price, image_url, stock_quantity").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "price", "image_url", "stock_quantity"}).
		AddRow("1", "Product 1", "Description 1", 10.0, "http://example.com/image1.jpg", 100).
		AddRow("2", "Product 2", "Description 2", 20.0, "http://example.com/image2.jpg", 200))

	req := &pb.ListAllProductsReq{}
	res, err := repo.ListAllProducts(req)
	assert.NoError(t, err)
	assert.Len(t, res.Products, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := repository.NewProductRepo(db)

	req := &pb.GetById{Id: "some-product-id"}

	mock.ExpectExec("UPDATE products").WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.DeleteProduct(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
