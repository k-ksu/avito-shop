package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/k-ksu/avito-shop/internal/model"
)

// JWTAuth ...
type JWTAuth struct {
	signingKey []byte
	tokenTTL   time.Duration
}

// NewJWTAuth ...
func NewJWTAuth(key string, tokenTTL time.Duration) *JWTAuth {
	return &JWTAuth{
		signingKey: []byte(key),
		tokenTTL:   tokenTTL,
	}
}

// ParseToken ...
func (j *JWTAuth) ParseToken(tokenString string) (model.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(_ *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})
	if err != nil {
		return model.Claims{}, fmt.Errorf("jwt.ParseWithClaims: %w", err)
	}

	claims, ok := token.Claims.(*model.Claims)
	if !ok {
		return model.Claims{}, errors.New("token.Claims to model.Claims failed")
	}

	return *claims, nil
}

// GenerateToken ...
func (j *JWTAuth) GenerateToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(j.tokenTTL)

	claims := &model.Claims{
		UserID:   user.ID,
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(j.signingKey)
	if err != nil {
		return "", fmt.Errorf("token.SignedString: %w", err)
	}

	return tokenString, nil
}
