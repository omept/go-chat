package rabbitmq

import (
	"log"
	"os"

	"github.com/ong-gtp/go-chat/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

var br Broker

func InitilizeBroker() (*amqp.Connection, *amqp.Channel) {
	log.Println("RabbitMQ connecting")

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
	log.Println("RabbitMQ connected")
	return conn, ch
}

func GetRabbitMQBroker() *Broker {
	return &br
}
