package account

type UniqueQuery struct {
	Value string `form:"value" binding:"required"`
	Type  string `form:"type" binding:"required"`
}
