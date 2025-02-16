package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gorilla/mux"
	"github.com/k-ksu/avito-shop/internal/consts"
	"github.com/k-ksu/avito-shop/internal/controller/http/helper"
	"github.com/k-ksu/avito-shop/internal/controller/http/middleware"
	"github.com/k-ksu/avito-shop/internal/model"
	httpSwagger "github.com/swaggo/http-swagger"
)

const apiName = "avito-shop"

type (
	// Route ...
	Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}

	// Routes ...
	Routes []Route

	// UsersServicer ...
	UsersServicer interface {
		GetUser(tokenString string) (model.User, error)
		AuthUser(ctx context.Context, name, pass string) (string, error)
		UserInfo(ctx context.Context, user model.User) (model.UserInfo, error)
	}

	// ShopServicer ...
	ShopServicer interface {
		SendCoins(ctx context.Context, fromUser, toUser model.User, amount int32) error
		BuyItem(ctx context.Context, user model.User, item string) error
	}

	// AvitoShopAPI ...
	AvitoShopAPI struct {
		users UsersServicer
		shop  ShopServicer
	}
)

// NewAvitoShopAPI ...
func NewAvitoShopAPI(users UsersServicer, shop ShopServicer) *AvitoShopAPI {
	return &AvitoShopAPI{
		users: users,
		shop:  shop,
	}
}

// RegisterGateway ...
func (t *AvitoShopAPI) RegisterGateway(router *mux.Router, addr string) {
	for _, route := range t.GetRoutes() {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = middleware.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	spec, err := loadSpec(addr)
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/swagger_spec", byteHandler(spec))
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger_spec"),
	))
}

// GetRoutes ...
func (t *AvitoShopAPI) GetRoutes() Routes {
	return Routes{
		Route{
			"APIAuthPost",
			strings.ToUpper("Post"),
			consts.AuthPath,
			t.APIAuthPost,
		},

		Route{
			"APIBuyItemGet",
			strings.ToUpper("Get"),
			"/api/buy/{item}",
			t.APIBuyItemGet,
		},

		Route{
			"APIInfoGet",
			strings.ToUpper("Get"),
			"/api/info",
			t.APIInfoGet,
		},

		Route{
			"APISendCoinPost",
			strings.ToUpper("Post"),
			"/api/sendCoin",
			t.APISendCoinPost,
		},
	}
}

func loadSpec(addr string) ([]byte, error) {
	file, err := os.Open(path.Join(consts.APIFolder, apiName+".json"))
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	defer file.Close()
	b, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	var spec map[string]any
	err = json.Unmarshal(b, &spec)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	spec["host"] = addr

	return json.Marshal(spec)
}

func byteHandler(b []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write(b)
		if err != nil {
			helper.WithError(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func (t *AvitoShopAPI) getUser(w http.ResponseWriter, r *http.Request) (model.User, error) {
	token, err := helper.ExtractToken(r)
	if err != nil {
		helper.WithError(w, http.StatusInternalServerError, err.Error())

		return model.User{}, err
	}

	if token == "" {
		helper.WithError(w, http.StatusUnauthorized, "token is empty")

		return model.User{}, errors.New("token is empty")
	}

	user, err := t.users.GetUser(token)
	if err != nil {
		helper.WithError(w, http.StatusInternalServerError, err.Error())

		return model.User{}, err
	}

	return user, nil
}
