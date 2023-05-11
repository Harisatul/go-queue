package event

import "github.com/streadway/amqp"

func getExchangeName() string {
	return "logs_topic"
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
}

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		getExchangeName(),
		"topic",
		false,
		false,
		false,
		false,
		nil)
}
