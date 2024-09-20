package app

import (
	"log"
	"path/filepath"
	"runtime"

	grpc "flashSale_gateway/internal/gRPC"
	"flashSale_gateway/internal/http"
	"flashSale_gateway/internal/http/handlers"
	"flashSale_gateway/internal/pkg/config"
	"flashSale_gateway/internal/pkg/kafka"
	"flashSale_gateway/internal/pkg/logger"

	"github.com/go-redis/redis/v8"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func Run(cfg config.Config) {
	logger := logger.NewLogger(basepath, cfg.LogPath)
	clients, err := grpc.NewClients(&cfg)
	if err != nil {
		logger.ERROR.Println("Failed to create gRPC clients", err)
		log.Fatal(err)
		return
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	//connect to kafka
	broker := []string{cfg.KafkaUrl}
	kafka, err := kafka.NewKafkaProducer(broker)
	if err != nil {
		logger.ERROR.Println("Failed to connect to Kafka", err)
		log.Fatal(err)
		return
	}
	defer kafka.Close()

	// make handler
	h := handlers.NewHandler(*clients, kafka, rdb, logger)

	// make gin
	router := http.NewGin(h)

	// start server
	router.Run(":5050")
}
