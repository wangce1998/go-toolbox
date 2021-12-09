package xjwt

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}
