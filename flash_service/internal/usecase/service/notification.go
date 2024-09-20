package service

import (
	"context"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/Mubinabd/flash_sale/internal/storage"
	"github.com/Mubinabd/flash_sale/internal/usecase/kafka"
)

type NotificationService struct {
	stg      storage.StorageI
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationService(stg storage.StorageI, kafka kafka.KafkaProducer) *NotificationService {
	return &NotificationService{stg: stg}
}

func (s *NotificationService) CreateNotification(ctx context.Context, req *pb.NotificationCreate) (*pb.Void, error) {
	return s.stg.Notification().CreateNotification(req)
}
func (s *NotificationService) DeleteNotification(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
	return s.stg.Notification().DeleteNotification(req)
}
func (s *NotificationService) UpdateNotification(ctx context.Context, req *pb.NotificationUpdate) (*pb.Void, error) {
	return s.stg.Notification().UpdateNotification(req)
}
func (s *NotificationService) GetNotifications(ctx context.Context, req *pb.NotifFilter) (*pb.NotificationList, error) {
	return s.stg.Notification().GetNotifications(req)
}
func (s *NotificationService) GetNotification(ctx context.Context, req *pb.GetById) (*pb.NotificationGet, error) {
	return s.stg.Notification().GetNotification(req)
}
