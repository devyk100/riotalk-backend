package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = "SSOMEITG"

func CreateRefreshToken(method *string, refreshToken *string) string {
	claims := &jwt.MapClaims{
		"token":  refreshToken,
		"method": method,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return signedToken
}

func CreateAccessToken(method *string, accessToken *string) (string, error) {
	claims := &jwt.MapClaims{
		"token":  accessToken,
		"method": method,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return signedToken, err
}
