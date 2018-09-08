package service

import "context"
import (
	proto "github.com/unliar/proto/account"
)

// Account 账户模块
type Account struct {
}

// GetUserInfo 是用来获取用户信息的接口
func (a *Account) GetUserInfo(ctx context.Context, req *proto.UIDInput, rsp *proto.UserInfo) error {
	return nil
}

// PostUserInfo 是用来修改用户信息的方法
func (a *Account) PostUserInfo(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {

	return nil
}

// UpdateUserInfo 是用来更新用户信息
func (a *Account) UpdateUserInfo(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {

	return nil
}

// GetToken 是用来获取合法token的接口
func (a *Account) GetToken(ctx context.Context, req *proto.UserInfo, rsp *proto.TokenMessage) error {

	return nil
}

// CheckToken 是用来检测用户token
func (a *Account) CheckToken(ctx context.Context, req *proto.TokenInput, rsp *proto.ResponseStatus) error {

	return nil
}
