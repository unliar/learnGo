package main

import (
	"context"
)
import proto "github.com/unliar/proto/pay"

type Pay struct {
}

func (p *Pay) GetPayInfo(ctx context.Context, req *proto.PayInfo, rsp *proto.ResponseStatus) error {
	result := &PayInfo{UID: req.GetUID()}
	if r := DB.Where(result).First(result).RowsAffected; r > 0 {
		rsp.Status = 1
		rsp.ErrMsg = "ok"
		rsp.PayInfo = &proto.PayInfo{
			UID:    result.UID,
			TenPay: result.TenPay,
			Alipay: result.Alipay,
		}
		return nil
	}
	rsp.Status = 2
	rsp.ErrMsg = "no record"
	return nil
}

func (p *Pay) PostPayInfo(ctx context.Context, req *proto.PayInfo, rsp *proto.ResponseStatus) error {
	result := &PayInfo{UID: req.UID}

	if err := DB.Where(result).FirstOrCreate(&PayInfo{
		UID:    req.UID,
		TenPay: req.TenPay,
		Alipay: req.Alipay}).First(result).Error; err != nil {
		rsp.Status = 2
		rsp.ErrMsg = err.Error()
		return nil
	}

	rsp.Status = 1
	rsp.ErrMsg = "ok"
	rsp.PayInfo = &proto.PayInfo{
		UID:    result.UID,
		TenPay: result.TenPay,
		Alipay: result.Alipay,
	}
	return nil
}

func (p *Pay) UpdatePayInfo(ctx context.Context, req *proto.PayInfo, rsp *proto.ResponseStatus) error {
	if err := DB.Model(&PayInfo{}).Where(&PayInfo{UID: req.UID}).Updates(map[string]interface{}{
		"alipay":  req.Alipay,
		"ten_pay": req.TenPay,
	}).Error; err != nil {
		rsp.Status = 2
		rsp.ErrMsg = err.Error()
		return nil
	}
	rsp.Status = 1
	rsp.ErrMsg = "ok"
	return nil
}
