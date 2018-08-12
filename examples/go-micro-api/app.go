package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-web"
	"log"
	ac "learnGo/examples/go-micro-api/account"
)

func main() {
	var err error
	
	service := web.NewService(
		web.Name("unliar-restful-api"),
	)
	
	service.Init()

	r := gin.Default()
	
	
	
	r.GET("/api/user/:uid", ac.GetUserBase)
	
	service.Handle("/", r)
	
	err = service.Run()
	
	if err != nil {
		log.Fatal(err)
	}
}
