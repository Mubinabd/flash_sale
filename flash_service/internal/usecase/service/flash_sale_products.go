package service

import (
	"context"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	st "github.com/Mubinabd/flash_sale/internal/storage"
	"github.com/Mubinabd/flash_sale/internal/usecase/kafka"
)

type FlashSaleProductService struct {
	storage  st.StorageI
	producer kafka.KafkaProducer
	pb.UnimplementedFlashSaleProductServiceServer
}

func NewFlashSaleProductService(storage st.StorageI, kafka kafka.KafkaProducer) *FlashSaleProductService {
    return &FlashSaleProductService{
        storage:  storage,
        producer: kafka,  
    }
}



func (s *FlashSaleProductService) CreateFlashSaleProduct(ctx context.Context, req *pb.CreateFlashSaleProductReq) (*pb.Void, error) {
	res, err := s.storage.FlashSaleProduct().CreateFlashSaleProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *FlashSaleProductService) UpdateFlashSaleProduct(ctx context.Context, req *pb.UpdateFlashSaleProductReq) (*pb.Void, error) {
	res, err := s.storage.FlashSaleProduct().UpdateFlashSaleProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *FlashSaleProductService) ListAllFlashSaleProducts(ctx context.Context, req *pb.ListAllFlashSaleProductsReq) (*pb.ListAllFlashSaleProductsRes, error) {
	res, err := s.storage.FlashSaleProduct().ListAllFlashSaleProducts(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *FlashSaleProductService) GetFlashSaleProduct(ctx context.Context, req *pb.GetById) (*pb.FlashSaleProduct, error) {
	res, err := s.storage.FlashSaleProduct().GetFlashSaleProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *FlashSaleProductService) DeleteFlashSaleProduct(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.FlashSaleProduct().DeleteFlashSaleProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
