package main

import (
	"REST-serverless/db"
	"REST-serverless/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
)

// REFER: https://gin-gonic.com/docs/

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(db.DBQueries.CreateUser(c.Request.Context(), db.CreateUserParams{
			Name:        "Yash",
			Username:    "Usernameseomt",
			Email:       "yashkumar@gmail.com",
			Img:         pgtype.Text{String: "fmoieaw"},
			Description: pgtype.Text{String: "foinwaelkf"},
		}))
		c.JSON(200, gin.H{
			"Success": "Yes dude!",
		})
	}
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	// all routes after this use DB, so initialisation is needed
	r.Use(middleware.InitDBMiddleware())
	r.GET("/hello", Test())
	r.Run(":8080") // Local development
}
