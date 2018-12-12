package main

import (
	"fmt"
	"github.com/micro/go-micro"
	proto "github.com/unliar/proto/pay"
	"learnGo/examples/go-micro-pay/config"
	"time"
)

func main() {
	fmt.Println("start to run main")

	service := micro.NewService(
		micro.Name(config.Config.ServiceName),
		micro.Version(config.Config.ServiceVersion),
		micro.RegisterTTL(config.Config.RegisterTTL*time.Second),
		micro.RegisterInterval(config.Config.RegisterInterval*time.Second),
	)
	service.Init()
	proto.RegisterPaySVHandler(service.Server(), new(Pay))
	err := service.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
