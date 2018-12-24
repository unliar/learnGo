package controller

import (
	"context"
)
import proto "github.com/unliar/proto/account"

func (a *AccountController) GetUserInfo(ctx context.Context, req *proto.UIDInput, res *proto.UserInfo) error {
	return nil
}

func (a *AccountController) GetUserInfoByToken(ctx context.Context, req *proto.UserInfoWithToken, rsp *proto.UserInfo) error {
	return nil
}

// PostUserInfo 是用来创建用户信息的方法
func (a *AccountController) PostUserInfo(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	return nil
}

// UpdateUserInfo 更新
func (a *AccountController) UpdateUserInfo(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	return nil
}

// PostUserInfo 是用来创建用户信息的方法
func (a *AccountController) GetToken(ctx context.Context, req *proto.UserInfo, rsp *proto.UserInfoWithToken) error {
	return nil
}

// PostUserInfo 是用来创建用户信息的方法
func (a *AccountController) CheckToken(ctx context.Context, req *proto.UserInfoWithToken, rsp *proto.ResponseStatus) error {
	return nil
}

// CheckLoginName 是用来检测用户登录名称的方法
func (a *AccountController) CheckLoginName(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	return nil
}

// CheckLoginName 是用来检测用户登录名称的方法
func (a *AccountController) CheckNickname(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	return nil
}

// CheckLoginName 是用来检测用户登录名称的方法
func (a *AccountController) RegisterUserByPassword(ctx context.Context, req *proto.RegisterInfo, rsp *proto.UserInfo) error {
	return nil
}
