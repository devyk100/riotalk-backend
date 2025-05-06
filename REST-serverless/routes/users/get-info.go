package users_route

import (
	"REST-serverless/routes/auth/google"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/db"
)

type UserInfo struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Image    string `json:"image"`
}

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, exists := c.Get("token")
		if !exists {
			fmt.Println(token, "IS THE TOKEN FOR GETTING USER DATA")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - The token was not passed from middleware"})
			return
		}
		method, exists := c.Get("method")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - The method was not passed from middleware"})
			return
		}

		tokenStr, ok := token.(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - The token is not a string"})
			return
		}

		if method == "google" {
			user, err := google.FetchGoogleUserData(&tokenStr)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			User, err := db.DBQueries.GetUserByEmail(c.Request.Context(), user.Email)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"name":     User.Name,
				"email":    User.Email,
				"img":      User.Img,
				"user_id":  User.ID,
				"username": User.Username,
			})
		} else {
			userId, exists := c.Get("userId")
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Authentication error, the userId field is missing"})
			}
			User, err := db.DBQueries.GetUserById(c.Request.Context(), userId.(int64))
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			}

			c.JSON(http.StatusOK, gin.H{
				"name":     User.Name,
				"email":    User.Email,
				"img":      User.Img,
				"user_id":  User.ID,
				"username": User.Username,
			})
		}
	}
}
