package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"learnGo/examles/go-grpc-blog/proto/Account"
	"learnGo/examles/go-grpc-blog/proto/Message"
	"log"
	"net"
)

const port = ":8081"

type server struct {
}

func (s *server) GetRes(ctx context.Context, req *Hello.Req) (*Hello.Res, error) {
	return &Hello.Res{Id: req.Id}, nil
}
func (s *server) GetUser(ctx context.Context, req *User.UserModel) (*User.UserModel, error) {
	return &User.UserModel{ID: 1, Name: "yoyo", Sex: "ok", Age: 12, Summary: "2222"}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("server run in %s", port)
	s := grpc.NewServer()
	Hello.RegisterHelloServiceServer(s, &server{})
	User.RegisterUserServiceServer(s, &server{})
	s.Serve(lis)

}
