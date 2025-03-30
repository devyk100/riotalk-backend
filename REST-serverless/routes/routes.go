package routes

import (
	"REST-serverless/middleware"
	auth_route "REST-serverless/routes/auth"
	channels_route "REST-serverless/routes/channels"
	servers_route "REST-serverless/routes/servers"
	users_route "REST-serverless/routes/users"
	"github.com/gin-gonic/gin"
)

func RoutesRouter(router *gin.Engine) *gin.RouterGroup {
	routesRouter := router.Group("/")
	routesRouter.Use(middleware.InitDBMiddleware())
	auth_route.AuthRouter(routesRouter)
	users_route.UsersRouter(routesRouter)
	servers_route.ServersRouter(routesRouter)
	channels_route.ChannelRouter(routesRouter)
	return routesRouter
}
