package main

import (
	"fmt"
	"github.com/micro/go-micro"
	proto "github.com/unliar/proto/account"
	SV "learnGo/examples/go-micro-account/src/service"
	"learnGo/examples/go-micro-account/src/utils"
	"time"
)

func main() {
	InitConfig()
	service := micro.NewService(
		micro.Name(Config.ServiceName),
		micro.Version(Config.Env+"-"+Config.ServiceVersion),
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
