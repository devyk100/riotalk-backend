package servers_route

import (
	"REST-serverless/db"
	"REST-serverless/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ChangeUserRoleRequest struct {
	UserId int64  `json:"user_id"`
	Role   string `json:"role"`
}

func ChangeUserRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request ChangeUserRoleRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		initiatorUserId, err := utils.ExtractUserIDFromContext(c)
		err = db.DBQueries.UpdateUserRole(c.Request.Context(), db.UpdateUserRoleParams{
			UserID:   initiatorUserId,
			UserID_2: request.UserId,
			Role:     db.UserRole(request.Role),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}
