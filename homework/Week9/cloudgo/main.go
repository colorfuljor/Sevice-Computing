package main

import (
	"os"

	"github.com/github-user/cloudgo/service"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8080"
)

func main() {
	//获取自定义的环境变量PORT作为端口值
	port := os.Getenv("PORT")

	//设置默认端口值
	if len(port) == 0 {
		port = PORT
	}

	//获取命令行参数-p作为端口值
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	//启动服务器
	service.Run(port)
}
