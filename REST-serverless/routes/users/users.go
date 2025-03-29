package users_route

import (
	"REST-serverless/middleware"
	"github.com/gin-gonic/gin"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func IsUsernameAvailable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

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
	// pass through the middleware
	usersRouter.GET("/info", middleware.AuthMiddleware(), GetUserInfo())
	return usersRouter
}
