package controller

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
