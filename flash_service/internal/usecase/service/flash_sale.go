package service

import (
	"context"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	st "github.com/Mubinabd/flash_sale/internal/storage"
	"github.com/Mubinabd/flash_sale/internal/usecase/kafka"
)

type FlashSaleService struct {
	storage  st.StorageI
	pb.UnimplementedFlashSaleServiceServer
}

func NewFlashSaleService(storage st.StorageI, kafka kafka.KafkaProducer) *FlashSaleService {
	return &FlashSaleService{
		storage: storage,
	}
}

func (s *FlashSaleService) CreateFlashSale(ctx context.Context, req *pb.CreateFlashSalesReq) (*pb.Void, error) {
	res, err := s.storage.FlashSale().CreateFlashSale(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *FlashSaleService) UpdateFlashSale(ctx context.Context, req *pb.UpdateFlashSalesReq) (*pb.Void, error) {
	res, err := s.storage.FlashSale().UpdateFlashSale(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *FlashSaleService) ListAllFlashSales(ctx context.Context, req *pb.ListAllFlashSalesReq) (*pb.ListAllFlashSalesRes, error) {
	res, err := s.storage.FlashSale().ListAllFlashSales(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *FlashSaleService) GetFlashSale(ctx context.Context, req *pb.GetById) (*pb.FlashSale, error) {
	res, err := s.storage.FlashSale().GetFlashSale(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *FlashSaleService) DeleteFlashSale(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.FlashSale().DeleteFlashSale(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *FlashSaleService) AddProductToFlashSale(ctx context.Context, req *pb.AddProductReq) (*pb.Void, error) {
	res, err := s.storage.FlashSale().AddProductToFlashSale(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *FlashSaleService) RemoveProductFromFlashSale(ctx context.Context, req *pb.RemoveProductReq) (*pb.Void, error) {
	res, err := s.storage.FlashSale().RemoveProductFromFlashSale(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *FlashSaleService) CancelFlashSale(ctx context.Context, req *pb.GetById) (*pb.CancelFlashSaleRes, error) {
	res, err := s.storage.FlashSale().CancelFlashSale(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (s *FlashSaleService) GetStoreLocation(ctx context.Context, req *pb.GetStoreLocationReq) (*pb.StoreLocation, error) {
	res, err := s.storage.FlashSale().GetStoreLocation(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}