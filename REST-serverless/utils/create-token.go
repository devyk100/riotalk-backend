package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var SECRET_KEY = []byte("SOMETIME")

func CreateRefreshToken(method string, refreshToken string, userId int64) string {
	claims := CustomTokenClaims{
		Token:  refreshToken,
		Method: method,
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 10)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return signedToken
}

func CreateAccessToken(method string, accessToken string, userId int64) (string, error) {
	claims := CustomTokenClaims{
		Token:  accessToken,
		Method: method,
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return signedToken, err
}
