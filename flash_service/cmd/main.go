package main

import (
	"github.com/Mubinabd/flash_sale/internal/app"
	"github.com/Mubinabd/flash_sale/internal/pkg/config"
)

func main() {
	config := config.Load()

	app.Run(&config)
}
