package handler

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k-ksu/avito-shop/internal/controller/http/helper"
	"github.com/k-ksu/avito-shop/internal/errs"
)

// APIBuyItemGet ...
func (t *AvitoShopAPI) APIBuyItemGet(w http.ResponseWriter, r *http.Request) {
	user, err := t.getUser(w, r)
	if err != nil {
		return
	}

	item := mux.Vars(r)["item"]

	if err = t.shop.BuyItem(r.Context(), user, item); err != nil {
		if errors.Is(err, errs.ErrNotEnoughMoney) {
			helper.WithError(w, http.StatusBadRequest, err.Error())

			return
		}

		helper.WithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
