package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/k-ksu/avito-shop/internal/consts"
	"github.com/k-ksu/avito-shop/internal/controller/http/desc"
)

// WithError ...
func WithError(w http.ResponseWriter, code int, message string) {
	b, err := json.Marshal(desc.ErrorResponse{Errors: message})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(code)
	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// ExtractToken ...
func ExtractToken(req *http.Request) (string, error) {
	cookie, err := req.Cookie(consts.AuthTokenCookie)
	if err == nil {
		return cookie.Value, nil // токен в куках
	}

	if !errors.Is(err, http.ErrNoCookie) { // ошибка, но не отсутсвие куки
		return "", fmt.Errorf("req.Cookie: %w", err)
	}

	return req.Header.Get("Authorization"), nil // пытаемся взять из заголовка
}
