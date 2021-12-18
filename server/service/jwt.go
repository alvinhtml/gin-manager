package service

import (
	"errors"
	"time"

	"github.com/alvinhtml/gin-manager/server/global"
	"github.com/dgrijalva/jwt-go"
)

type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// @title VerifyToken
// @description Verifies a JWT token
// @return    err             error
func VerifyToken(token string) (err error) {
	jwtToken, err := jwt.ParseWithClaims(
		string(token),
		&customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(global.CONFIG.JWT.SigningKey), nil
		},
	)

	if err != nil {
		return err
	}

	claims, ok := jwtToken.Claims.(*customClaims)
	if !ok {
		return errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return errors.New("jwt is expired")
	}
	return nil
}

// @title CreateToken
// @description Creates a JWT token
// @return    token           string
// @return    err             error
func CreateToken(username string) (signedToken string, err error) {
	claims := customClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: global.CONFIG.JWT.ExpiresTime,
			Issuer:    global.CONFIG.JWT.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(global.CONFIG.JWT.SigningKey))

	return signedToken, err
}
