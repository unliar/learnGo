package service

import "context"
import (
	proto "learnGo/examples/go-micro-account/proto"
)

type Account struct {
}

// GetUserInfo 是用来处理相关用户信息的接口
func (a *Account) GetUserInfo(ctx context.Context, req *proto.UserId, rsp *proto.UserInfo) error {
	uc := proto.UserContact{}
	ub := proto.UserBase{}
	var duc UserContact
	DB.Find(&duc, "uid = ? ", req.UID).First(&duc)
	uc.Phone, uc.Email, uc.WeiBoId, uc.WeChatId, uc.Id = duc.Phone, duc.Email, duc.WeiBoId, duc.WeChatId, duc.UID
	rsp.Base = &ub
	rsp.Contact = &uc

	return nil
}

// GetUserBase 是用来获取用户基础信息的接口
func (a *Account) GetUserBase(ctx context.Context, req *proto.UserId, rsp *proto.UserBase) error {
	rsp.Nickname = "qq"
	return nil
}

// GetUserContact 是用来获取用户联络信息的接口
func (a *Account) GetUserContact(ctx context.Context, req *proto.UserId, rsp *proto.UserContact) error {
	rsp.Phone = "99999"
	return nil
}
