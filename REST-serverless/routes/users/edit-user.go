package users_route

import (
	"REST-serverless/db"
	"REST-serverless/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EditUserRequest struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Img         string `json:"img"`
	Description string `json:"description"`
}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqPayload EditUserRequest
		err := c.ShouldBindJSON(&reqPayload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID, err := utils.ExtractUserIDFromContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = db.DBQueries.UpdateUser(c.Request.Context(), db.UpdateUserParams{
			ID:      userID,
			Column2: reqPayload.Name,
			Column3: reqPayload.Username,
			Column4: reqPayload.Img,
			Column5: reqPayload.Description,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}
