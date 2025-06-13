package route

import (
	"gin-web-server/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 版本路由
	apiV1 := r.Group("/v1")

	// 用户资源路由组
	users := apiV1.Group("/user")
	{
		users.POST("/register", controller.Register) // 用户注册
		users.POST("/login", controller.Login)       // 用户登录
	}

	return r
}
