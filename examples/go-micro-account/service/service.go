package service

import "context"
import (
	proto "github.com/unliar/proto/account"
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

// PostUserBase 是用来创建用户基础信息的接口
func (a *Account) PostUserBase(ctx context.Context, req *proto.UserBase, rsp *proto.ResponseStatus) error {
	rsp.ErrMsg = "you bad bad PostUserBase"
	return nil
}

// PutUserBase 是用来更新用户基础信息的接口
func (a *Account) PutUserBase(ctx context.Context, req *proto.UserBase, rsp *proto.ResponseStatus) error {
	rsp.ErrMsg = "you bad bad PutUserBase"
	return nil
}

// PostUserContact 是用来创建用户联系信息的接口
func (a *Account) PostUserContact(ctx context.Context, req *proto.UserContact, rsp *proto.ResponseStatus) error {
	rsp.ErrMsg = "you bad bad PostUserContact"
	return nil
}

// PutUserContact 是用来更新用户联系方式的接口
func (a *Account) PutUserContact(ctx context.Context, req *proto.UserContact, rsp *proto.ResponseStatus) error {
	rsp.ErrMsg = "you bad bad PutUserContact"
	rsp.Status = 1
	return nil
}

// DelUserContact 是用来删除用户联系方式的接口
func (a *Account) DelUserContact(ctx context.Context, req *proto.UserId, rsp *proto.ResponseStatus) error {
	rsp.ErrMsg = "you bad bad DelUserContact"
	return nil
}
