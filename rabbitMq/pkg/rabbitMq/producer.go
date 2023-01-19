package rabbitMq

import (
	"context"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Producer struct {
	channel *amqp.Channel
}

func (p *Producer) Publish(ctx context.Context, exchange, routingKey, msg string) error {
	return p.channel.PublishWithContext(
		ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain", //纯文本
			DeliveryMode: amqp.Persistent,
			MessageId:    uuid.New().String(),
			Timestamp:    time.Now(),
			Type:         "",
			Body:         []byte(msg),
		},
	)
}
