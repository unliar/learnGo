package main

import (
	"fmt"
	"github.com/micro/go-micro"
	proto "github.com/unliar/proto/account"
	SV "learnGo/examples/go-micro-account/service"
	"learnGo/examples/go-micro-account/utils"
	"time"
)

func main() {
	service := micro.NewService(
		micro.Name("unliar-account"),
		micro.Version("beta-1.1.1"),
		micro.RegisterInterval(time.Second*15),
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
