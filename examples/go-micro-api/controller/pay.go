package controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/unliar/proto/pay"
)

// GetPayInfo 用于获取指定id的支付信息
func (p *PayController) GetPayInfo(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 64)

	if err != nil {
		c.JSON(422, &APIRSP{
			StatusCode: 422,
			Result:     nil,
			Detail:     err.Error(),
		})
		return
	}
	rsp, err := PayService.GetPayInfo(context.TODO(), &pay.PayInfo{
		UID: uid,
	})
	if err != nil {
		c.JSON(400, &APIRSP{
			StatusCode: 400,
			Result:     nil,
			Detail:     err.Error(),
		})

		return
	}
	if rsp.Status == 2 {
		c.JSON(500, &APIRSP{
			StatusCode: 200,
			Result:     rsp,
			Detail:     "not found",
		})
		return
	}
	c.SecureJSON(200, &APIRSP{
		StatusCode: 200,
		Result:     rsp,
		Detail:     "ok",
	})
}

// PostPayInfo 用户创建支付信息
func (p *PayController) PostPayInfo(c *gin.Context) {
	info := &PayInfoRequest{}
	info.UID = c.GetInt64("UID")
	if err := c.ShouldBind(info); err != nil {
		c.JSON(422, &APIRSP{
			StatusCode: 422,
			Result:     nil,
			Detail:     err.Error(),
		})
		return
	}

	rsp, err := PayService.PostPayInfo(context.TODO(), &pay.PayInfo{
		UID:    info.UID,
		Alipay: info.Alipay,
		TenPay: info.TenPay,
	})

	if err != nil {
		c.JSON(500, &APIRSP{
			StatusCode: 200,
			Result:     nil,
			Detail:     err.Error(),
		})
		return
	}
	c.JSON(200, &APIRSP{
		StatusCode: 200,
		Result:     rsp,
		Detail:     "ok",
	})

}

// UpdatePayInfo 用户更新支付信息
func (p *PayController) UpdatePayInfo(c *gin.Context) {
	info := PayInfoRequest{}
	info.UID = c.GetInt64("UID")
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(422, &APIRSP{
			StatusCode: 422,
			Result:     nil,
			Detail:     err.Error(),
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
	c.JSON(200, &APIRSP{
		StatusCode: 200,
		Result:     rsp,
		Detail:     "ok",
	})
}
