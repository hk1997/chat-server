package main

import (
	broker "chatserver/broker"
	"fmt"
)

func main() {
	// Get a greeting message and print it.
	fmt.Println("hello world this works")

	// rabbitMq.Consumer("hello", func(body []byte) {
	// 	fmt.Println("Received message " + string(body))
	// })

	broker := broker.GetMessagingBroker("amqp://guest:guest@localhost:5672/")

	// broker.Consumer("hello", func(body []byte) {
	// 	fmt.Println("Received message " + string(body))
	// })

	broker.Publish([]string{"hello"}, []byte("Its working"))

}
