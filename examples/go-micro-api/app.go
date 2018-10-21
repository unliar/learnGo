package main

import (
	ac "learnGo/examples/go-micro-api/account"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-web"
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
	// 健康检查api
	r.GET("/api/health", ac.GetHealthStatus)

	// 检查用户登录名手机号昵称是否重复的接口
	r.GET("/api/unique", ac.GetValueIsUnique)
	// 获取用户信息
	r.GET("/api/users/:uid", ac.GetUserInfo)

	// 注册用户
	r.POST("/api/users", ac.PostUserInfo)

	// 修改用户基础信息
	r.PUT("/api/users", ac.UpdateUserInfo)

	// 创建||刷新登录token
	r.POST("/api/tokens", ac.PostToken)

	service.Handle("/", r)

	err = service.Run()

	if err != nil {
		log.Fatal(err)
	}
}
