package service

import (
	"context"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	st "github.com/Mubinabd/flash_sale/internal/storage"
	"github.com/Mubinabd/flash_sale/internal/usecase/kafka"
)

type UserService struct {
	storage  st.StorageI
	pb.UnimplementedUserServiceServer
}

func NewUserService(storage st.StorageI, kafka kafka.KafkaProducer) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (s *UserService) GetProfile(ctx context.Context, req *pb.GetByID) (*pb.UserRes, error) {
	res, err := s.storage.User().GetProfile(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) EditProfile(ctx context.Context, req *pb.UserRes) (*pb.UserRes, error) {
	res, err := s.storage.User().EditProfile(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) ChangePassword(ctx context.Context, req *pb.ChangePasswordReq) (*pb.Void, error) {
	res, err := s.storage.User().ChangePassword(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) GetSetting(ctx context.Context, req *pb.GetByID) (*pb.Setting, error) {
	res, err := s.storage.User().GetSetting(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) EditSetting(ctx context.Context, req *pb.SettingReq) (*pb.Void, error) {
	res, err := s.storage.User().EditSetting(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.GetByID) (*pb.Void, error) {
	res, err := s.storage.User().DeleteUser(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
