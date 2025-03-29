package auth_route

import (
	"REST-serverless/routes/auth/google"
	"REST-serverless/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshToken() gin.HandlerFunc {
	/*
		1. Get the refresh token --> extract refresh token, and type
		2. Get the new access token if google or generate new if email
		3. Wrap it around in custom payload
	*/
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh_token")
		fmt.Println(refreshToken, "IS THE REFRESH TOKEN")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No refresh token"})
			return
		}
		token, method, err := utils.ParseToken(refreshToken)
		if err != nil {
			fmt.Println("")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Some JWT parsing error " + err.Error()})
			return
		}

		if method == "google" {
			google.RefreshTokenGoogle(token, c)
		} else {

		}

	}
}
