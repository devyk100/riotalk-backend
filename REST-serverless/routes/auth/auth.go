package auth_route

import (
	"REST-serverless/routes/auth/email"
	"REST-serverless/routes/auth/google"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRouter(router *gin.RouterGroup) *gin.RouterGroup {
	authRouter := router.Group("/auth")
	google.GoogleAuthRouter(authRouter)
	email.EmailAuthRouter(authRouter)
	authRouter.GET("/refresh-token", RefreshToken())
	authRouter.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("refresh_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No refresh token"})
		} else {
			c.String(http.StatusOK, cookie)
		}
	})
	return authRouter
}
