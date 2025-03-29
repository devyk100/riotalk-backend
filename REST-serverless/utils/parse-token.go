package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString string) (string, string, error) {
	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure it uses HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return SECRET_KEY, nil
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to parse token: %w", err)
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", fmt.Errorf("invalid token")
	}

	// Convert values safely
	accessToken, ok1 := claims["token"].(string)
	method, ok2 := claims["method"].(string)

	if !ok1 || !ok2 {
		return "", "", fmt.Errorf("malformed token claims")
	}

	return accessToken, method, nil
}
