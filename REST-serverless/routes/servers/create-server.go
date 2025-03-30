package servers_route

import (
	"REST-serverless/db"
	"REST-serverless/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

type CreateServerRequest struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	Img         string `json:"img"`
	Banner      string `json:"banner"`
}

/*
Requires the auth middleware
*/
func CreateServer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqPayload CreateServerRequest
		if err := c.ShouldBindJSON(&reqPayload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userId, err := utils.ExtractUserIDFromContext(c)
		serverId, err := db.DBQueries.CreateServerAndMapping(c.Request.Context(), db.CreateServerAndMappingParams{
			Name: reqPayload.Name,
			Description: pgtype.Text{
				String: reqPayload.Description,
				Valid:  true,
			},
			Img: pgtype.Text{
				String: reqPayload.Img,
				Valid:  true,
			},
			Banner: pgtype.Text{
				String: reqPayload.Banner,
				Valid:  true,
			},
			UserID: userId,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusCreated, gin.H{"id": serverId, "success": true})
		}
	}
}
