package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin World!")
	})

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080") // Start the server on port 8080
}