package middleware

import (
	"REST-serverless/db"
	"github.com/gin-gonic/gin"
)

func InitDBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if db.DBQueries == nil {
			err := db.InitDb(c.Request.Context())
			if err != nil {
				panic(err)
			}
		}
		c.Next()
	}
}
