package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-web"
	ac "learnGo/examples/go-micro-api/account"
	"log"
)

func main() {
	var err error

	service := web.NewService(
		web.Name("unliar-restful-api"),
		// 端口设置 否则是随机端口
		web.Address(":8088"),
		web.Version("beta-1.1.1"),
	)

	service.Init()

	r := gin.Default()
	r.GET("/api/health", ac.GetHealthStatus)
	r.GET("/api/users/:uid", ac.GetUserBase)

	service.Handle("/", r)

	err = service.Run()

	if err != nil {
		log.Fatal(err)
	}
}
