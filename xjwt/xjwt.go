package xjwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/wangce1998/go-toolbox/xerror"
	"time"
)

const Secret = "wangce1998"

var (
	TokenNotfound = xerror.New(11001, "未携带令牌")
	TokenInvalid  = xerror.New(11002, "%s")
	TokenExpired  = xerror.New(11003, "令牌已过期")
)

type JWTClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}

func NewToken(userID int64) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
		UserID: userID,
	})
}

func TokenToString(token *jwt.Token) (string, error) {
	return token.SignedString([]byte(Secret))
}

func Parse(tokenStr string) (JWTClaims, xerror.XError) {
	var jc JWTClaims
	if len(tokenStr) == 0 {
		return jc, TokenNotfound
	}
	token, err := jwt.ParseWithClaims(tokenStr, &jc, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return jc, TokenExpired
			} else {
				return jc, xerror.New(TokenInvalid.Code(), err.Error())
			}
		}
	}
	if !token.Valid {
		return jc, xerror.New(TokenInvalid.Code(), "令牌无效")
	}

	return jc, nil
}
