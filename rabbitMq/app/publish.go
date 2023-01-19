package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/houseofcat/turbocookedrabbit/v2/pkg/tcr"
	logger "github.com/jonny-mark/package-demo/rabbitMq/pkg/log"
	"github.com/jonny-mark/package-demo/rabbitMq/pkg/rabbitMq"
	"log"
)

var (
	exchangeName = "test-exchange"
)

func Publish(c *gin.Context) {
	rabbitMq.NewRabbitService()
	id, err := uuid.NewUUID()
	if err != nil {
		logger.Error(err)
	}
	letter := &tcr.Letter{
		LetterID: id,
		Body:     []byte("Hello World"),
		Envelope: &tcr.Envelope{
			Exchange:     "exchangeName",
			RoutingKey:   "guest-routing-key1",
			ContentType:  "text/plain",
			Mandatory:    false,
			Immediate:    false,
			Priority:     0,
			DeliveryMode: 2,
		},
	}

	rabbitMq.RabbitService.Publisher.Publish(letter, true)
	err = rabbitMq.RabbitService.PublishData([]byte("Hello World"), "exchangeName", "guest-routing-key", nil)
	if err != nil {
		log.Fatal(err)
	}

	consumer, err := rabbitMq.RabbitService.GetConsumer("RabbitConsumerOne")
	if err != nil {
		log.Fatal(err)
	}
	consumer.StartConsuming()

	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}
