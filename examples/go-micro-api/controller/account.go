package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	ASV "github.com/unliar/proto/account"
	"learnGo/examples/go-micro-api/utils"
	"strconv"
)

var (
	// AccountSVService 账户服务
	AccountSVService ASV.AccountSVService
)

// APIRSP api错误返回值
type APIRSP struct {
	StatusCode int64       `json:"statusCode"`
	Detail     interface{} `json:"detail"`
	Result     interface{} `json:"result"`
}

type AccountContoller struct {
}

func init() {
	AccountSVService = ASV.NewAccountSVService("unliar-account", client.DefaultClient)

}

// GetUserInfo 根据用户id获取账户信息
func (a *AccountContoller) GetUserInfo(c *gin.Context) {
	var err error
	uid := c.Param("uid")
	UID, err := strconv.ParseInt(uid, 10, 64)

	if err != nil {
		c.JSON(500, &APIRSP{
			StatusCode: 400,
			Detail:     err,
		})
		return
	}

	resp, err := AccountSVService.GetUserInfo(context.TODO(), &ASV.UIDInput{
		UID: UID,
	})

	if err != nil {
		c.JSON(500, &APIRSP{
			StatusCode: 500,
			Detail:     err,
		})
		return
	}

	if resp.Id == UID {
		c.JSON(200, &APIRSP{
			StatusCode: 200,
			Result:     resp,
		})
		return
	}
	c.JSON(404, &APIRSP{
		StatusCode: 404,
		Detail:     "cant find user by this uid",
	})
}

// PostUserInfo 创建用户
func (a *AccountContoller) PostUserInfo(c *gin.Context) {
	var loginRequest LoinRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(422, &APIRSP{
			StatusCode: 422,
			Detail:     "params error",
			Result:     err,
		})
		return
	}
	// 登录类型
	types := []string{"phone", "email", "LoginName"}
	if !utils.ContainItem(types, c.PostForm("type")) {
		c.JSON(400, &APIRSP{
			StatusCode: 422,
			Detail:     "no match types",
			Result:     nil,
		})
		return
	}
	data, err := AccountSVService.RegisterUserByPassword(context.TODO(), &ASV.CheckPasswordInput{
		Type:     loginRequest.Type,
		Value:    loginRequest.Value,
		Password: loginRequest.Password,
	})
	if err != nil {
		c.JSON(400, &APIRSP{
			StatusCode: 400,
			Detail:     err,
			Result:     nil,
		})
		return
	}
	c.JSON(200, &APIRSP{
		StatusCode: 200,
		Result:     data,
	})
}

// UpdateUserInfo 更新用户信息
func (a *AccountContoller) UpdateUserInfo(c *gin.Context) {

	c.JSON(200, &APIRSP{
		StatusCode: 200,
	})
}

// GetHealthStatus 用于获取服务状态
func (a *AccountContoller) GetHealthStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  200,
		"message": "api server ok",
	})
}

// PostToken 是用来获取登录token凭证的
func (a *AccountContoller) PostToken(c *gin.Context) {

	//判断是刷新还是获取新的token
	t := c.Query("type")
	var loginRequest LoinRequest
	if t == "refresh" {
		// 获取cookies里的token
		Token, err := c.Cookie("USER_TOKEN")
		fmt.Println("isok", Token)
		if err != nil {
			c.JSON(403, &APIRSP{
				StatusCode: 403,
				Detail:     "NO TOKEN",
			})
			return
		}

		// 获取根据token获取用户信息并且生成新的token
		rsp, err := AccountSVService.GetUserInfoByToken(context.TODO(), &ASV.TokenInput{
			Token: Token,
		})

		if err != nil {
			c.JSON(400, &APIRSP{
				StatusCode: 403,
				Detail:     "call CheckToken error ",
				Result:     err,
			})
			return
		}
		if rsp.Status != 1 {
			c.JSON(403, &APIRSP{
				StatusCode: 403,
				Detail:     "user staus forbidden",
				Result:     nil,
			})
			return
		}
		// 生成新的token
		tokenMessage, err := AccountSVService.GetToken(context.TODO(), rsp.UserInfo)
		if err != nil {
			c.JSON(400, &APIRSP{
				StatusCode: 400,
				Detail:     "GetToken service error",
				Result:     err,
			})
			return
		}
		c.SetCookie("USER_TOKEN", tokenMessage.Token, 7200, "/", "", false, false)
		c.JSON(200, &APIRSP{
			StatusCode: 200,
			Detail:     "OK",
			Result:     tokenMessage,
		})

		return
	}
	// 此时是登录

	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(400, &APIRSP{
			StatusCode: 422,
			Detail:     "model error",
			Result:     err,
		})
		return
	}
	// 登录类型
	types := []string{"phone", "email", "LoginName"}
	if !utils.ContainItem(types, c.PostForm("type")) {
		c.JSON(400, &APIRSP{
			StatusCode: 422,
			Detail:     "no match types",
			Result:     nil,
		})
		return
	}
	resp, err := AccountSVService.CheckPassword(context.TODO(), &ASV.CheckPasswordInput{
		Type:     loginRequest.Type,
		Value:    loginRequest.Value,
		Password: loginRequest.Password,
	})
	if err != nil {
		c.JSON(500, &APIRSP{
			StatusCode: 500,
			Detail:     err,
			Result:     nil,
		})
		return
	}
	tokenMsg, err := AccountSVService.GetToken(context.TODO(), resp.UserInfo)
	if err != nil {
		c.JSON(500, &APIRSP{
			StatusCode: 500,
			Detail:     err,
			Result:     nil,
		})
		return
	}
	c.SetCookie("USER_TOKEN", tokenMsg.Token, 7200, "/", "", false, false)
	c.JSON(400, &APIRSP{
		StatusCode: 400,
		Detail:     "hi-PostToken",
		Result:     resp,
	})

}

// GetValueIsUnique 是检查用户登录名手机号昵称是否重复的接口
func (a *AccountContoller) GetValueIsUnique(c *gin.Context) {

	var uq UniqueQuery
	if err := c.ShouldBindQuery(&uq); err != nil {
		c.JSON(400, &APIRSP{
			StatusCode: 422,
			Detail:     "queryString model error",
			Result:     err,
		})
		return
	}

	t, v := uq.Type, uq.Value

	switch t {
	case "phone":
		rsp, err := AccountSVService.CheckPhone(context.TODO(), &ASV.UserInfo{
			Phone: v,
		})
		if err != nil {
			c.JSON(500, &APIRSP{
				StatusCode: 500,
				Detail:     "server err",
				Result:     err,
			})
			return
		}
		c.JSON(200, &APIRSP{
			StatusCode: 200,
			Detail:     t,
			Result:     rsp})
	case "nickname":
		rsp, err := AccountSVService.CheckNickname(context.TODO(), &ASV.UserInfo{Nickname: v})
		if err != nil {
			c.JSON(500, &APIRSP{
				StatusCode: 500,
				Detail:     "server err",
				Result:     err,
			})
			return
		}
		c.JSON(200, &APIRSP{
			StatusCode: 200,
			Detail:     t,
			Result:     rsp,
		})
	case "loginName":
		rsp, err := AccountSVService.CheckLoginName(context.TODO(), &ASV.UserInfo{LoginName: v})
		if err != nil {
			c.JSON(500, &APIRSP{
				StatusCode: 500,
				Detail:     "server err",
				Result:     err,
			})
			return
		}
		c.JSON(200, &APIRSP{
			StatusCode: 200,
			Detail:     t,
			Result:     rsp,
		})
	default:
		c.JSON(200, &APIRSP{
			StatusCode: 400,
			Detail:     "no matched type,value must be oneof phone nickname loginName",
			Result:     nil,
		})
	}

}
