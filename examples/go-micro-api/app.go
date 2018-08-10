package main

import "github.com/micro/go-web"
func main()  {
	service:=web.NewService(
		web.Name("unliar-api"),
		)
}