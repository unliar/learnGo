package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"time"
)

// MicroWrapCall 是用来包装rpc client的
func MicroWrapCall(c client.CallFunc) client.CallFunc {
	return func(ctx context.Context, address string, req client.Request, rsp interface{}, opts client.CallOptions) error {
		t := time.Now()
		err := c(ctx, address, req, rsp, opts)
		fmt.Printf("HiWrapCall==> %s -  %s - %s - %v ", address, req.Service(), req.Method(), time.Since(t))
		return err
	}
}

// LoggerInfo 是service的请求信息
type LoggerInfo struct {
	Method  interface{} `json:"method"`
	Payload interface{} `json:"payload"`
	Service interface{} `json:"service"`
	Spend   interface{} `json:"spend"`
	Meta    interface{} `json:"meta"`
}

// MicroWrapHandler 是用来包装service的
func MicroWrapHandler(s server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		t := time.Now()
		metaData, _ := metadata.FromContext(ctx)
		err := s(ctx, req, rsp)
		logger := LoggerInfo{
			Method:  req.Method(),
			Payload: req.Request(),
			Service: req.Service(),
			Spend:   time.Since(t).String(),
			Meta:    metaData,
		}
		s, _ := json.Marshal(logger)

		fmt.Printf("%s", s)
		return err
	}
}
