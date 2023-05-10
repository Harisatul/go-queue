package main

import (
	"github.com/streadway/amqp"
	"go-queue/event"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	consumer, err := event.NewCustomer(connection)
	if err != nil {
		panic(err)
	}
	list := []string{"printqueue"}
	consumer.Listen(list)
}
