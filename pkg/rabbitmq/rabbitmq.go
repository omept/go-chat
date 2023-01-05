package rabbitmq

import (
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/ong-gtp/go-chat/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

var br Broker

func InitilizeBroker(logger log.Logger) (*amqp.Connection, *amqp.Channel) {
	level.Info(logger).Log("RabbitMQ ", "connecting")

	rmqHost := os.Getenv("RMQ_HOST")
	rmqUserName := os.Getenv("RMQ_USERNAME")
	rmqPassword := os.Getenv("RMQ_PASSWORD")
	rmqPort := os.Getenv("RMQ_PORT")
	dsn := "amqp://" + rmqUserName + ":" + rmqPassword + "@" + rmqHost + ":" + rmqPort + "/"

	conn, err := amqp.Dial(dsn)
	errors.ErrorCheck(err)

	ch, err := conn.Channel()
	errors.ErrorCheck(err)

	br.SetUp(ch)
	level.Info(logger).Log("RabbitMQ ", "connected")
	return conn, ch
}

func GetRabbitMQBroker() *Broker {
	return &br
}
