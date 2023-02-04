package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/riicarus/loveshop/conf"
)

type AuthClaims struct {
	Account   string `json:"account"`
	LoginType string `json:"loginType"`
	jwt.StandardClaims
}

func GenToken(account, loginType string) (string, error) {
	c := AuthClaims{
		account,
		loginType,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(conf.ServiceConf.Jwt.Expire)).Unix(),
			Issuer: conf.ServiceConf.Jwt.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString([]byte(conf.ServiceConf.Jwt.Secret))
}