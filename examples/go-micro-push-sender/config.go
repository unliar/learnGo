package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

var Ch *amqp.Channel

func init() {
	fmt.Println("config starting....")
	var err error
	connection, err := amqp.Dial("amqp://unliar:19930224@localhost:5672/")
	if err != nil {
		fmt.Printf("%s====%s", "fail connect mq server ", err)
	}
	fmt.Println("config starting....connection.Channel")
	Ch, err = connection.Channel()
	if err != nil {
		fmt.Printf("%s===%s", "chanel error", err)
	}

}
