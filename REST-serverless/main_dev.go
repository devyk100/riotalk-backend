package main

import "github.com/gin-gonic/gin"

// REFER: https://gin-gonic.com/docs/

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, AWS Lambda!")
	})

	r.Run(":8080") // Local development
}
