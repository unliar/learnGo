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
	// 健康检查api
	r.GET("/api/health", ac.GetHealthStatus)
	// 用户信息
	r.GET("/api/users/:uid", ac.GetUserBase)
	// 注册用户
	r.POST("/api/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ok": 1,
		})
	})
	// 修改用户基础信息
	r.PUT("/api/users/:uid/baseInfo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ok": 1,
		})
	})
	// 修改用户联系信息
	r.PUT("/api/users/:uid/userContact", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ok": 1,
		})
	})
	// 创建||刷新登录token
	r.POST("/api/tokens", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ok": 1,
		})
	})

	service.Handle("/", r)

	err = service.Run()

	if err != nil {
		log.Fatal(err)
	}
}
