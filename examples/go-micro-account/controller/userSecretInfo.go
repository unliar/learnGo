package controller

import (
	"context"
)
import proto "github.com/unliar/proto/account"

// CheckPhone 检测用户手机号是否被绑定
func (a *AccountController) CheckPhone(ctx context.Context, req *proto.UserSecretInfo, rsp *proto.ResponseStatus) error {
	return nil
}
