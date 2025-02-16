package helper

import (
	"bytes"
	"encoding/json"
	"github.com/k-ksu/avito-shop/internal/controller/http/desc"
	"net/http"
)

func AuthUser(addr, name, pass string) (string, error) {
	loginReq := desc.AuthRequest{
		Username: name,
		Password: pass,
	}

	b, err := json.Marshal(loginReq)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(addr+"/api/auth", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	var authRsp desc.AuthResponse
	if err = json.NewDecoder(resp.Body).Decode(&authRsp); err != nil {
		return "", err
	}

	return authRsp.Token, nil
}
