package account

// UniqueQuery 是地址栏参数
type UniqueQuery struct {
	Value string `form:"value" binding:"required"` // 对应类型的值
	Type  string `form:"type" binding:"required"`  // 类型
}

// LoinRequest 是用于登录的
type LoinRequest struct {
	Type  string `form:"type" json:"type" binding:"required"`   // 用于指明类型
	Value string `form:"value" json:"value" binding:"required"` // 一般是密码或者验证码
	Key   string `form:"key" json:"key" binding:"required"`     // 指定类型对应的值
}
