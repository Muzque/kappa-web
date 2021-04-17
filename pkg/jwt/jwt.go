package jwt

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"kappa-web/config"
	"time"
)

const Key = "token"

type Credentials struct {
	Account string `json:"account"`
	Password string `json:"password"`
}

type Claims struct {
	Account string      `json:"account"`
	Role string         `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(account, password string) (string, error){
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.Val.JWTTokenTtl) * time.Second)

	if account != "mario" || password != "35" {
		return "", errors.New("Account/Password is not correct")
	}

	claims := Claims{
		account,
		"member",
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "Kappa-Web",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(config.Val.JWTSecret))
	return token, err
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
