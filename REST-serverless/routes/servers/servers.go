package servers_route

import (
	"REST-serverless/middleware"
	"github.com/gin-gonic/gin"
)

func ServersRouter(router *gin.RouterGroup) *gin.RouterGroup {
	serverRouter := router.Group("/servers")
	serverRouter.GET("/list", middleware.AuthMiddleware(), ListAllServers())
	serverRouter.POST("/create", middleware.AuthMiddleware(), CreateServer())
	serverRouter.GET("/accept-invite", middleware.AuthMiddleware(), AcceptInvite())
	serverRouter.POST("/create-invite", middleware.AuthMiddleware(), CreateInvite())
	serverRouter.PUT("/change-role", middleware.AuthMiddleware(), ChangeUserRole())
	return serverRouter
}
