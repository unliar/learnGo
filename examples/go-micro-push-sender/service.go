package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

// SendString 发送一个文本信息
func SendString(s string) error {
	fmt.Println("qaq")
	Q, err := Ch.QueueDeclare("go-push", true, false, false, false, nil)
	fmt.Print("start SendString")
	if err != nil {
		fmt.Print(err)
	}
	err = Ch.Publish("", Q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(s),
	})
	return err
}
