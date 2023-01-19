package rabbitMq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Connection struct {
	tempId int64
	active map[int64]int64
	conn   *amqp.Connection
	err    *amqp.Error
}

// NewConnection connect to rabbitmq
func NewConnection(config *Config) (*Connection, error) {
	//guest:guest@localhost:5672
	uri := fmt.Sprintf("amqp://%s:%s@%s", config.User, config.Password, config.Addr)

	conn, err := amqp.DialConfig(uri, amqp.Config{
		Vhost:      config.Vhost,
		ChannelMax: config.ChannelMax,
		Heartbeat:  config.Heartbeat,
	})
	if err != nil {
		return nil, err
	}
	return &Connection{
		conn:   conn,
		active: make(map[int64]int64),
	}, err
}
