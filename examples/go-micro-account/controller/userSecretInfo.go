package controller

import (
	"context"
	"fmt"
	"learnGo/examples/go-micro-account/service"
)
import proto "github.com/unliar/proto/account"

// CheckPhone 检测用户手机号是否被绑定
func (a *AccountController) CheckPhone(ctx context.Context, req *proto.UserSecretInfo, rsp *proto.ResponseStatus) error {
	_, err := service.GetSecretInfo(&service.UserSecretInfo{Phone: req.Phone})
	r, err := service.PostSecretInfo(&service.UserSecretInfo{Phone: req.Phone})
	fmt.Println(r.UID)
	if err != nil {
		rsp.Status = 2
		rsp.ErrMsg = err.Error()
		return nil
	}
	rsp.Status = 1
	rsp.ErrMsg = "have records"
	return nil
}
