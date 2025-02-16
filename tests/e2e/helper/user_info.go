package helper

import (
	"encoding/json"
	"github.com/k-ksu/avito-shop/internal/controller/http/desc"
	"github.com/k-ksu/avito-shop/internal/model"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func UserInfo(t *testing.T, addr, token string) model.UserInfo {
	client := &http.Client{}
	req, err := http.NewRequest("GET", addr+"/api/info", nil)
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	require.NoError(t, err)

	var authRsp desc.InfoResponse
	err = json.NewDecoder(resp.Body).Decode(&authRsp)
	require.NoError(t, err)

	var userInfo model.UserInfo

	userInfo.Coins = authRsp.Coins

	userInfo.Inventory = make([]model.Inventory, 0, len(authRsp.Inventory))
	for _, inventory := range authRsp.Inventory {
		userInfo.Inventory = append(userInfo.Inventory, model.Inventory{
			Quantity: inventory.Quantity,
			Type:     inventory.Type,
		})
	}

	userInfo.CoinHistory.Sent = make([]model.SentCoins, 0, len(authRsp.CoinHistory.Sent))
	for _, s := range authRsp.CoinHistory.Sent {
		userInfo.CoinHistory.Sent = append(userInfo.CoinHistory.Sent, model.SentCoins{
			Amount: s.Amount,
			ToUser: s.ToUser,
		})
	}

	userInfo.CoinHistory.Received = make([]model.ReceivedCoins, 0, len(authRsp.CoinHistory.Received))
	for _, r := range authRsp.CoinHistory.Received {
		userInfo.CoinHistory.Received = append(userInfo.CoinHistory.Received, model.ReceivedCoins{
			Amount:   r.Amount,
			FromUser: r.FromUser,
		})
	}

	return userInfo
}
