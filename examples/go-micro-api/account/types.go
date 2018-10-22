package account

// UniqueQuery 是地址栏参数
type UniqueQuery struct {
	Value string `form:"value" binding:"required"`
	Type  string `form:"type" binding:"required"`
}
