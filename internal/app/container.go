package app

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k-ksu/avito-shop/config"
	"github.com/k-ksu/avito-shop/internal/controller/http/middleware"
	"github.com/k-ksu/avito-shop/internal/repository"
	"github.com/k-ksu/avito-shop/internal/repository/cache"
	"github.com/k-ksu/avito-shop/internal/repository/wrapper"
	"github.com/k-ksu/avito-shop/internal/service"
	"github.com/k-ksu/avito-shop/pkg/postgres"
)

type (
	// Container ...
	Container struct {
		Services     Services
		Servers      Servers
		Repositories Repositories
	}

	// Services ...
	Services struct {
		Auth  *service.JWTAuth
		Users *service.Users
		Shop  *service.Shop
	}

	// Repositories ...
	Repositories struct {
		Wrapper            *wrapper.Transaction
		Users              *repository.Users
		TransactionHistory *repository.TransactionHistory
		Merch              *repository.Merch
		ShopHistory        *repository.ShopHistory
		MerchCache         cache.Merch
	}

	// Servers ...
	Servers struct {
		httpRouter *mux.Router
		httpServer *http.Server
	}
)

// NewContainer ...
func NewContainer(ctx context.Context, cfg *config.Config) *Container {
	sRepositories := initRepos(ctx, cfg)
	sServices := initServices(ctx, cfg, sRepositories)

	return &Container{
		Services:     sServices,
		Servers:      initServers(cfg, sServices),
		Repositories: sRepositories,
	}
}

func initServices(ctx context.Context, cfg *config.Config, sRepos Repositories) Services {
	auth := service.NewJWTAuth(cfg.SigningKey, cfg.TokenTTL)
	shop := service.NewShop(
		sRepos.Wrapper,
		sRepos.TransactionHistory,
		sRepos.Users,
		sRepos.Merch,
		sRepos.ShopHistory,
		sRepos.MerchCache,
	)

	if err := shop.WarmUpCache(ctx); err != nil {
		log.Fatal(err)
	}

	return Services{
		Auth:  auth,
		Users: service.NewUsers(sRepos.Wrapper, sRepos.Users, auth, sRepos.TransactionHistory, sRepos.ShopHistory),
		Shop:  shop,
	}
}

func initRepos(ctx context.Context, cfg *config.Config) Repositories {
	cl, err := postgres.NewClient(ctx, cfg.URL)
	if err != nil {
		log.Fatal("init pg client", err)
	}

	return Repositories{
		Wrapper:            wrapper.NewTransaction(cl),
		Users:              repository.NewUsers(cl),
		TransactionHistory: repository.NewTransactionHistory(cl),
		Merch:              repository.NewMerch(cl),
		ShopHistory:        repository.NewShopHistory(cl),
		MerchCache:         cache.NewMerch(),
	}
}

func initServers(cfg *config.Config, sServices Services) Servers {
	router := mux.NewRouter().StrictSlash(true)
	handler := middleware.Auth(router, sServices.Auth)
	server := &http.Server{ //nolint:gosec
		Addr:    cfg.Host + ":" + cfg.Port,
		Handler: handler,
	}

	return Servers{
		httpRouter: router,
		httpServer: server,
	}
}
