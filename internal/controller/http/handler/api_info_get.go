package handler

import (
	"encoding/json"
	"net/http"

	"github.com/k-ksu/avito-shop/internal/controller/http/desc"
	"github.com/k-ksu/avito-shop/internal/controller/http/helper"
)

// APIInfoGet ...
func (t *AvitoShopAPI) APIInfoGet(w http.ResponseWriter, r *http.Request) {
	user, err := t.getUser(w, r)
	if err != nil {
		return
	}

	userInfo, err := t.users.UserInfo(r.Context(), user)
	if err != nil {
		helper.WithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	var rsp desc.InfoResponse
	rsp.Coins = userInfo.Coins
	rsp.Inventory = make([]desc.InfoResponseInventory, 0, len(userInfo.Inventory))
	for _, item := range userInfo.Inventory {
		rsp.Inventory = append(rsp.Inventory, desc.InfoResponseInventory{
			Type:     item.Type,
			Quantity: item.Quantity,
		})
	}

	rsp.CoinHistory = new(desc.InfoResponseCoinHistory)
	rsp.CoinHistory.Sent = make([]desc.InfoResponseCoinHistorySent, 0, len(userInfo.CoinHistory.Sent))
	for _, item := range userInfo.CoinHistory.Sent {
		rsp.CoinHistory.Sent = append(rsp.CoinHistory.Sent, desc.InfoResponseCoinHistorySent{
			ToUser: item.ToUser,
			Amount: item.Amount,
		})
	}

	rsp.CoinHistory.Received = make([]desc.InfoResponseCoinHistoryReceived, 0, len(userInfo.CoinHistory.Received))
	for _, item := range userInfo.CoinHistory.Received {
		rsp.CoinHistory.Received = append(rsp.CoinHistory.Received, desc.InfoResponseCoinHistoryReceived{
			FromUser: item.FromUser,
			Amount:   item.Amount,
		})
	}

	b, err := json.Marshal(rsp)
	if err != nil {
		helper.WithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		helper.WithError(w, http.StatusInternalServerError, err.Error())
	}
}
