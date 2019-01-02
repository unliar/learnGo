package service

import (
	"errors"
	proto "github.com/unliar/proto/account"
)

func GetUserInfoByUID(id int64) (*proto.UserInfo, error) {
	m := &UserInfo{}

	if r := DB.First(m, id).RecordNotFound(); r {
		return m.ToProto(), errors.New("no user")
	}
	return m.ToProto(), nil
}

func GetUserInfo(u *UserInfo) (*proto.UserInfo, error) {
	if r := DB.Where(u).First(u).RecordNotFound(); r {
		return u.ToProto(), errors.New("not found")
	}
	return u.ToProto(), nil
}
