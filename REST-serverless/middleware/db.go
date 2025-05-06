package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shared/db"
)

func InitDBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if db.DBQueries == nil {
			err := db.InitDb(c.Request.Context())
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		c.Next()
	}
}
