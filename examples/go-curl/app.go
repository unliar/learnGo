package main

import (
	"github.com/gin-gonic/gin"

	"learnGo/examples/go-curl/conf"
	"learnGo/examples/go-curl/router"
)

func main() {
	mode := conf.CurrentConf.Mode
	// 设置mode
	gin.SetMode(mode)
	// 非debug模式不打印彩色
	if mode != "debug" {
		gin.DisableConsoleColor()
	}
	r := router.InitRoute()
	r.Run(conf.CurrentConf.Port)
}
