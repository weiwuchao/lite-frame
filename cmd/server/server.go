package server

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"lite-frame/pkg/router"
	"os"
)

type SeverConfig struct {
	Port    int
	LogPath string
}

type HttpServer struct {
	Config *SeverConfig
}

func NewServer() *HttpServer {
	return &HttpServer{
		Config: &SeverConfig{},
	}
}

func (h *HttpServer) InitCommand() {
	flag.IntVar(&h.Config.Port, "port", 8080, "insecure http port")
	flag.StringVar(&h.Config.LogPath, "log-path", "gin.log", "log path")
}

func (h *HttpServer) InitLog() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()
	f, err := os.Create(h.Config.LogPath)
	if err != nil {
		fmt.Printf("create gin.log error:%v\n", err)
		return
	}
	// 需要同时将日志写入文件和控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func (h *HttpServer) Run() {
	// 创建Gin的默认引擎
	r := gin.Default()
	g := r.Group("/api/lite-frame/v1")
	router.InitRouter(g)
	err := r.Run(fmt.Sprintf(":%d", h.Config.Port))
	if err != nil {
		return
	}
}
