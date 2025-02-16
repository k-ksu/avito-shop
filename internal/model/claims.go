package model

import "github.com/dgrijalva/jwt-go"

// Claims ...
type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
