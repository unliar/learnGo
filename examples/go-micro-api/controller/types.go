package controller

import (
	"github.com/micro/go-micro/client"
	"github.com/unliar/proto/account"
	"github.com/unliar/proto/pay"
)

// ErrorMsg 错误信息
type ErrorMsg struct {
	Code    int64  `json:"code"`
	Detail  string `json:"detail"`
	Time    int64  `json:"time"`
	Message string `json:"message"`
}

// UniqueQuery 是地址栏参数
type UniqueQuery struct {
	Value string `form:"value" binding:"required"` // 对应类型的值
	Type  string `form:"type" binding:"required"`  // 类型
}

// LoinRequest 是用于登录的
type LoinRequest struct {
	Type     string `form:"type" json:"type" binding:"required"`         // 用于指明类型
	Value    string `form:"value" json:"value" binding:"required"`       // 手机号 邮箱 登录名
	Password string `form:"password" json:"password" binding:"required"` // 密码
}
type PayInfoRequest struct {
	UID    int64  `form:"uid" json:"uid"`
	Alipay string `form:"alipay" json:"alipay" binding:"required"`
	TenPay string `form:"tenpay" json:"tenpay" binding:"required"`
}

// APIRSP api错误返回值
type APIRSP struct {
	StatusCode int64       `json:"statusCode"`
	Detail     interface{} `json:"detail"`
	Result     interface{} `json:"result"`
}

type AccountController struct {
}

type PayController struct {
}

var AccountService = account.NewAccountSVService("unliar-account", client.DefaultClient)
var PayService = pay.NewPaySVService("unliar-pay", client.DefaultClient)
