package email

import "github.com/gin-gonic/gin"

func EmailAuthRouter(router *gin.RouterGroup) *gin.RouterGroup {
	emailAuthRouter := router.Group("/email")
	emailAuthRouter.POST("/login", EmailLogin())
	return emailAuthRouter
}
