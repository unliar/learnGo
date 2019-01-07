package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	ASV "github.com/unliar/proto/account"
	"strconv"
)

// GetUserInfo 根据用户id获取账户信息
func (a *AccountController) GetUserInfo(c *gin.Context) {
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

	resp, err := AccountService.GetUserInfo(context.TODO(), &ASV.UIDInput{
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
func (a *AccountController) PostUserInfo(c *gin.Context) {
	params := &RegisterRequest{}
	if err := c.ShouldBind(params); err != nil {
		c.JSON(422, &APIRSP{
			StatusCode: 422,
			Result:     nil,
			Detail:     err.Error(),
		})
		return
	}

	r, err := AccountService.RegisterUserByPassword(context.TODO(), &ASV.RegisterInfo{
		Nickname:  params.Nickname,
		LoginName: params.LoginName,
		Password:  params.Password,
	})
	if err != nil {
		c.JSON(500, &APIRSP{
			StatusCode: 422,
			Result:     nil,
			Detail:     err.Error(),
		})
		return
	}
	c.JSON(200, &APIRSP{
		StatusCode: 200,
		Result:     r,
	})
}

// UpdateUserInfo 更新用户信息
func (a *AccountController) UpdateUserInfo(c *gin.Context) {
	// 暂时不做吧 可能需要做单独的接口
	c.JSON(200, &APIRSP{
		StatusCode: 200,
	})
}

// GetHealthStatus 用于获取服务状态
func (a *AccountController) GetHealthStatus(c *gin.Context) {
	c.JSON(200, &APIRSP{
		StatusCode: 200,
		Detail:     "Server Status OK",
	})
}

// PostToken 是用来获取登录token凭证的
func (a *AccountController) PostToken(c *gin.Context) {

	c.SetCookie("USER_TOKEN", "qaq", 7200, "/", "", false, false)
	c.JSON(400, &APIRSP{
		StatusCode: 400,
		Detail:     "hi-PostToken",
		Result:     nil,
	})

}

// GetValueIsUnique 是检查用户登录名手机号昵称是否重复的接口
func (a *AccountController) GetValueIsUnique(c *gin.Context) {

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
		rsp, err := AccountService.CheckPhone(context.TODO(), &ASV.UserSecretInfo{
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
		rsp, err := AccountService.CheckNickname(context.TODO(), &ASV.UserInfo{Nickname: v})
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
		rsp, err := AccountService.CheckLoginName(context.TODO(), &ASV.UserInfo{LoginName: v})
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

func (a *AccountController) JWTAuth(t ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("hei~ you are in jwtauth")

		token, _ := c.Cookie("USER_TOKEN")

		// 需要token
		if len(t) == 0 {
			// token为空
			if token == "" {
				c.AbortWithStatusJSON(403, &APIRSP{
					StatusCode: 403,
					Detail:     "need token but empty",
					Result:     nil,
				})
				return
			}
			// token不为空
			r, err := AccountService.GetUserInfoByToken(context.TODO(), &ASV.UserInfoWithToken{
				Token: token,
			})
			if err != nil || r.Status != 1 {
				fmt.Println("JWTAuth failed")
				c.AbortWithStatusJSON(403, &APIRSP{
					StatusCode: 403,
					Detail:     err.Error(),
					Result:     nil,
				})
				return
			}
			fmt.Println("hi JWTAuth let you go~")
			c.Set("UID", r.Id)
			c.Next()
			return

		}
		// 可选的token
		if len(t) > 0 {
			fmt.Println("当前为可选token")
			// token不为空
			if token != "" {
				fmt.Println("==>有token")
				r, err := AccountService.GetUserInfoByToken(context.TODO(), &ASV.UserInfoWithToken{
					Token: token,
				})
				if err != nil || r.Status != 1 {
					fmt.Println("JWTAuth failed")
					c.AbortWithStatusJSON(403, &APIRSP{
						StatusCode: 403,
						Detail:     err.Error(),
						Result:     nil,
					})
					return
				}
				fmt.Println("hi JWTAuth let you go~")
				c.Set("UID", r.Id)
				c.Next()
				return
			}
			fmt.Println("没有token")
			c.Next()
		}
		// 没有token 而且 需要检查token

	}
}
