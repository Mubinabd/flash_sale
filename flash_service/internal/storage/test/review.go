package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/Mubinabd/flash_sale/internal/storage/repository"
)

func TestCreateReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewReviewRepo(db)

	mock.ExpectExec(`INSERT INTO reviews`).
		WithArgs("user1", "product1", 5, "Great product!", time.Now().Format(time.RFC3339)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(`INSERT INTO products`).
		WithArgs("product1", 5.0, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := &pb.CreateReviewReq{
		UserId:     "user1",
		ProductId:  "product1",
		Rating:     5,
		ReviewText: "Great product!",
		CreatedAt:  time.Now().Format(time.RFC3339),
	}

	_, err = repo.CreateReview(req)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetProductRating(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewReviewRepo(db)

	mock.ExpectQuery(`SELECT average_rating, total_reviews`).
		WithArgs("product1").
		WillReturnRows(sqlmock.NewRows([]string{"average_rating", "total_reviews"}).
			AddRow(4.5, 10))

	req := &pb.GetProductRatingReq{
		ProductId: "product1",
	}

	res, err := repo.GetProductRating(req)
	assert.NoError(t, err)
	assert.Equal(t, "product1", res.ProductId)
	assert.Equal(t, float64(4.5), res.AverageRating)
	assert.Equal(t, int64(10), res.TotalReviews)
	assert.NoError(t, mock.ExpectationsWereMet())
}
