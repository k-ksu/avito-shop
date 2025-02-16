package helper

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func BuyItem(t *testing.T, item, addr, token string) {
	t.Helper()

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, addr+"/api/buy/"+item, nil)
	req.Header.Add("Authorization", token)
	req.Header.Add("accept", "application/json")

	req.WithContext(context.Background())
	resp, err := client.Do(req)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
