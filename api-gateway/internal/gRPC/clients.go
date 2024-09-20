package grpc

import (
	"flashSale_gateway/internal/pkg/config"
	pb "flashSale_gateway/internal/pkg/genproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	Auth             pb.AuthServiceClient
	Order            pb.OrderServiceClient
	FlashSale        pb.FlashSaleServiceClient
	FlashSaleProduct pb.FlashSaleProductServiceClient
	Product          pb.ProductServiceClient
	Transaction      pb.TransactionServiceClient
	User             pb.UserServiceClient
	Notification     pb.NotificationServiceClient
	Social           pb.SocialSharingServiceClient
	Review           pb.ReviewServiceClient
}

func NewClients(cfg *config.Config) (*Clients, error) {
	service_conn, err := grpc.NewClient("flash_sale_service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	userClient := pb.NewUserServiceClient(service_conn)
	orderClient := pb.NewOrderServiceClient(service_conn)
	productClient := pb.NewProductServiceClient(service_conn)
	authClient := pb.NewAuthServiceClient(service_conn)
	flashSaleClient := pb.NewFlashSaleServiceClient(service_conn)
	flashSaleProductClient := pb.NewFlashSaleProductServiceClient(service_conn)
	transactionClient := pb.NewTransactionServiceClient(service_conn)
	notificationClient := pb.NewNotificationServiceClient(service_conn)
	reviewClient := pb.NewReviewServiceClient(service_conn)
	socialClient := pb.NewSocialSharingServiceClient(service_conn)

	return &Clients{
		Auth:             authClient,
		Order:            orderClient,
		FlashSale:        flashSaleClient,
		Product:          productClient,
		Transaction:      transactionClient,
		User:             userClient,
		FlashSaleProduct: flashSaleProductClient,
		Notification:     notificationClient,
		Review:           reviewClient,
		Social:           socialClient,
	}, nil
}
