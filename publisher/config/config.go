package config

import (
	"github.com/streadway/amqp"
	"go-queue/event"
)

func CreateEventEmitter() (event.Emitter, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		return event.Emitter{}, err
	}

	emitter, err := event.NewEventEmitter(conn)
	if err != nil {
		return emitter, err
	}

	return emitter, nil
}
