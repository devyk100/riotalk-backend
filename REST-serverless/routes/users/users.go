package users_route

import (
	"REST-serverless/middleware"
	"github.com/gin-gonic/gin"
)

func SetUsername() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUserToUserChats() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UsersRouter(router *gin.RouterGroup) *gin.RouterGroup {
	usersRouter := router.Group("/users")
	usersRouter.POST("/create", CreateUserFromEmail())
	// pass through the middleware
	usersRouter.GET("/info", middleware.AuthMiddleware(), GetUserInfo())
	usersRouter.PUT("/edit", middleware.AuthMiddleware(), EditUser())
	return usersRouter
}
