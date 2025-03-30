package email

import (
	"REST-serverless/db"
	"REST-serverless/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmailLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func EmailLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			Take the username/email and password, and verify in the DB
			Make the refresh token, and send it, and redirect to "auth-success"
		*/
		var reqPayload EmailLoginRequest
		err := c.ShouldBindJSON(&reqPayload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := db.DBQueries.GetPasswordFromUserNameEmail(c.Request.Context(), reqPayload.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if user.Password.Valid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "If signed in from Google, you must set the password"})
			return
		}
		valid := utils.ComparePasswordHash(reqPayload.Password, user.Password.String)
		if valid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}
		token := utils.CreateRefreshToken("email", "email", user.ID)
		fmt.Println(token, "is the token, generated")
		c.SetCookie("refresh_token", token, 60*60*24, "/", "", false, true)
		c.Redirect(http.StatusFound, "http://localhost:3000/auth-success")
	}
}
