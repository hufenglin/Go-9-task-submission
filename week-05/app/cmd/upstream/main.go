package main

import (
	"hystrix-demo/app/server"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	server.NewUpStreamServer(
		10,
		50,            //熔断阈值，总请求数为50
		0.8,           //总失败率为80%
		time.Second*6, //熔断生效时长6秒
	).Run(":9000")
}
