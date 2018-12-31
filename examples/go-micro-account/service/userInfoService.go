package service

import (
	"errors"
	proto "github.com/unliar/proto/account"
)

func GetUserInfoByUID(id int64) (*proto.UserInfo, error) {
	m := &UserInfo{}

	if r := DB.First(m, id).RowsAffected; r > 0 {
		return m.ToProto(), nil
	}
	return m.ToProto(), errors.New("no user")
}
