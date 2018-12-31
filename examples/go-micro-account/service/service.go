package service

import (
	"context"
	"fmt"
	"learnGo/examples/go-micro-account/config"
	"strconv"
	"time"

	"github.com/micro/go-micro/errors"

	jwt "github.com/dgrijalva/jwt-go"
	proto "github.com/unliar/proto/account"
)

// SignKey 用于token签名
var SignKey = []byte(config.Config.JWTTokenKey + config.Config.Env)

// MD5Key 用于md5加密
var MD5Key = config.Config.MD5Key + config.Config.Env

// Account 账户模块
type Account struct {
}

// Payload 是用来发token的
type Payload struct {
	UID    int64 `json:"uid"`
	Status int64 `json:"userStatus"`
	jwt.StandardClaims
}

// GetUserInfo 是用来获取用户信息的接口
func (a *Account) GetUserInfo(ctx context.Context, req *proto.UIDInput, rsp *proto.UserInfo) error {
	result := &proto.UserInfo{}
	if r := DB.First(result, req.GetUID()).RowsAffected; r > 0 {
		*rsp = *result
		return nil
	}
	return errors.NotFound("400", "user not found %v", req.UID)
}

// PostUserInfo 是用来创建用户信息的方法
func (a *Account) PostUserInfo(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	result := &proto.UserInfo{}
	IsOk := DB.NewRecord(result)
	if IsOk {
		DB.Create(result)
		rsp.Status = 1
		rsp.ErrMsg = strconv.FormatInt(result.Id, 10)
		return nil
	}
	rsp.Status = 0
	rsp.ErrMsg = "not work"
	return nil
}

// UpdateUserInfo 是用来更新用户信息
func (a *Account) UpdateUserInfo(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	result := &proto.UserInfo{}

	count := DB.Model(result).Where("id = ?", req.GetId()).Updates(req).RowsAffected
	fmt.Println("UpdateUserInfo====>count", count)
	if count == 1 {
		rsp.Status = proto.Status_Ok
		rsp.ErrMsg = "ok"
		return nil
	}
	rsp.Status = proto.Status_Failed
	rsp.ErrMsg = "check if you have the correct id or has updated"
	return nil
}

// GetToken 是用来获取合法token的接口
func (a *Account) GetToken(ctx context.Context, req *proto.UserInfo, rsp *proto.UserInfoWithToken) error {
	claims := Payload{
		req.GetId(),
		req.GetStatus(),
		jwt.StandardClaims{
			Issuer:    "accountSrv",
			Subject:   "AccountToken",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 70000,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SignKey)
	rsp.Token = signedToken
	return err
}

// CheckToken 是用来检测用户token
func (a *Account) CheckToken(ctx context.Context, req *proto.UserInfoWithToken, rsp *proto.ResponseStatus) error {

	token, err := jwt.ParseWithClaims(req.Token, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return SignKey, nil
	})
	if err != nil {
		rsp.ErrMsg = "Parse Token error"
		rsp.Status = 2
		return nil
	}
	if claims, ok := token.Claims.(*Payload); ok && token.Valid {
		fmt.Println(claims.UID, claims.Status, claims)
		rsp.Status = 1
		rsp.ErrMsg = ""
		return nil
	}
	rsp.Status = 2
	rsp.ErrMsg = "Token not ok"
	return nil
}

// CheckNickname 是用来检测用户昵称
func (a *Account) CheckNickname(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	result := &proto.UserInfo{Nickname: req.GetNickname()}
	if r := DB.Where(result).First(result).RowsAffected; r > 0 {
		rsp.Status = proto.Status_Failed
		rsp.ErrMsg = "nickname is used"
		return nil
	}

	rsp.Status = proto.Status_Ok
	rsp.ErrMsg = "nickname is not used"
	return nil
}

// UpdatePassword 是更新用户密码的接口
func (a *Account) UpdatePassword(ctx context.Context, req *proto.UserPasswordInfo, rsp *proto.ResponseStatus) error {
	result := &UserPass{Password: req.GetPassword()}
	if r := DB.Model(result).Where("uid = ?", req.GetUID()).Updates(result).RowsAffected; r > 0 {
		rsp.Status = proto.Status_Ok
		rsp.ErrMsg = "ok"
		return nil
	}
	rsp.Status = proto.Status_Failed
	rsp.ErrMsg = "fail update"
	return nil
}

// CheckLoginName 是检查登录名称的接口
func (a *Account) CheckLoginName(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	result := &proto.UserInfo{LoginName: req.GetLoginName()}
	if r := DB.Where(result).First(result).RowsAffected; r > 0 {
		rsp.Status = proto.Status_Failed
		rsp.ErrMsg = "the login_name is used"
		return nil
	}
	rsp.Status = proto.Status_Ok
	rsp.ErrMsg = "login_name not used"
	return nil

}

// CheckPhone 是检查手机号的接口
func (a *Account) CheckPhone(ctx context.Context, req *proto.UserSecretInfo, rsp *proto.ResponseStatus) error {
	result := &proto.UserSecretInfo{Phone: req.GetPhone()}
	if r := DB.First(result).First(result).RowsAffected; r > 0 {
		rsp.Status = 2
		rsp.ErrMsg = "the phone is used"
		return nil
	}
	rsp.Status = 1
	rsp.ErrMsg = "phone not used"
	return nil
}

// GetUserInfoByToken 是用token获取用户信息的接口
func (a *Account) GetUserInfoByToken(ctx context.Context, req *proto.UserInfoWithToken, rsp *proto.UserInfo) error {
	// 检测token
	token, _ := jwt.ParseWithClaims(req.GetToken(), &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return SignKey, nil
	})
	// token 合法--->获取用户信息
	if claims, ok := token.Claims.(*Payload); ok && token.Valid {
		fmt.Println(claims.UID, claims.Status, claims)
		result := &proto.UserInfo{}
		DB.First(result, claims.UID)
		rsp = result
		return nil
	}
	return nil
}

// CheckPassword 是用于检测账户登录的接口
func (a *Account) CheckPassword(ctx context.Context, req *proto.PasswordInput, rsp *proto.UserInfoWithToken) error {

	return nil
}

// RegisterUserByPassword 是用于密码注册的方法
func (a *Account) RegisterUserByPassword(ctx context.Context, req *proto.RegisterInfo, rsp *proto.UserInfo) error {

	return nil
}

// GetUserPasswordUpdatedTime 获取用户更新密码的时间
func (a *Account) GetUserPasswordUpdatedTime(ctx context.Context, req *proto.UIDInput, rsp *proto.UserPasswordInfo) error {
	return nil
}
