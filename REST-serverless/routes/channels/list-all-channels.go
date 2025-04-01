package channels_route

import (
	"REST-serverless/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListAllChannels() gin.HandlerFunc {
	return func(c *gin.Context) {
		server_id := c.Query("server_id")
		//channelType := c.Param("type") --> ADD THIS FILTER IN THE FUTURE
		userId, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized no userId exists"})
		}
		userIdInt64, ok := userId.(int64)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User Id is not a number"})
		}
		serverIdInt, err := strconv.Atoi(server_id)
		serverIdInt64 := int64(serverIdInt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		list, err := db.DBQueries.GetChannelList(c.Request.Context(), db.GetChannelListParams{
			UserID:   userIdInt64,
			ServerID: serverIdInt64,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, list)
	}
}
