package handlers

import (
	grpc "flashSale_gateway/internal/gRPC"
	"flashSale_gateway/internal/pkg/kafka"
	"flashSale_gateway/internal/pkg/logger"

	"github.com/go-redis/redis/v8"
)

type Handler struct {
	Clients  grpc.Clients
	Producer kafka.KafkaProducer
	Redis    *redis.Client
	Logger   *logger.Logger
}

func NewHandler(clients grpc.Clients, producer kafka.KafkaProducer, redis *redis.Client, logger *logger.Logger) *Handler {
	return &Handler{Clients: clients, Producer: producer, Redis: redis, Logger: logger}
}
