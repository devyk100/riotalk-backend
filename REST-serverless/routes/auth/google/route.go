package google

import (
	"REST-serverless/middleware"
	"github.com/gin-gonic/gin"
)

func GoogleAuthRouter(router *gin.RouterGroup) *gin.RouterGroup {
	googleAuthRouter := router.Group("/google")
	googleAuthRouter.GET("/callback", middleware.InitDBMiddleware(), GoogleCallback())
	googleAuthRouter.GET("/get-oauth-url", GetOauthURL())
	googleAuthRouter.GET("/initiate", InitiateGoogleAuth())
	return googleAuthRouter
}
