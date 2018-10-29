package main

import (
	"fmt"
	"github.com/micro/go-micro"
	proto "github.com/unliar/proto/pay"
	"time"
)

func main() {
	fmt.Println("start to run main")

	service := micro.NewService(
		micro.Name("unliar-pay"),
		micro.Version("v1.0.0.1"),
		micro.RegisterTTL(time.Second*60),
		micro.RegisterInterval(time.Second*15),
	)
	service.Init()
	proto.RegisterPaySVHandler(service.Server(), new(Pay))
	err := service.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
