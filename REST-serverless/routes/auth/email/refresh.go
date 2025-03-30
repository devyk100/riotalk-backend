package email

import (
	"REST-serverless/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshTokenEmail(refreshToken string, userId int64, c *gin.Context) {
	method := "email"
	token, err := utils.CreateAccessToken(method, refreshToken, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make the access token"})
		return
	}
	fmt.Println("It reached here")
	c.JSON(http.StatusOK, gin.H{"accessToken": token})
}
