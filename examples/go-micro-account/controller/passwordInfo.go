package controller

import (
	"context"

	"google.golang.org/grpc"

	proto "github.com/unliar/proto/account"
)

// AccountController 账户模块
type AccountController struct {
}

// GetUserPasswordUpdatedTime 获取账户密码更新时间
func (a *AccountController) GetUserPasswordUpdatedTime(ctx context.Context, req *proto.UIDInput, rsp *proto.UserPasswordInfo) error {
	return grpc.Errorf(1100001, "%s", "qaq")
}
