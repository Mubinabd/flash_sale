package app

import (
	"log"
	"net"

	"github.com/Mubinabd/flash_sale/internal/pkg/config"
	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/Mubinabd/flash_sale/internal/pkg/postgres"
	"github.com/Mubinabd/flash_sale/internal/storage/repository"
	"github.com/Mubinabd/flash_sale/internal/usecase/kafka"
	"github.com/Mubinabd/flash_sale/internal/usecase/service"
	"google.golang.org/grpc"
)

func Run(cf *config.Config) {
	// connect to postgres
	pgm, err := postgres.New(cf)
	if err != nil {
		log.Fatal(err)
	}
	defer pgm.Close()
	// connect to kafka producer
	kf, err := kafka.NewKafkaProducer([]string{cf.KafkaUrl})
	if err != nil {
		log.Fatal(err)
	}

	// repo
	db := repository.NewStorage(pgm.DB)

	k_handler := KafkaHandler{
		auth:             service.NewAuthService(db, kf),
		user:             service.NewUserService(db, kf),
		order:            service.NewOrderService(db, kf),
		product:          service.NewProductService(db, kf),
		notification:     service.NewNotificationService(db, kf),
		flashSaleProduct: service.NewFlashSaleProductService(db, kf),
		flashSale:        service.NewFlashSaleService(db, kf),
	}

	if err := Register(&k_handler, cf); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", cf.GRPCPort)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	// set grpc server
	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, service.NewAuthService(db, kf))
	pb.RegisterUserServiceServer(server, service.NewUserService(db, kf))
	pb.RegisterFlashSaleProductServiceServer(server, service.NewFlashSaleProductService(db, kf))
	pb.RegisterFlashSaleServiceServer(server, service.NewFlashSaleService(db, kf))
	pb.RegisterNotificationServiceServer(server, service.NewNotificationService(db, kf))
	pb.RegisterOrderServiceServer(server, service.NewOrderService(db, kf))
	pb.RegisterProductServiceServer(server, service.NewProductService(db, kf))
	pb.RegisterReviewServiceServer(server, service.NewReviewService(db, kf))
	pb.RegisterSocialSharingServiceServer(server, service.NewSocialService(db, kf))

	// start server
	log.Println("Server started on", cf.GRPCPort)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	defer lis.Close()
}
