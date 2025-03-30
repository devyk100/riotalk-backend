package channels_route

import (
	"REST-serverless/db"
	"REST-serverless/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EditChannelRequest struct {
	ChannelID    int64  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	AllowedRoles string `json:"allowed_roles"`
}

func EditChannel() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqPayload EditChannelRequest
		err := c.ShouldBindJSON(&reqPayload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userId, err := utils.ExtractUserIDFromContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = db.DBQueries.UpdateChannel(c.Request.Context(), db.UpdateChannelParams{
			ID:      reqPayload.ChannelID,
			UserID:  userId,
			Column3: reqPayload.Name,
			Column4: reqPayload.AllowedRoles,
			Column5: reqPayload.Description,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}
