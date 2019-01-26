package controller

import (
	"context"
	proto "github.com/unliar/proto/account"
	cpt "github.com/unliar/utils/go/crypto"
	"learnGo/examples/go-micro-account/service"
	"strings"
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
	ts := strings.TrimSpace(req.Password)
	hash, err := cpt.HashString(ts)
	if err != nil {
		rsp.Status = 2
		rsp.ErrMsg = err.Error()
		return nil
	}
	r, err := service.UpdateUserPasswordInfo(&service.UserPass{UID: req.UID, Password: hash})
	if err != nil {
		rsp.Status = r.Status
		rsp.ErrMsg = err.Error()
		return nil
	}
	rsp.Status = r.Status
	return nil
}

// CheckPassword 是检测用户密码的接口
func (a *AccountController) CheckPassword(ctx context.Context, req *proto.PasswordInput, rsp *proto.ResponseStatus) error {
	ts := strings.TrimSpace(req.Password)
	p, err := service.GetUserPasswordInfo(req.UID)
	if err != nil {
		rsp.Status = 2
		rsp.ErrMsg = err.Error()
		return err
	}
	r := cpt.MatchHashString(ts, p.Password)
	if r {
		rsp.Status = 1
		rsp.ErrMsg = "check passed"
		return nil
	}
	rsp.Status = 2
	rsp.ErrMsg = "check failed"
	return nil
}
