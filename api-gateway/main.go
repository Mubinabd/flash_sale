package main

import (
	"flashSale_gateway/internal/app"
	"flashSale_gateway/internal/pkg/config"
)

func main() {
	cfg := config.Load()

	app.Run(cfg)
}