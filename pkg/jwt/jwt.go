package jwt

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"kappa-web/config"
)

const Key = "token"

type Claims struct {
	Account string      `json:"account"`
	Role string         `json:"role"`
	jwt.StandardClaims
}

func VerifyToken(token string) (account string, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Val.JWTSecret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok || !tokenClaims.Valid {
		return "", errors.New("tokenClaims invalid")
	}

	return claims.Account, nil
}
