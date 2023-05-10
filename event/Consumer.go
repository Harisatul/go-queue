package event

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	return declareExchange(channel)
}

func NewCustomer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}
	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}
	return consumer, nil
}

func (consumer *Consumer) Listen(topics []string) error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	queue, err := declareRandomQueue(channel)
	if err != nil {
		return err
	}

	for _, s := range topics {
		err := channel.QueueBind(queue.Name,
			s,
			getExchangeName(),
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}
	msgs, err := channel.Consume(queue.Name, "", true,
		false, false, false, nil)

	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {

		for d := range msgs {
			str := fmt.Sprintf("Received a message: %v", d.Body)
			log.Println(str)
		}
	}()
	log.Printf("[*] Waiting for message [Exchange, Queue][%s, %s]. To exit press CTRL+C", getExchangeName(), queue.Name)
	<-forever
	return nil
}
