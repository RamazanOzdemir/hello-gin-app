package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ramazanozdemir/hello-gin-app/handlers"
)

// SetupRouter initializes the Gin router with the necessary routes.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define routers 
	r.GET("/hello", handlers.HelloHandler)
	
	return r
}