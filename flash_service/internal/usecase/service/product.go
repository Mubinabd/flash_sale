package service

import (
	"context"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	st "github.com/Mubinabd/flash_sale/internal/storage"
	"github.com/Mubinabd/flash_sale/internal/usecase/kafka"
)

type ProductService struct {
	storage st.StorageI
	pb.UnimplementedProductServiceServer
}

func NewProductService(storage st.StorageI, kafka kafka.KafkaProducer) *ProductService {
	return &ProductService{
		storage: storage,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductReq) (*pb.Void, error) {
	res, err := s.storage.Product().CreateProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductReq) (*pb.Void, error) {
	res, err := s.storage.Product().UpdateProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ProductService) ListAllProducts(ctx context.Context, req *pb.ListAllProductsReq) (*pb.ListAllProductsRes, error) {
	res, err := s.storage.Product().ListAllProducts(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetById) (*pb.Products, error) {
	res, err := s.storage.Product().GetProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	res, err := s.storage.Product().DeleteProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
