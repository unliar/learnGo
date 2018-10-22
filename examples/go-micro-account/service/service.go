package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	proto "github.com/unliar/proto/account"
)

// SignKey 用于token签名
var SignKey = []byte("hikey")

// Account 账户模块
type Account struct {
}

// Payload 是用来发token的
type Payload struct {
	UID    int64 `json:"uid"`
	Status int32 `json:"userStatus"`
	jwt.StandardClaims
}

// GetUserInfo 是用来获取用户信息的接口
func (a *Account) GetUserInfo(ctx context.Context, req *proto.UIDInput, rsp *proto.UserInfo) error {
	result := &proto.UserInfo{}
	DB.First(result, req.GetUID())
	*rsp = *result
	return nil
}

// PostUserInfo 是用来创建用户信息的方法
func (a *Account) PostUserInfo(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	result := &proto.UserInfo{}
	IsOk := DB.NewRecord(result)
	if IsOk {
		DB.Create(result)
		rsp.Status = 1
		rsp.ErrMsg = strconv.FormatInt(result.Id,10)
		return nil
	}
	rsp.Status = 0
	rsp.ErrMsg = "not work"
	return nil
}

// UpdateUserInfo 是用来更新用户信息
func (a *Account) UpdateUserInfo(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	result := &proto.UserInfo{}
	DB.First(result, req.GetId())
	// 更新年龄
	if req.GetAge() != 0 && result.Age != req.GetAge() {
		result.Age = req.GetAge()
	}
	// 性别
	if req.GetGender() != 0 && result.Gender != req.GetGender() {
		result.Gender = req.GetGender()
	}
	// 头像
	if req.GetAvatar() != "" && result.Avatar != req.GetAvatar() {
		result.Avatar = req.GetAvatar()
	}
	// 位置
	if req.GetLocation() != "" && result.Location != req.GetLocation() {
		result.Location = req.GetLocation()
	}
	// 手机
	if req.GetPhone() != "" && result.Phone != req.GetPhone() {
		result.Phone = req.GetPhone()
	}
	// 邮箱
	if req.GetEmail() != "" && result.Email != req.GetEmail() {
		result.Email = req.GetEmail()
	}
	// 微信
	if req.GetWeChatId() != "" && result.WeChatId != req.GetWeChatId() {
		result.WeChatId = req.GetWeChatId()
	}
	// Brief简介
	if req.GetBrief() != "" && result.Brief != req.GetBrief() {
		result.Brief = req.GetBrief()
	}
	// 地区码
	if req.GetNationCode() != "" && result.NationCode != req.GetNationCode() {
		result.NationCode = req.GetNationCode()
	}
	DB.Save(result)
	return nil
}

// GetToken 是用来获取合法token的接口
func (a *Account) GetToken(ctx context.Context, req *proto.UserInfo, rsp *proto.TokenMessage) error {
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
func (a *Account) CheckToken(ctx context.Context, req *proto.TokenInput, rsp *proto.ResponseStatus) error {

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
	result := &UserInfo{}
	DB.First(result, "nickname=?", req.GetNickname())
	fmt.Println("CheckNickname", result)
	if result.ID > 0 {
		rsp.Status = 2
		rsp.ErrMsg = "the Nickname is used"
		return nil
	}
	rsp.Status = 1
	rsp.ErrMsg = "not used"
	return nil
}

// UpdatePassword 是更新用户密码的接口
func (a *Account) UpdatePassword(ctx context.Context, req *proto.UpdatePassInput, rsp *proto.ResponseStatus) error {
	return nil
}

// CheckLoginName 是检查登录名称的接口
func (a *Account) CheckLoginName(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	result := &UserInfo{}
	DB.First(result, "login_name=?", req.GetLoginName())
	fmt.Println("CheckLoginName", result)
	if result.ID > 0 {
		rsp.Status = 2
		rsp.ErrMsg = "the login_name is used"
		return nil
	}
	rsp.Status = 1
	rsp.ErrMsg = "login_name not used"
	return nil

}

// CheckPhone 是检查手机号的接口
func (a *Account) CheckPhone(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	result := &UserInfo{}
	DB.First(result, "phone=?", req.GetPhone())
	fmt.Println("CheckPhone", result)
	if result.ID > 0 {
		rsp.Status = 2
		rsp.ErrMsg = "the phone is used"
		return nil
	}
	rsp.Status = 1
	rsp.ErrMsg = "phone not used"
	return nil
}

// GetUserInfoByToken 是用token获取用户信息的接口
func (a *Account) GetUserInfoByToken(ctx context.Context, req *proto.TokenInput, rsp *proto.UserInfoByTokenResponse) error {
	// 检测token
	token, _ := jwt.ParseWithClaims(req.GetToken(), &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return SignKey, nil
	})
	// token 合法--->获取用户信息
	if claims, ok := token.Claims.(*Payload); ok && token.Valid {
		fmt.Println(claims.UID, claims.Status, claims)
		rsp.Status = 1
		result := &proto.UserInfo{}
		DB.First(result, claims.UID)
		rsp.UserInfo = result
		return nil
	}
	// 用id获取信息
	rsp.Status = 2
	rsp.UserInfo = nil
	return nil
}
