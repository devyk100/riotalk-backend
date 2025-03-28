package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type TokenPayload struct {
	Token  string `json:"token"`
	Method string `json:"method"`
}

func ParseToken(tokenString string) (*TokenPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return SECRET_KEY, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &TokenPayload{
			Token:  claims["token"].(string),
			Method: claims["method"].(string),
		}, nil
	}
	return nil, err
}
