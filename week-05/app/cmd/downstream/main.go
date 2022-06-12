package main

import (
	"hystrix-demo/app/server"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	server.NewDownStreamServer(0.2).Run(":8000")
}
