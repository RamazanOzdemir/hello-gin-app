package handlers

import (
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.String(200, "Hello,Gin World!!!")
}