package service

import (
	"context"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	st "github.com/Mubinabd/flash_sale/internal/storage"
	"github.com/Mubinabd/flash_sale/internal/usecase/kafka"
)

type SocialService struct {
	storage  st.StorageI
	pb.UnimplementedSocialSharingServiceServer
}

func NewSocialService(storage st.StorageI, kafka kafka.KafkaProducer) *SocialService {
	return &SocialService{
		storage: storage,
	}
}

func (s *SocialService) ShareDeal(ctx context.Context, req *pb.ShareDealReq) (*pb.Void, error) {
	res, err := s.storage.Social().ShareDeal(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SocialService) GetSharingStats(ctx context.Context, req *pb.GetSharingStatsReq) (*pb.SharingStatsRes, error) {
	res, err := s.storage.Social().GetSharingStats(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
