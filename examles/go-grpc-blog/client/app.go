package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learnGo/examles/go-grpc-blog/proto/Account"
	"learnGo/examles/go-grpc-blog/proto/Message"
	"net/http"
)

const address = "localhost:8081"
const defaultName = "HelloService"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintln(w, err)
	}
	defer conn.Close()
	h := Hello.NewHelloServiceClient(conn)
	hr, err := h.GetRes(context.Background(), &Hello.Req{})
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, hr)

}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintln(w, err)
	}
	defer conn.Close()

	u := User.NewUserServiceClient(conn)
	ur, err := u.GetUser(context.Background(), &User.UserModel{ID: 1, Sex: "222", Name: "555", Age: 56, Summary: "55"})
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, ur)
}

func main() {

	http.HandleFunc("/HelloService", HelloHandler)
	http.HandleFunc("/UserService", UserHandler)
	fmt.Println("server starting")
	http.ListenAndServe(":8082", nil)

}
