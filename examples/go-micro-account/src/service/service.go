package service

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/errors"
	"github.com/satori/go.uuid"
	"learnGo/examples/go-micro-account/src/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	proto "github.com/unliar/proto/account"
)

// SignKey 用于token签名
var SignKey = []byte("hikey")
var MD5Key = "hi-md5"

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
func (a *Account) UpdatePassword(ctx context.Context, req *proto.UpdatePassInput, rsp *proto.ResponseStatus) error {
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
func (a *Account) CheckPhone(ctx context.Context, req *proto.UserInfo, rsp *proto.ResponseStatus) error {
	result := &proto.UserInfo{Phone: req.GetPhone()}
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

// CheckPassword 是用于检测账户登录的接口
func (a *Account) CheckPassword(ctx context.Context, req *proto.CheckPasswordInput, rsp *proto.UserInfoByTokenResponse) error {
	var t string
	UserInfo := &proto.UserInfo{}
	PasswordMD5 := utils.CreateMD5(req.Password, MD5Key)
	switch req.GetType() {
	case "phone":
		t = "user_infos.phone = ? AND user_passes.password = ?"
	case "email":
		t = "user_infos.email = ? AND user_passes.password = ?"
	case "loginName":
		t = "user_infos.login_name = ? AND user_passes.password = ?"
	default:
		return errors.BadRequest("400", "not matched type===>%s", req.GetType())
	}
	q := "left join user_passes on user_passes.uid = user_infos.id"
	if r := DB.Model(UserInfo).Joins(q).Where(t, req.GetValue(), PasswordMD5).First(UserInfo).RowsAffected; r > 0 {
		fmt.Println("CheckPassword db result=============>", UserInfo)
		rsp.Status = 1
		rsp.UserInfo = UserInfo
		return nil
	}
	rsp.Status = 2
	rsp.UserInfo = nil
	return nil
}

// RegisterUserByPassword 是用于密码注册的方法
func (a *Account) RegisterUserByPassword(ctx context.Context, req *proto.CheckPasswordInput, rsp *proto.UserInfo) error {
	// 设置用户信息表
	User := &UserInfo{}
	// 生成随机昵称和登录名
	UUID := uuid.Must(uuid.NewV4())
	User.Nickname = fmt.Sprintf("%s", UUID)
	User.LoginName = fmt.Sprintf("%s", UUID)
	// 设置密码表
	Pass := &UserPass{}
	PasswordMD5 := utils.CreateMD5(req.GetPassword(), MD5Key)
	Pass.Password = PasswordMD5
	// 判断手机号 邮箱 登录名是否使用过
	switch req.GetType() {
	case "phone":
		User.Phone = req.GetValue()
		if r := DB.First(User, "phone = ?", req.GetValue()).RowsAffected; r > 0 {
			return errors.BadRequest("4001", "%s", "手机已存在")
		}

	case "email":
		User.Email = req.GetValue()
		if r := DB.First(User, "email = ?", req.GetValue()).RowsAffected; r > 0 {
			return errors.BadRequest("4002", "%s", "邮箱已存在")
		}

	case "LoginName":
		User.LoginName = req.GetValue()
		if r := DB.First(User, "login_name = ?", req.GetValue()).RowsAffected; r > 0 {
			return errors.BadRequest("4003", "%s", "登录名已存在")
		}
	default:
		return errors.BadRequest("400", "type errors ===> %s", req.GetType())

	}
	// 启动事务操作
	tx := DB.Begin()
	if err := tx.Create(User).Error; err != nil {
		tx.Rollback()
		return nil
	}
	Pass.UID = int64(User.ID)
	if err := tx.Create(Pass).Error; err != nil {
		tx.Rollback()
		return nil
	}
	fmt.Print("user====>", User)
	tx.Commit()
	rsp.Id = int64(User.ID)
	rsp.LoginName = User.LoginName
	rsp.Phone = User.Phone
	rsp.Email = User.Email
	return nil
}
