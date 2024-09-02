package main

import (
	"fc-gin/pkg/conf"
	"fc-gin/router"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := router.NewRouter()
	router.Run(os.Getenv("APP_HOST"))
}
