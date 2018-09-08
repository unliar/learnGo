package main

import (
	"fmt"
	SV "learnGo/examples/go-micro-account/service"
	"time"

	"github.com/micro/go-micro"
	proto "github.com/unliar/proto/account"
)

func main() {
	var err error
	service := micro.NewService(
		micro.Name("unliar-account"),
		micro.Version("beta-1.1.1"),
		micro.RegisterInterval(time.Second*15),
		micro.RegisterTTL(time.Second*15),
	)
	service.Init()
	proto.RegisterAccountSVHandler(service.Server(), new(SV.Account))
	err = service.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// CONSUL_HTTP_ADDR=$consuleader go run main.go  --registry=consul --registry_address=$consuleader--selector=cache --server=grpc --client=grpc
