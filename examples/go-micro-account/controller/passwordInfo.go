package controller

import (
	"context"
	proto "github.com/unliar/proto/account"
	cpt "github.com/unliar/utils/go/crypto"
	"learnGo/examples/go-micro-account/config"
	"learnGo/examples/go-micro-account/service"
)

// AccountController 账户模块
type AccountController struct {
}

// GetUserPasswordUpdatedTime 获取账户密码更新时间
func (a *AccountController) GetUserPasswordUpdatedTime(ctx context.Context, req *proto.UIDInput, rsp *proto.UserPasswordInfo) error {

	r, err := service.GetUserPasswordInfo(req.GetUID())
	if err != nil {
		return err
	}
	*rsp = *r
	return nil
}

// UpdatePassword 是更新用户密码的接口
func (a *AccountController) UpdatePassword(ctx context.Context, req *proto.UserPasswordInfo, rsp *proto.ResponseStatus) error {
	hash := cpt.GetMD5(req.Password, config.Config.MD5Key)
	r, err := service.UpdateUserPasswordInfo(&service.UserPass{UID: req.UID, Password: hash})
	if err != nil {
		rsp.Status = r.Status
		return nil
	}
	rsp.Status = r.Status
	return nil
}

// CheckPassword 是检测用户密码的接口
func (a *AccountController) CheckPassword(ctx context.Context, req *proto.PasswordInput, rsp *proto.UserInfoWithToken) error {

	return nil
}
