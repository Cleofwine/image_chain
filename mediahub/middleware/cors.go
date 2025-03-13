package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	// 添加跨域逻辑
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders: []string{
			"Origin", "Content-Length", "Content-Type",
		},
		AllowMethods: []string{
			"GET", "POST", "DELETE", "PUT", "HEAD", "OPTIONS",
		},
	})
}
