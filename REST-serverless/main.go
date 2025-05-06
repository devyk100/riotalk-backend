package main

import (
	"REST-serverless/db"
	"REST-serverless/redis"
	"REST-serverless/routes"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

// REFER: https://github.com/awslabs/aws-lambda-go-api-proxy

var ginLambda *ginadapter.GinLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	db.InitDb(context.Background())
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

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

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

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
