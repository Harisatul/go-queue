package event

import (
	"github.com/streadway/amqp"
	"log"
)

type Emitter struct {
	connection *amqp.Connection
}

func (e *Emitter) setup() error {
	channel, err := e.connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()
	return declareExchange(channel)
}

func (e *Emitter) Push(numbers []int, severity string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	byteData := make([]byte, len(numbers))
	for i, num := range numbers {
		byteData[i] = byte(num)
	}

	channel.Publish(getExchangeName(), severity, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        byteData,
	})
	log.Printf("Sending message : %v -> %s", numbers, getExchangeName())
	return nil
}

func NewEventEmitter(conn *amqp.Connection) (Emitter, error) {
	emitter := Emitter{
		connection: conn,
	}

	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}
	return emitter, err
}
