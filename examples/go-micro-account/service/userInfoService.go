package service

import (
	"errors"
	proto "github.com/unliar/proto/account"
)

func GetUserInfoByUID(id int64) (*proto.UserInfo, error) {
	m := &UserInfo{}

	if r := DB.First(m, id).RecordNotFound(); !r {
		return m.ToProto(), nil
	}
	return m.ToProto(), errors.New("no user")
}
