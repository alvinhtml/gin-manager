package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/alvinhtml/gin-manager/server/global"
	"github.com/alvinhtml/gin-manager/server/model"
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

	// 打印 err
	fmt.Println(err)

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
func CreateToken(username string) (token model.Jwt, err error) {
	token.ExpiresAt = time.Now().Add(time.Duration(global.CONFIG.JWT.ExpiresTime) * time.Second)

	claims := customClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: token.ExpiresAt.Unix(),
			Issuer:    global.CONFIG.JWT.Issuer,
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := t.SignedString([]byte(global.CONFIG.JWT.SigningKey))

	if err != nil {
		return model.Jwt{}, err
	}

	token.Token = signedToken

	// 当前时间加上过期时间

	err = global.DB.Create(&token).Error

	fmt.Printf(signedToken)

	return token, nil
}

// @title ParseToken
// @description Parses a JWT token
// @return    token           string
// @return    err             error
func ParseToken(token string) (username string, err error) {
	jwtToken, err := jwt.ParseWithClaims(
		string(token),
		&customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(global.CONFIG.JWT.SigningKey), nil
		},
	)

	if err != nil {
		return "", err
	}

	claims, ok := jwtToken.Claims.(*customClaims)
	if !ok {
		return "", errors.New("couldn't parse claims")
	}

	return claims.Username, nil
}
