package service

import (
	"context"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/Mubinabd/flash_sale/internal/storage"
	"github.com/Mubinabd/flash_sale/internal/usecase/kafka"
)

type AuthService struct {
	storage  storage.StorageI
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(storage storage.StorageI, kafka kafka.KafkaProducer) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.Void, error) {
	res, err := s.storage.Auth().Register(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.User, error) {
	res, err := s.storage.Auth().Login(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AuthService) ForgotPassword(ctx context.Context, req *pb.GetByEmail) (*pb.Void, error) {
	res, err := s.storage.Auth().ForgotPassword(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AuthService) ResetPassword(ctx context.Context, req *pb.ResetPassReq) (*pb.Void, error) {
	res, err := s.storage.Auth().ResetPassword(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AuthService) SaveRefreshToken(ctx context.Context, req *pb.RefToken) (*pb.Void, error) {
	res, err := s.storage.Auth().SaveRefreshToken(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AuthService) GetAllUsers(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserRes, error) {
	res, err := s.storage.Auth().GetAllUsers(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *AuthService) GetUserById(ctx context.Context, req *pb.GetById) (*pb.UserRes, error) {
	res, err := s.storage.Auth().GetUserById(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
