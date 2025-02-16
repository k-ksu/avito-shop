package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/k-ksu/avito-shop/internal/controller/http/desc"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func SendCoin(t *testing.T, addr, to, token string) {
	var sendCoinReq desc.SendCoinRequest
	sendCoinReq.ToUser = to
	sendCoinReq.Amount = 100

	b, err := json.Marshal(sendCoinReq)
	require.NoError(t, err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, addr+"/api/sendCoin", bytes.NewBuffer(b))
	req.Header.Add("Authorization", token)
	req.Header.Add("accept", "application/json")

	req.WithContext(context.Background())
	resp, err := client.Do(req)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
