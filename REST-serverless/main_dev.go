package main

import (
	"REST-serverless/db"
	"REST-serverless/redis"
	"REST-serverless/routes"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"time"
)

// REFER: https://gin-gonic.com/docs/

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Println(db.DBQueries.CreateUser(c.Request.Context(), db.CreateUserParams{
		//	Name:        "Yash",
		//	Username:    "Usernameseomt",
		//	Email:       "yashkumar@gmail.com",
		//	Img:         pgtype.Text{String: "fmoieaw"},
		//	Description: pgtype.Text{String: "foinwaelkf"},
		//}))
		//c.JSON(200, gin.H{
		//	"Success": "Yes dude!",
		//})
	}
}

func main2() {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow frontend domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Important: allows cookies
		MaxAge:           12 * time.Hour,
	}))

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file", err.Error())
	}

	err = redis.InitRedisClient(context.Background())
	if err != nil {
		panic(err)
		return
	}

	// all routes after this use DB, so initialisation is needed
	err = db.InitDb(context.Background())
	if err != nil {
		fmt.Println("Error initializing db", err.Error())
		return
	}
	//r.Use(middleware.InitDBMiddleware())
	routes.RoutesRouter(r)
	err = r.Run(":8080")
	if err != nil {
		fmt.Println("At r.Run", err.Error())
		return
	}
}
