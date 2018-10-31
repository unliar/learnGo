package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unliar/proto/pay"
	"strconv"
)

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

	c.SecureJSON(200, rsp)
}

// PostPayInfo 用户创建支付信息
func (p *PayController) PostPayInfo(c *gin.Context) {
	info := PayInfoRequest{}
	info.UID = c.GetInt64("UID")
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(422, gin.H{
			"Status": 2,
			"ErrMsg": err.Error(),
		})
		return
	}

	rsp, err := PayService.PostPayInfo(context.TODO(), &pay.PayInfo{
		UID:    info.UID,
		Alipay: info.Alipay,
		TenPay: info.TenPay,
	})

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

	c.JSON(403, gin.H{
		"Status": 2,
		"ErrMsg": "have you long in?",
	})
}

// UpdatePayInfo 用户更新支付信息
func (p *PayController) UpdatePayInfo(c *gin.Context) {
	info := PayInfoRequest{}
	info.UID = c.GetInt64("UID")
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(422, gin.H{
			"Status": 2,
			"ErrMsg": err.Error(),
		})
		return
	}
	info.UID = c.GetInt64("UID")
	fmt.Println("===put==>", info)
	rsp, err := PayService.UpdatePayInfo(context.TODO(), &pay.PayInfo{
		UID:    info.UID,
		Alipay: info.Alipay,
		TenPay: info.TenPay,
	})
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, rsp)
}
