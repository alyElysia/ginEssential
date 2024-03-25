package routes

import (
	"ginEssential/controller"
	"ginEssential/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.GET("/api/auth/login", controller.Login)
	//使用中间件保护用户信息
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return r
}
