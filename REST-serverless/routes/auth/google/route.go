package google

import "github.com/gin-gonic/gin"

func GoogleAuthRouter(router *gin.RouterGroup) *gin.RouterGroup {
	googleAuthRouter := router.Group("/google")
	googleAuthRouter.GET("/callback", GoogleCallback())
	googleAuthRouter.GET("/get-oauth-url", GetOauthURL())
	googleAuthRouter.GET("/initiate", InitiateGoogleAuth())
	return googleAuthRouter
}
