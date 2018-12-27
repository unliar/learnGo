package service

import (
	"errors"
	proto "github.com/unliar/proto/account"
)

// QuerySecretInfo 用于查询SecretInfo表数据
func GetSecretInfo(u *proto.UserSecretInfo) (*proto.UserSecretInfo, error) {
	if r := DB.Where(u).First(u).RowsAffected; r > 0 {
		return u, nil
	}
	return u, errors.New("no that secretInfo")
}

// PutSecretInfo 用于更新SecretInfo表数据
func PutSecretInfo(u *proto.UserSecretInfo) (*proto.UserSecretInfo, error) {
	if r := DB.Model(&UserSecretInfo{UID: u.UID}).Updates(u).RowsAffected; r > 0 {
		return u, nil
	}
	return u, errors.New("no that secretInfo")
}

// PostSecretInfo 用于新增SecretInfo表数据
func PostSecretInfo(u *proto.UserSecretInfo) (*proto.UserSecretInfo, error) {
	if r := DB.Where(u).First(u).RowsAffected; r > 0 {
		return u, nil
	}
	return u, errors.New("no that secretInfo")
}

// DeleteSecretInfo 用于删除SecretInfo表数据
func DeleteSecretInfo(u *proto.UserSecretInfo) (*proto.UserSecretInfo, error) {
	if r := DB.Where(u).First(u).RowsAffected; r > 0 {
		return u, nil
	}
	return u, errors.New("no that secretInfo")
}
