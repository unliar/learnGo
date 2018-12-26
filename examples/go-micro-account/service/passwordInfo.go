package service

import (
	"errors"
	proto "github.com/unliar/proto/account"
)

// GetUserPasswordInfo 是用来获取用户密码信息的接口
func GetUserPasswordInfo(id int64) (*proto.UserPasswordInfo, error) {
	// 协议
	up := &proto.UserPasswordInfo{}
	// db 模型
	dup := &UserPass{UID: id}
	if r := DB.Where(dup).First(dup).RowsAffected; r > 0 {
		up.UID = id
		up.UpdatedAt = dup.UpdatedAt.Unix()
		up.CreatedAt = dup.CreatedAt.Unix()
		up.Password = dup.Password
		return up, nil
	}
	return up, errors.New("no user passInfo")
}
