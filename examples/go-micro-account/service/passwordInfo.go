package service

import (
	"errors"
	proto "github.com/unliar/proto/account"
)

// GetUserPasswordInfo 是用来获取用户密码信息的接口
func GetUserPasswordInfo(id int64) (*proto.UserPasswordInfo, error) {
	// db 模型
	dup := &UserPass{UID: id}
	if r := DB.Where(dup).First(dup).RecordNotFound(); r {

		return dup.ToProto(), errors.New("no user passInfo")
	}
	return dup.ToProto(), nil
}

func UpdateUserPasswordInfo(u *UserPass) (*proto.ResponseStatus, error) {
	if r := DB.Model(&UserPass{}).
		Where(UserPass{UID: u.UID}).
		Updates(UserPass{Password: u.Password}).
		RowsAffected; r > 0 {
		return &proto.ResponseStatus{
			Status: 1,
		}, nil
	}
	// 更新失败
	return &proto.ResponseStatus{Status: 2}, errors.New("update failed")
}
