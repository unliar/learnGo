package account

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	ASV "github.com/unliar/proto/account"
)

var (
	// AccountSVService 账户服务
	AccountSVService ASV.AccountSVService
)

// Resp api错误返回值
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
	c.JSON(200, &APIRSP{
		StatusCode: 200,
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
