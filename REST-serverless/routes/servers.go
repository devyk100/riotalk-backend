package routes

import "github.com/gin-gonic/gin"

func ListAllServers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateServer() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func EditServer() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ServerRouter(router *gin.Engine) *gin.RouterGroup {
	serverRouter := router.Group("/servers")
	serverRouter.GET("/list", ListAllServers())
	return serverRouter
}
