package service

import (
	"context"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	st "github.com/Mubinabd/flash_sale/internal/storage"
	"github.com/Mubinabd/flash_sale/internal/usecase/kafka"
)

type OrderService struct {
	storage  st.StorageI
	pb.UnimplementedOrderServiceServer
}

func NewOrderService(storage st.StorageI, kafka kafka.KafkaProducer) *OrderService {
	return &OrderService{
		storage: storage,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.Void, error) {
	res, err := s.storage.Order().CreateOrder(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderReq) (*pb.Void, error) {
	res, err := s.storage.Order().UpdateOrder(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OrderService) ListAllOrders(ctx context.Context, req *pb.ListAllOrdersReq) (*pb.ListAllOrdersRes, error) {
	res, err := s.storage.Order().ListAllOrders(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetById) (*pb.Order, error) {
	res, err := s.storage.Order().GetOrder(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.Order().DeleteOrder(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}



func (s *OrderService) GetOrderHistory(ctx context.Context, req *pb.OrderHistoryReq) (*pb.OrderHistoryRes, error) {
	res, err := s.storage.Order().GetOrderHistory(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (s *OrderService) CancelOrder(ctx context.Context, req *pb.GetById) (*pb.CancelOrderRes, error) {
	res, err := s.storage.Order().CancelOrder(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

