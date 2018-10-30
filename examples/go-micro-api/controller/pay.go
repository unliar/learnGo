package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unliar/proto/pay"
	"strconv"
)

type PayController struct {
}

// GetPayInfo 用于获取指定id的支付信息
func (p *PayController) GetPayInfo(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 64)

	if err != nil {
		c.JSON(422, gin.H{
			"Status": 2,
			"ErrMsg": err.Error(),
		})
		return
	}
	rsp, err := PayService.GetPayInfo(context.TODO(), &pay.PayInfo{
		UID: uid,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"result": rsp,
	})
}

// PostPayInfo 用户创建支付信息
func (p *PayController) PostPayInfo(c *gin.Context) {
	P := &pay.PayInfo{}
	if _, ok := c.Get("UID"); ok == true {
		if err := c.ShouldBind(P); err != nil {
			fmt.Println("ppppp", P)
			c.JSON(422, gin.H{
				"Status": 2,
				"ErrMsg": err.Error(),
			})
			return
		}

		rsp, err := PayService.PostPayInfo(context.TODO(), P)

		if err != nil {
			c.JSON(500, gin.H{
				"Status": 2,
				"ErrMsg": err.Error(),
			})
			return
		}
		c.JSON(500, gin.H{
			"Status": 1,
			"ErrMsg": "ok",
			"Result": rsp,
		})
		return
	}
	c.JSON(403, gin.H{
		"Status": 2,
		"ErrMsg": "have you long in?",
	})
}

// UpdatePayInfo 用户更新支付信息
func (p *PayController) UpdatePayInfo(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"statusCode": 200,
	})
}
