package user

import (
	"github.com/gin-gonic/gin"
	"lite-frame/pkg/controller/user"
)

func RegisterUser(group *gin.RouterGroup) {
	group.GET("/users", user.ListUser)
	group.GET("/users/:userId", user.GetUserById)
	group.POST("/users", user.Create)
}
