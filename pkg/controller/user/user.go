package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "lite-frame/apis/v1"
	"lite-frame/pkg/common"
	"lite-frame/pkg/service"
)

func ListUser(c *gin.Context) {
	name := c.Query("name")
	selector := v1.UserSelector{Name: name}
	page := common.BuildPageParams(c)
	users, count, err := service.ListUser(selector, page)
	rb, _ := json.Marshal(users)
	c.JSON(200, gin.H{
		"items": string(rb),
		"count": count,
		"err":   err.Error(),
	})
}

func GetUserById(c *gin.Context) {
	userId := c.Param("userId")
	userName := c.Query("userName")
	c.JSON(200, gin.H{
		"message":  "Hello, Gin!",
		"userId":   userId,
		"userName": userName,
	})
}

func Create(c *gin.Context) {
	body := common.ReadRequestBody(c.Request)
	user := v1.User{}
	if err := json.Unmarshal(body, &user); err != nil {
		c.JSON(500, gin.H{"message": "Create User failed!"})
		return
	}
	err := service.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"message": "Create User failed!"})
		return
	}
	c.JSON(200, gin.H{
		"message": "Create User Success!",
	})
}
