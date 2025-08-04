package main

import (
	"github.com/ramazanozdemir/hello-gin-app/router"
)


func main() {
	r := router.SetupRouter()

	r.Run(":8080") // Start the server on port 8080
}