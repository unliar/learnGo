package service

import "context"
import proto "learnGo/examples/go-micro-account/proto"

type Account struct {
}

func (a *Account) GetUserInfo(ctx context.Context, req *proto.UserId, rsp *proto.UserInfo) error {
	return nil
}

func (a *Account) GetUserBase(ctx context.Context, req *proto.UserId, rsp *proto.UserBase) error {
	return nil
}
func (a *Account) GetUserContact(ctx context.Context, req *proto.UserId, rsp *proto.UserContact) error {
	return nil
}
