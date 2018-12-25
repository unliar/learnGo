package service

import (
	"errors"
	proto "github.com/unliar/proto/account"
)

func GetUserInfoByUID(id int64) (*proto.UserInfo, error) {
	m := &proto.UserInfo{}
	if r := DB.First(m, id).RowsAffected; r > 0 {
		return m, nil
	}
	return m, errors.New("no user")
}
