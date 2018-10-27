package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-web"
	"learnGo/examples/go-micro-api/routes"
	"log"
	"time"
)

// ErrorMsg 错误信息
type ErrorMsg struct {
	Code    int64  `json:"code"`
	Detail  string `json:"detail"`
	Time    int64  `json:"time"`
	Message string `json:"message"`
}

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
	// 404 error
	r.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{
			"statusCode": 404,
			"data":       nil,
			"error": ErrorMsg{
				Code:    999,
				Detail:  "what'up? you may have the wrong api path!",
				Time:    time.Now().Unix(),
				Message: "failed",
			},
		})
	})

	// 添加账户模块
	routes.AccountRouter(r)
	// 添加健康检查
	routes.AddHealth(r)
	service.Handle("/", r)

	err = service.Run()

	if err != nil {
		log.Fatal(err)
	}
}
