package controller

import (
	"context"
	"github.com/micro/go-micro/errors"
	proto "github.com/unliar/proto/account"
	"learnGo/examples/go-micro-account/service"
)

// AccountController 账户模块
type AccountController struct {
}

// GetUserPasswordUpdatedTime 获取账户密码更新时间
func (a *AccountController) GetUserPasswordUpdatedTime(ctx context.Context, req *proto.UIDInput, rsp *proto.UserPasswordInfo) error {
	r := &proto.UserPasswordInfo{UID: req.GetUID()}
	if e := service.DB.Model(r).First(r).Error; e != nil {
		return errors.BadRequest(string(1100001), "not found")
	}
	rsp.UID = req.UID
	rsp.CreatedAt = r.CreatedAt
	rsp.UpdatedAt = r.UpdatedAt
	return nil
}
