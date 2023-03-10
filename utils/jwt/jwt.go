package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaim struct {
	Identity int64  `json:"user_id"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

var key = []byte("helloWord")

func GenerateToken(identity int64, name string) (string, error) {
	uc := UserClaim{
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(time.Hour)).Unix(), //设置过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)

	return token.SignedString(key)

}

func AnalyzeToken(tokenStr string) (*UserClaim, error) {
	uc := new(UserClaim)
	token, err := jwt.ParseWithClaims(tokenStr, uc, func(token *jwt.Token) (i interface{}, err error) {
		return key, nil
	})
	if err != nil {
		return nil, errors.New("analyze token error")
	}
	if token.Valid { // 校验token
		return uc, nil
	}
	return nil, errors.New("invalid token")
}
