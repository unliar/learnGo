package service

import proto "github.com/unliar/proto/account"

// 用于将数据模型转换成协议模型

// ToProto 用户密码转proto
func (u *UserPass) ToProto() *proto.UserPasswordInfo {
	return &proto.UserPasswordInfo{
		UID:       u.UID,
		Password:  u.Password,
		CreatedAt: u.CreatedAt.Unix(),
		UpdatedAt: u.UpdatedAt.Unix(),
	}
}

func (u *UserInfo) ToProto() *proto.UserInfo {
	return &proto.UserInfo{
		Id:         int64(u.ID),
		LoginName:  u.LoginName,
		Nickname:   u.Nickname,
		Age:        u.Age,
		Gender:     proto.Gender(u.Gender),
		Avatar:     u.Avatar,
		Location:   u.Location,
		Profession: u.Profession,
		Status:     u.Status,
		Brief:      u.Brief,
	}
}

func (u *UserSecretInfo) ToProto() *proto.UserSecretInfo {
	return &proto.UserSecretInfo{
		UID:      u.UID,
		Email:    u.Email,
		Phone:    u.Phone,
		WeChatId: u.WeChatId,
		QQId:     u.QQId,
		RealName: u.RealName,
	}
}
