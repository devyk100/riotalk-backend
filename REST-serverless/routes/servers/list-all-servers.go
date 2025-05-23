package servers_route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/db"
	"shared/utils"
)

func ListAllServers() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := utils.ExtractUserIDFromContext(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		list, err := db.DBQueries.GetServersList(c.Request.Context(), userId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, list)
	}
}
