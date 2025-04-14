package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置中间件
func Cors() gin.HandlerFunc {
	// 配置 CORS 中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // 允许所有来源的跨域请求
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	config.AllowCredentials = true
	return cors.New(config)
}
