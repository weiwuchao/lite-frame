package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Auth(c *gin.Context) {

	fmt.Println("enter Auth in...")
	start := time.Now()
	c.Next() //调用后续的处理函数
	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)

	fmt.Printf("request url is %s\n", c.Request.URL)

}
