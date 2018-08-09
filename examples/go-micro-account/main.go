package main

import (
	"fmt"
	"github.com/micro/go-micro"
	proto "learnGo/examples/go-micro-account/proto"
	SV "learnGo/examples/go-micro-account/service"
	"time"
)

func main() {
	var err error
	service := micro.NewService(
		micro.Name("unliar-account"),
		micro.Version("beta-1.0.1"),
		micro.RegisterInterval(time.Second*15),
		micro.RegisterTTL(time.Second*30),
	)
	service.Init()
	proto.RegisterAccountSVHandler(service.Server(), new(SV.Account))
	err = service.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// CONSUL_HTTP_ADDR=$consuleader go run main.go  --registry=consul --registry_address=$consuleader--selector=cache --server=grpc --client=grpc