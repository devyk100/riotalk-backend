package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/utils"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No token provided"})
			c.Abort()
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		token, method, userId, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid token"})
			c.Abort()
			return
		}
		// Do not forget that this value has to be used everywhere
		c.Set("token", token)
		c.Set("method", method)
		c.Set("userId", userId)
		c.Next()
	}
}
