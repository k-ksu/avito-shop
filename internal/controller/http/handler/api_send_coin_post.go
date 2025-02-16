package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/k-ksu/avito-shop/internal/controller/http/desc"
	"github.com/k-ksu/avito-shop/internal/controller/http/helper"
	"github.com/k-ksu/avito-shop/internal/errs"
	"github.com/k-ksu/avito-shop/internal/model"
)

// APISendCoinPost ...
func (t *AvitoShopAPI) APISendCoinPost(w http.ResponseWriter, r *http.Request) {
	user, err := t.getUser(w, r)
	if err != nil {
		return
	}

	var sendCoinReq desc.SendCoinRequest
	if err = json.NewDecoder(r.Body).Decode(&sendCoinReq); err != nil {
		helper.WithError(w, http.StatusBadRequest, err.Error())

		return
	}

	if sendCoinReq.Amount <= 0 {
		helper.WithError(w, http.StatusBadRequest, "amount is invalid")

		return
	}

	if sendCoinReq.ToUser == user.Name {
		helper.WithError(w, http.StatusBadRequest, "you cannot send to yourself")

		return
	}

	if err = t.shop.SendCoins(
		r.Context(),
		user,
		model.User{Name: sendCoinReq.ToUser},
		sendCoinReq.Amount,
	); err != nil {
		if errors.Is(err, errs.ErrNotEnoughMoney) || errors.Is(err, errs.ErrUserNotExists) {
			helper.WithError(w, http.StatusBadRequest, err.Error())

			return
		}

		helper.WithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
