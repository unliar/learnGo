package main

import (
	"fmt"
	"time"
)

// 超时机制
func main() {

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	select {
	case <-timeout:
		fmt.Println("this is what i need")
	default:
		fmt.Println("this is defualt value")
	}
}
