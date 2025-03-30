package channels_route

import (
	"REST-serverless/db"
	"REST-serverless/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

type CreateChannelRequest struct {
	Name         string  `json:"name"`
	Type         *string `json:"type"`
	ServerId     int64   `json:"server_id"`
	AllowedRoles string  `json:"allowed_roles"`
	Description  string  `json:"description"`
}

func CreateChannel() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqPayload CreateChannelRequest
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
		channelType := ""
		if reqPayload.Type == nil {
			channelType = "chat"
		} else {
			channelType = *reqPayload.Type
		}
		channel, err := db.DBQueries.CreateChannelIfAuthorized(c.Request.Context(), db.CreateChannelIfAuthorizedParams{
			Name:         reqPayload.Name,
			Type:         db.ChannelType(channelType),
			ServerID:     reqPayload.ServerId,
			AllowedRoles: db.UserRole(reqPayload.AllowedRoles),
			Description: pgtype.Text{
				String: reqPayload.Description,
				Valid:  true,
			},
			UserID: userID,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": channel.ID, "success": true})
	}
}
