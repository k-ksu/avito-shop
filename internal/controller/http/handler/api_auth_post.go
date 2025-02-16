package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/k-ksu/avito-shop/internal/consts"
	"github.com/k-ksu/avito-shop/internal/controller/http/desc"
	"github.com/k-ksu/avito-shop/internal/controller/http/helper"
	"github.com/k-ksu/avito-shop/internal/errs"
)

// APIAuthPost ...
func (t *AvitoShopAPI) APIAuthPost(w http.ResponseWriter, r *http.Request) {
	var authReq desc.AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&authReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	token, err := t.users.AuthUser(r.Context(), authReq.Username, authReq.Password)
	if err != nil {
		if errors.Is(err, errs.ErrInvalidPassword) {
			http.Error(w, err.Error(), http.StatusUnauthorized)

			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	b, err := json.Marshal(desc.AuthResponse{
		Token: token,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  consts.AuthTokenCookie,
		Value: token,
		Path:  "/",
	})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		helper.WithError(w, http.StatusInternalServerError, err.Error())
	}
}
