package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type CustomTokenClaims struct {
	Token  string `json:"token"`
	Method string `json:"method"`
	UserID int64  `json:"userId"`
	jwt.RegisteredClaims
}

var SECRET_KEY = []byte("SOMETIME")

func ParseToken(tokenString string) (string, string, int64, error) {
	// Parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure it uses HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return SECRET_KEY, nil
	})
	if err != nil {
		return "", "", -1, fmt.Errorf("failed to parse token: %w", err)
	}

	// Extract claims
	claims, ok := token.Claims.(*CustomTokenClaims)
	if !ok || !token.Valid {
		return "", "", -1, fmt.Errorf("invalid token")
	}

	return claims.Token, claims.Method, claims.UserID, nil
}
