package routers

import (
	"TTGM/control"

	"github.com/gin-gonic/gin"
)

// Routes 配置所有路由
func Routes(r *gin.Engine) {
	// 配置静态文件服务
	r.Static("/static", "./front")

	// 配置默认路由，重定向到前端页面
	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/static/index.html")
	})

	// API路由
	r.GET("/random-article", control.RandomArticle)

	// 配置CORS中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
}
