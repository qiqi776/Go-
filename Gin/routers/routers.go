package routers

import (
	"GoStudy/handler"
	"GoStudy/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 2. 定义一个路由组，前缀是 /api/v1
	//    这是一个 "公开" 组，任何人都可以访问
	publicGroup := r.Group("/api/v1")
	{
		// "雇佣" auth_handler.go 里的 LoginHandler
		// 负责处理 /api/v1/login 路径的 POST 请求
		publicGroup.POST("/login", handler.LoginHandler)
	}

	// 3. 定义另一个路由组，前缀也是 /api/v1
	//    这是一个 "私有" 组，必须通过认证
	privateGroup := r.Group("/api/v1")

	// ！！！关键： "雇佣" jwt_auth.go 里的 JWTAuthMiddleware
	//    告诉 Gin：这个组 (privateGroup) 里的所有路由，
	//    都必须先经过 JWTAuthMiddleware 这个 "保安" 的检查。
	privateGroup.Use(middleware.JWTAuthMiddleware())
	{
		// "雇佣" hello_handler.go 里的 HelloHandler
		// 负责处理 /api/v1/hello 路径的 GET 请求
		// (只有通过了保安检查的请求才能到达这里)
		privateGroup.GET("/hello", handler.HelloHandler)
	}
	return r
}
