package event

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"go-queue/worker/sender"
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

func (e *Emitter) Push(patients sender.Patient, severity string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	jsonData, err := json.Marshal(patients)
	if err != nil {
		return err
	}

	channel.Publish(getExchangeName(), severity, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        jsonData,
	})
	log.Printf("Sending message : %v -> %s", patients, getExchangeName())
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
