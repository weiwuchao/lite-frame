package main

import (
	"flag"
	"lite-frame/cmd/server"
	"lite-frame/pkg/dao/factory"
)

func main() {

	// 初始化命令行参数
	s := server.NewServer()
	s.InitCommand()
	flag.Parse()

	//设置日志路径
	s.InitLog()

	//初始化数据库
	factory.InitDB()

	//启动rest服务
	s.Run()
}
