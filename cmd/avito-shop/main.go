package main

import (
	"context"
	"log"

	"github.com/k-ksu/avito-shop/config"
	"github.com/k-ksu/avito-shop/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.New()
	if err != nil {
		log.Print("Failed to load config", err)

		return
	}

	if err = app.Run(ctx, cfg); err != nil {
		log.Print("App execution failed", err)
	}
}
