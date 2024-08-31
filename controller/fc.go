package api

import (
	"github.com/gin-gonic/gin"
	"go_fc/pkg/serializer"
)

// 函数计算调用入口
func FcInvoke(c *gin.Context) {

	c.JSON(200, serializer.Response{
		Code:    0,
		Message: "Pong",
	})
}

// 函数计算初始化
func FcInitialize(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code:    0,
		Message: "Pong",
	})
}

// 函数计算Header头
func FcHeader(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code:    0,
		Message: "Pong",
	})
}
