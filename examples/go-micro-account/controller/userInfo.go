package controller

import (
	"context"
	"errors"
	"fmt"
	proto "github.com/unliar/proto/account"
	"learnGo/examples/go-micro-account/config"
	"learnGo/examples/go-micro-account/service"
)

func (a *AccountController) GetUserInfo(ctx context.Context, req *proto.UIDInput, res *proto.UserInfo) error {
	r, err := service.GetUserInfoByUID(req.GetUID())
	if err != nil {
		return err
	}
	*res = *r
	return nil
}

func (a *AccountController) GetUserInfoByToken(ctx context.Context, req *proto.UserInfoWithToken, rsp *proto.UserInfo) error {
	token, tokenKey := req.Token, config.Config.JWTTokenKey

	uid, err := service.ParseToken(token, tokenKey)

	if err != nil {
		return err
	}

	r, err := service.GetUserInfoByUID(uid)

	if err != nil {
		return err
	}
	*rsp = *r
	return nil
}

// PostUserInfo 是用来创建用户信息的方法 看起来不会用了
func (a *AccountController) PostUserInfo(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	return nil
}

// UpdateUserInfo 更新
func (a *AccountController) UpdateUserInfo(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	return nil
}

// GetToken 是用来获取用户token的方法
func (a *AccountController) GetToken(ctx context.Context, req *proto.UserInfo, rsp *proto.UserInfoWithToken) error {
	// 获取用户信息
	u, err := service.GetUserInfoByUID(req.Id)
	if err != nil {
		rsp.Token = ""
		return err
	}
	// 检测用户状态
	if u.Status != 1 {
		rsp.Token = ""
		return errors.New("user status not ok")
	}
	// 获取密码信息
	p, err := service.GetUserPasswordInfo(u.Id)
	if err != nil {
		rsp.Token = ""
		return err
	}

	Token, _ := service.GeneratorToken(
		service.TokenPayload{
			UID:         req.Id,
			Status:      int32(u.Status),
			PassUpdated: p.UpdatedAt,
		},
		config.Config.JWTTokenKey,
	)
	fmt.Println("contorller token", Token)
	rsp.Token = Token
	rsp.UserInfo = u
	return nil
}

// CheckToken 是用来检测用户token的方法
func (a *AccountController) CheckToken(ctx context.Context, req *proto.UserInfoWithToken, rsp *proto.ResponseStatus) error {
	token, tokenKey := req.Token, config.Config.JWTTokenKey
	_, err := service.ParseToken(token, tokenKey)
	if err != nil {
		// token 无效
		rsp.Status = 2
		return err
	}
	rsp.Status = 1
	return nil
}

// CheckLoginName 是用来检测用户登录名称的方法
func (a *AccountController) CheckLoginName(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	u := &service.UserInfo{Nickname: req.LoginName}
	_, err := service.GetUserInfo(u)
	if err != nil {
		// LoginName可用 不存在记录
		rsp.Status = 1
		return nil
	}
	// LoginName不可用 已经存在记录
	rsp.Status = 2
	return nil

}

// CheckLoginName 是用来检测用户昵称是否重复的方法
func (a *AccountController) CheckNickname(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	u := &service.UserInfo{Nickname: req.Nickname}
	_, err := service.GetUserInfo(u)
	if err != nil {
		// 昵称可用 不存在记录
		rsp.Status = 1
		return nil
	}
	// 昵称不可用 已经存在了记录
	rsp.Status = 2
	return nil
}

// CheckLoginName 是用来密码注册的接口
func (a *AccountController) RegisterUserByPassword(ctx context.Context, req *proto.RegisterInfo, rsp *proto.UserInfo) error {

	tx := service.DB.Begin()

	u := &service.UserInfo{LoginName: req.LoginName, Nickname: req.Nickname}
	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&service.UserPass{Password: req.Password, UID: int64(u.ID)}).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	*rsp = *u.ToProto()
	return err
}
