package rabbitMq

import amqp "github.com/rabbitmq/amqp091-go"

type Channel struct {
	ch            *amqp.Channel
	notifyClose   chan *amqp.Error
	notifyConfirm chan amqp.Confirmation
}
