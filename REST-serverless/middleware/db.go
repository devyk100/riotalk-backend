package middleware

import (
	"REST-serverless/db"
	"fmt"
	"github.com/gin-gonic/gin"
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
