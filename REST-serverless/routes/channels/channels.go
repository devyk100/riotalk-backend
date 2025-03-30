package channels_route

import (
	"REST-serverless/middleware"
	"github.com/gin-gonic/gin"
)

func GetChannelChats() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ChannelRouter(router *gin.RouterGroup) *gin.RouterGroup {
	channelRouter := router.Group("/channels")
	channelRouter.GET("/list", middleware.AuthMiddleware(), ListAllChannels())
	channelRouter.POST("/create", middleware.AuthMiddleware(), CreateChannel())
	channelRouter.PUT("/edit", middleware.AuthMiddleware(), EditChannel())
	return channelRouter
}
