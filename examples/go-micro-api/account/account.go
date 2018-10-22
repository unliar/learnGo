package account

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	ASV "github.com/unliar/proto/account"
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

func init() {
	AccountSVService = ASV.NewAccountSVService("unliar-account", client.DefaultClient)
}

// GetUserInfo 根据用户id获取账户信息
func GetUserInfo(c *gin.Context) {
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
func PostUserInfo(c *gin.Context) {
	data, _ := AccountSVService.PostUserInfo(context.TODO(), &ASV.UserInfo{
		LoginName: "admin",
		Nickname:  "admin",
	})
	c.JSON(200, &APIRSP{
		StatusCode: 200,
		Result:     data,
	})
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(c *gin.Context) {

	c.JSON(200, &APIRSP{
		StatusCode: 200,
	})
}

// GetHealthStatus 用于获取服务状态
func GetHealthStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  200,
		"message": "api server ok",
	})
}

// PostToken 是用来获取登录token凭证的
func PostToken(c *gin.Context) {
	var loginRequest LoinRequest

	//判断是刷新还是获取新的token
	t := c.Query("type")
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

		if err != nil || rsp.Status != 1 {
			c.JSON(400, &APIRSP{
				StatusCode: 403,
				Detail:     "call CheckToken error ",
				Result:     err,
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
	switch loginRequest.Type {
	case "email":
		c.JSON(200, &APIRSP{
			StatusCode: 200,
			Detail:     "type==>" + loginRequest.Type,
			Result:     nil,
		})

	case "phone":

		c.JSON(200, &APIRSP{
			StatusCode: 200,
			Detail:     "type==>" + loginRequest.Type,
			Result:     nil,
		})

	case "loginName":
		c.JSON(200, &APIRSP{
			StatusCode: 200,
			Detail:     "type==>" + loginRequest.Type,
			Result:     nil,
		})

	default:
		c.JSON(400, &APIRSP{
			StatusCode: 400,
			Detail:     "no match type",
			Result:     nil,
		})
	}

}

// GetValueIsUnique 是检查用户登录名手机号昵称是否重复的接口
func GetValueIsUnique(c *gin.Context) {

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
