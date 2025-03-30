package servers_route

import (
	"REST-serverless/db"
	"REST-serverless/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

type CreateInviteRequest struct {
	ExpiryTime *int64 `json:"expiry_time"`
	Uses       *int32 `json:"uses"`
	ServerId   int64  `json:"server_id"`
}

func CreateInvite() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqPayload CreateInviteRequest
		err := c.ShouldBindJSON(&reqPayload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		inviteId := utils.RandomString(10)
		var uses *pgtype.Int4
		var expiry *pgtype.Int8
		if reqPayload.Uses != nil {
			uses = &pgtype.Int4{Valid: true, Int32: *reqPayload.Uses}
		} else {
			uses = &pgtype.Int4{Valid: false} // Marks as NULL in DB
		}
		if reqPayload.ExpiryTime != nil {
			expiry = &pgtype.Int8{
				Int64: *reqPayload.ExpiryTime,
				Valid: true,
			}
		} else {
			expiry = &pgtype.Int8{Valid: false}
		}
		if uses != nil || expiry != nil {
			userId, err := utils.ExtractUserIDFromContext(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			invite, err := db.DBQueries.CreateServerInvite(c.Request.Context(), db.CreateServerInviteParams{
				ID:       inviteId,
				ServerID: reqPayload.ServerId,
				Expiry:   *expiry,
				Uses:     *uses,
				UserID:   userId,
				Valid:    true,
			})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, gin.H{"id": invite.ID})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Either uses, or "})
		}
	}
}
