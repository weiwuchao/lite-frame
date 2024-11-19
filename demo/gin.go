package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	v1 "lite-frame/apis/v1"
	"net/http"
	"os"
	"reflect"
	"time"
)

func testGin() {

	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件。
	f, err := os.Create("gin.log")
	if err != nil {
		fmt.Printf("create gin.log error:%v\n", err)
		return
	}
	//gin.DefaultWriter = f
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 创建Gin的默认引擎
	r := gin.Default()
	r.Use(m1)
	r.With()

	// 定义一个GET请求的处理器函数
	r.GET("/hello", func(c *gin.Context) {
		fmt.Println("hello gin")
		time.Sleep(5 * time.Second)
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	}, m1)

	r.LoadHTMLGlob("./template/*")
	r.GET("/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "admin",
			"pwd":  "123456",
		})
	})

	r.GET("/param/id/:id", func(c *gin.Context) {
		name := c.Query("name")
		pwd := c.Query("pwd")
		id := c.Param("id")
		fmt.Printf("params=%#v\n", c.Params)
		fmt.Printf("id=%s", id)
		// fmt.Printf("name:%s ; pwd:%s",name,pwd)
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"pwd":  pwd,
		})
	})

	r.POST("/user", func(c *gin.Context) {
		var u v1.User
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": GetValidMsg(err, &u),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
		fmt.Printf("%#v\n", u)
	})

	r.GET("/gournate", func(c *gin.Context) {

		cp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println("path=" + cp.Request.URL.Path)
			fmt.Println("gournate")
		}()
		fmt.Println("gournate222")
	})

	r.GET("/redirect", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/hello"
		r.HandleContext(c)
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "templates/404.html", nil)
	})

	// 启动HTTP服务器，默认监听在本地的8080端口
	r.Run(":8090")
}

func GetValidMsg(err error, obj any) string {
	// 使用的时候，需要传obj的指针
	getObj := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}

	return err.Error()
}

func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	start := time.Now()
	//c.Next() //调用后续的处理函数
	c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
}
