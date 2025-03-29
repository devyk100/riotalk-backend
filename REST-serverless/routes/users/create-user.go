package users_route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserRequest struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Img         string `json:"img"`
	Description string `json:"description"`
	Provider    string `json:"provider"`
}

func CreateUserFromEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqPayload CreateUserRequest
		err := c.ShouldBindJSON(&reqPayload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
