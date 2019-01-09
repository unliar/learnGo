package controller

import (
	"context"
	"learnGo/examples/go-micro-account/service"
)
import proto "github.com/unliar/proto/account"

// CheckPhone 检测用户手机号是否被绑定
func (a *AccountController) CheckPhone(ctx context.Context, req *proto.UserSecretInfo, rsp *proto.ResponseStatus) error {
	_, err := service.GetSecretInfo(&service.UserSecretInfo{Phone: req.Phone})
	if err != nil {
		rsp.Status = 2
		rsp.ErrMsg = err.Error()
		return nil
	}
	rsp.Status = 1
	rsp.ErrMsg = "have records"
	return nil
}

// GetUserUIDByUserSecretInfo 获取用户的UID
func (a *AccountController) GetUserUIDByUserSecretInfo(ctx context.Context, req *proto.UserSecretInfo, rsp *proto.UIDInput) error {
	r, err := service.GetSecretInfo(&service.UserSecretInfo{
		Phone:    req.Phone,
		Email:    req.Email,
		RealName: req.RealName,
		WeChatId: req.WeChatId,
		QQId:     req.QQId,
	})
	if err != nil {
		return err
	}
	rsp.UID = r.UID
	return nil
}

// GetUserInfoByUserSecretInfo 获取用户的信息
func (a *AccountController) GetUserInfoByUserSecretInfo(ctx context.Context, req *proto.UserSecretInfo, rsp *proto.UserInfo) error {
	r, err := service.GetSecretInfo(&service.UserSecretInfo{
		Phone:    req.Phone,
		Email:    req.Email,
		RealName: req.RealName,
		WeChatId: req.WeChatId,
		QQId:     req.QQId,
	})
	if err != nil {
		return err
	}
	u, err := service.GetUserInfoByUID(r.UID)
	*rsp = *u
	return nil
}
