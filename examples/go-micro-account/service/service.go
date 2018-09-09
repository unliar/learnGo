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
	result := &UserInfo{}
	DB.First(result, req.GetUID())
	rsp.Id = int64(result.ID)
	rsp.LoginName = result.LoginName
	rsp.Age = result.Age
	rsp.Gender = proto.Gender(result.Gender)
	rsp.Avatar = result.Avatar
	rsp.Location = result.Location
	rsp.Profession = result.Profession
	rsp.Status = int32(result.Status)
	rsp.Phone = result.Phone
	rsp.Email = result.Email
	rsp.WeChatId = result.WeChatId
	rsp.QQId = result.QQId
	rsp.Brief = result.Brief
	rsp.NationCode = result.NationCode
	return nil
}

// PostUserInfo 是用来创建用户信息的方法
func (a *Account) PostUserInfo(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	result := &UserInfo{
		LoginName:  req.GetLoginName(),
		Age:        req.GetAge(),
		Gender:     int64(req.GetGender()),
		Avatar:     req.GetAvatar(),
		Location:   req.GetLocation(),
		Profession: req.GetProfession(),
		Status:     int64(req.GetStatus()),
		Phone:      req.GetPhone(),
		Email:      req.GetEmail(),
		WeChatId:   req.GetWeChatId(),
		QQId:       req.GetQQId(),
		Brief:      req.GetBrief(),
		NationCode: req.GetNationCode(),
	}
	IsOk := DB.NewRecord(result)
	if IsOk {
		DB.Create(result)
		rsp.Status = 1
		rsp.ErrMsg = strconv.FormatUint(uint64(result.ID), 10)
	}
	rsp.Status = 0
	rsp.ErrMsg = "not work"
	return nil
}

// UpdateUserInfo 是用来更新用户信息
func (a *Account) UpdateUserInfo(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	result := &UserInfo{}
	DB.First(result, req.GetId())
	// 更新年龄
	if req.GetAge() != 0 && result.Age != req.GetAge() {
		result.Age = req.GetAge()
	}
	// 性别
	if req.GetGender() != 0 && result.Gender != int64(req.GetGender()) {
		result.Gender = int64(req.GetGender())
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
