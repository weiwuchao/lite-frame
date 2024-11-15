package router

import (
	"github.com/gin-gonic/gin"
	"lite-frame/pkg/middleware"
	"lite-frame/pkg/router/user"
)

func InitRouter(group *gin.RouterGroup) {
	// 添加中间件
	group.Use(middleware.Auth)

	// 注册用户路由
	user.RegisterUser(group)
}
