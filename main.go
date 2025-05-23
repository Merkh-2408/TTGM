package main

import (
	"TTGM/modsls"
	"TTGM/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建gin实例
	r := gin.Default()

	// 初始化数据库
	if err := modsls.InitDB(); err != nil {
		panic("Failed to initialize database: " + err.Error())
	}

	// 静态文件服务已在routers.Routes中配置

	// 设置路由
	routers.Routes(r)

	// 启动服务器
	r.Run(":8080")
}
