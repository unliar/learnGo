package main

import "C"
import (
	"fmt"
	"github.com/micro/go-micro"
	proto "github.com/unliar/proto/account"
	"learnGo/examples/go-micro-account/config"
	SV "learnGo/examples/go-micro-account/service"
	"learnGo/examples/go-micro-account/utils"
	"time"
)

func main() {
	// 配置初始化
	config.InitConfig()

	service := micro.NewService(
		micro.Name(config.Config.ServiceName),
		micro.Version(config.Config.Env+"-"+config.Config.ServiceVersion),
		micro.RegisterInterval(15*time.Second),
		micro.RegisterTTL(time.Second*60),
		micro.WrapCall(utils.MicroWrapCall),
		micro.WrapHandler(utils.MicroWrapHandler),
	)

	service.Init()

	proto.RegisterAccountSVHandler(service.Server(), new(SV.Account))

	err := service.Run()

	if err != nil {
		fmt.Println(err)
	}
}

// CONSUL_HTTP_ADDR=$consuleader go run main.go  --registry=consul --registry_address=$consuleader--selector=cache --server=grpc --client=grpc
