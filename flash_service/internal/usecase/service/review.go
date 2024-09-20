package service

import (
	"context"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	st "github.com/Mubinabd/flash_sale/internal/storage"
	"github.com/Mubinabd/flash_sale/internal/usecase/kafka"
)

type ReviewService struct {
	storage  st.StorageI
	pb.UnimplementedReviewServiceServer
}

func NewReviewService(storage st.StorageI, kafka kafka.KafkaProducer) *ReviewService {
	return &ReviewService{
		storage: storage,
	}
}

func (s *ReviewService) CreateReview(ctx context.Context, req *pb.CreateReviewReq) (*pb.Void, error) {
	res, err := s.storage.Review().CreateReview(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (s *ReviewService) GetProductRating(ctx context.Context, req *pb.GetProductRatingReq) (*pb.ProductRatingRes, error) {
	res, err := s.storage.Review().GetProductRating(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
