package middleware

import (
	"REST-serverless/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func validateToken(token string) bool {
	return true
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No token provided"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
