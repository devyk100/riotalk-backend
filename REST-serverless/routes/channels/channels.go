package channels_route

import "github.com/gin-gonic/gin"

func ListAllChannels() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateChannel() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func EditChannel() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetChannelChats() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ChannelRouter(router *gin.RouterGroup) *gin.RouterGroup {
	channelRouter := router.Group("/channels")
	channelRouter.GET("/list")
	return channelRouter
}
