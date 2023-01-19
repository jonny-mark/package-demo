package rabbitMq

// Consumer define consumer for rabbitmq
type Consumer struct {
	routingKey string
	exchange   string
}
