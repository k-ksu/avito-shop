package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/k-ksu/avito-shop/config"
	httpcontroller "github.com/k-ksu/avito-shop/internal/controller/http/handler"
)

// Run ...
func Run(ctx context.Context, cfg *config.Config) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	cont := NewContainer(ctx, cfg)

	impl := httpcontroller.NewAvitoShopAPI(cont.Services.Users, cont.Services.Shop)
	impl.RegisterGateway(cont.Servers.httpRouter, cfg.SwaggerAddr)

	go func() {
		if err := cont.Servers.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err, "Server error")
		}
	}()

	log.Println("Server started at address:", cont.Servers.httpServer.Addr)

	<-quit
	log.Println("Shutting down server...")

	if err := cont.Servers.httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

	return nil
}
