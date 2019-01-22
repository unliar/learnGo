package routes

import "github.com/gin-gonic/gin"
import "learnGo/examples/go-micro-api/controller"

// 账户模块

// AccountRouter 是用于添加账户相关路由
func AccountRouter(r *gin.Engine) *gin.Engine {
	ac := controller.AccountController{}
	// 检查用户登录名手机号昵称是否重复的接口
	r.GET("/api/account/unique", ac.GetValueIsUnique)
	// 获取用户信息
	r.GET("/api/account/users/:uid", ac.GetUserInfo)

	// 注册用户
	r.POST("/api/account/users", ac.PostUserInfo)

	// 修改用户基础信息
	r.PUT("/api/account/users", ac.JWTAuth(), ac.UpdateUserInfo)

	// 创建||刷新登录token 登录逻辑
	r.POST("/api/account/tokens", ac.JWTAuth("option"), ac.PostToken)

	// 删除token逻辑
	r.DELETE("/api/account/tokens", ac.RemoveToken)
	return r
}
