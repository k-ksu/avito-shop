package middleware

import (
	"net/http"
	"strings"

	"github.com/k-ksu/avito-shop/internal/consts"
	"github.com/k-ksu/avito-shop/internal/controller/http/helper"
	"github.com/k-ksu/avito-shop/internal/model"
)

// Auther ...
//
// nolint:misspell
type Auther interface {
	ParseToken(tokenString string) (model.Claims, error)
}

// Auth ...
//
// nolint:misspell
func Auth(next http.Handler, auth Auther) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if isAuthException(req.URL.Path) {
			next.ServeHTTP(w, req)

			return
		}

		token, err := helper.ExtractToken(req)
		if err != nil {
			helper.WithError(w, http.StatusInternalServerError, err.Error())

			return
		}

		if token == "" {
			helper.WithError(w, http.StatusUnauthorized, "token is empty")

			return
		}

		claims, err := auth.ParseToken(token)
		if err != nil {
			helper.WithError(w, http.StatusInternalServerError, err.Error())

			return
		}

		if err = claims.Valid(); err != nil {
			helper.WithError(w, http.StatusUnauthorized, "token expired, please reauthenticate")

			return
		}

		next.ServeHTTP(w, req)
	}
}

func isAuthException(urlPath string) bool {
	return urlPath == consts.AuthPath || strings.HasPrefix(urlPath, consts.SwaggerPath)
}
