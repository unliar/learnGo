package main

import (
	"github.com/gin-gonic/gin"

	"learnGo/examples/go-curl/conf"
	"learnGo/examples/go-curl/router"
)

func main() {
	gin.SetMode(conf.CurrentConf.Mode)
	r := router.InitRoute()

	r.Run(conf.CurrentConf.Port)
}
