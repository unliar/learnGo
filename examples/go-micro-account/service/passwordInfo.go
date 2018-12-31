package service

import (
	"errors"
	proto "github.com/unliar/proto/account"
)

// GetUserPasswordInfo 是用来获取用户密码信息的接口
func GetUserPasswordInfo(id int64) (*proto.UserPasswordInfo, error) {
	// db 模型
	dup := &UserPass{UID: id}
	if r := DB.Where(dup).First(dup).RowsAffected; r > 0 {

		return dup.ToProto(), nil
	}
	return dup.ToProto(), errors.New("no user passInfo")
}
