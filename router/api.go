package router

import (
	"github.com/gin-gonic/gin"
	api "go_fc/controller"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	route := gin.Default()

	// 路由
	fc := route.Group("/")
	{
		// 默认入口
		fc.POST("invoke", api.FcInvoke)

		// 初始化
		fc.POST("initialize", api.FcInitialize)

		// 获取headers数据
		fc.POST("header/:header", api.FcHeader)
	}
	return route
}
