package route

import (
	"server/controller"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// API v1 路由分组
	apiV1 := r.Group("/v1")

	// 用户资源路由组
	users := apiV1.Group("/users")
	{
		users.POST("/register", controller.Register)               // 用户注册
		users.POST("/login", controller.Login)                     // 用户登录
		users.GET("/me", middleware.JWTAuth(), controller.Profile) // 当前用户信息
	}

	return r
}
