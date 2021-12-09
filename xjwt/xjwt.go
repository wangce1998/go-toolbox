package xjwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const Secret = "wangce1998"

type JWTClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}

func New(userID int64) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
		UserID: userID,
	})
}
