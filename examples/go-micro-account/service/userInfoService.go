package service

import (
	"errors"
	"github.com/jinzhu/gorm"
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

func UpdateUserInfo(u *UserInfo) (*proto.ResponseStatus, error) {
	if r := DB.Model(u).Where(
		UserInfo{
			Model: gorm.Model{ID: u.ID},
		}).Updates(UserInfo{
		Nickname:   u.Nickname,
		Age:        u.Age,
		Gender:     u.Gender,
		Location:   u.Location,
		Profession: u.Profession,
		Brief:      u.Brief,
	}).RowsAffected; r <= 0 {
		return &proto.ResponseStatus{Status: 2}, errors.New("update fail")
	}
	return &proto.ResponseStatus{Status: 1}, nil
}
