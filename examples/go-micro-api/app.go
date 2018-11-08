package main

import (
	"learnGo/examples/go-micro-api/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-web"
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
	// 添加跨域中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://192.168.0.143:3000", "http://localhost:3000", "http://192.168.31.236:3000", "http://127.0.0.1:3000"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// 404 error
	r.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{
			"statusCode": 404,
			"data":       nil,
			"error":      "not matched router path",
		})
	})

	// 添加账户模块
	routes.AccountRouter(r)
	// 添加pay模块
	routes.PayRouter(r)
	// 添加健康检查
	routes.AddHealth(r)

	service.Handle("/", r)

	err = service.Run()

	if err != nil {
		log.Fatal(err)
	}
}
