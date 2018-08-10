package service

import "context"
import (
	proto "learnGo/examples/go-micro-account/proto"
	"fmt"
)

type Account struct {
}

func (a *Account) GetUserInfo(ctx context.Context, req *proto.UserId, rsp *proto.UserInfo) error {
	fmt.Println(rsp.Contact)
	uc:=proto.UserContact{}
	ub:=proto.UserBase{}
	uc.Phone="2222222"
	rsp.Base=&ub
	rsp.Contact=&uc
	return nil
}

func (a *Account) GetUserBase(ctx context.Context, req *proto.UserId, rsp *proto.UserBase) error {
	rsp.Nickname = "qq"
	return nil
}
func (a *Account) GetUserContact(ctx context.Context, req *proto.UserId, rsp *proto.UserContact) error {
	rsp.Phone = "99999"
	return nil
}
