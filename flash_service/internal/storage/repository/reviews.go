package repository

import (
	"database/sql"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
)

type ReviewRepo struct {
	DB *sql.DB
}

func NewReviewRepo(db *sql.DB) *ReviewRepo {
	return &ReviewRepo{
		DB: db,
	}
}

func (r *ReviewRepo) CreateReview(req *pb.CreateReviewReq) (*pb.Void, error) {
	_, err := r.DB.Exec(`INSERT INTO reviews (user_id, product_id, rating, review_text, created_at)
        VALUES ($1, $2, $3, $4, $5)`, req.UserId, req.ProductId, req.Rating, req.ReviewText, req.CreatedAt)

	if err != nil {
		return nil, err
	}

	err = r.updateProductRating(req.ProductId)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *ReviewRepo) GetProductRating(req *pb.GetProductRatingReq) (*pb.ProductRatingRes, error) {
	var res pb.ProductRatingRes
	err := r.DB.QueryRow(`
        SELECT average_rating, total_reviews
        FROM products
        WHERE id = $1
    `, req.ProductId).Scan(&res.AverageRating, &res.TotalReviews)

	if err != nil {
		return nil, err
	}

	res.ProductId = req.ProductId
	return &res, nil
}

func (r *ReviewRepo) updateProductRating(productId string) error {
	var totalReviews int64
	var sumRatings int64

	err := r.DB.QueryRow(`
        SELECT COUNT(*), SUM(rating)
        FROM reviews
        WHERE product_id = $1
    `, productId).Scan(&totalReviews, &sumRatings)

	if err != nil {
		return err
	}

	if totalReviews > 0 {
		averageRating := float64(sumRatings) / float64(totalReviews)
		_, err = r.DB.Exec(`
            INSERT INTO products (id, average_rating, total_reviews)
            VALUES ($1, $2, $3)
            ON CONFLICT (id) DO UPDATE
            SET average_rating = EXCLUDED.average_rating,
                total_reviews = EXCLUDED.total_reviews
        `, productId, averageRating, totalReviews)
		if err != nil {
			return err
		}
	} else {
		_, err = r.DB.Exec(`
            INSERT INTO products (id, average_rating, total_reviews)
            VALUES ($1, 0, 0)
            ON CONFLICT (id) DO UPDATE
            SET average_rating = 0,
                total_reviews = 0
        `, productId)
		if err != nil {
			return err
		}
	}

	return nil
}
