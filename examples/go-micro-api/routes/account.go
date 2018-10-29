package routes

import "github.com/gin-gonic/gin"
import ctl "learnGo/examples/go-micro-api/controller"

// 账户模块
var AC ctl.AccountContoller

// AccountRouter 是用于添加账户相关路由
func AccountRouter(r *gin.Engine) *gin.Engine {

	// 检查用户登录名手机号昵称是否重复的接口
	r.GET("/api/account/unique", AC.GetValueIsUnique)
	// 获取用户信息
	r.GET("/api/account/users/:uid", AC.GetUserInfo)

	// 注册用户
	r.POST("/api/account/users", AC.PostUserInfo)

	// 修改用户基础信息
	r.PUT("/api/account/users", AC.JWTAuth(), AC.UpdateUserInfo)

	// 创建||刷新登录token 登录逻辑
	r.POST("/api/account/tokens", AC.JWTAuth("option"), AC.PostToken)
	return r
}
