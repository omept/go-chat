package rabbitmq

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ong-gtp/go-chat/pkg/errors"
	"github.com/ong-gtp/go-chat/pkg/intetrnal/websocket"
	"github.com/ong-gtp/go-chat/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type StockRequest struct {
	RoomId uint   `json:"RoomId"`
	Code   string `json:"Code"`
}

type StockResponse struct {
	RoomId  uint   `json:"RoomId"`
	Message string `json:"Message"`
}

type Broker struct {
	ReceiverQueue  amqp.Queue
	PublisherQueue amqp.Queue
	Channel        *amqp.Channel
}

// Setup creates(or connects if not existing) the reciever and publisher queues
func (b *Broker) SetUp(ch *amqp.Channel) {
	receiverQueue := os.Getenv("STKBT_RECEIVER_QUEUE")
	publisherQueue := os.Getenv("STKBT_PUBLISHER_QUEUE")

	q1, err := ch.QueueDeclare(
		receiverQueue, // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	errors.ErrorCheck(err)

	q2, err := ch.QueueDeclare(
		publisherQueue, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	errors.ErrorCheck(err)

	b.ReceiverQueue = q1
	b.PublisherQueue = q2
	b.Channel = ch
}

// PublishMessage sends messages to the stock-bot's receiver queue
func (b *Broker) PublishMessage(requestBody chan []byte) {
	for body := range requestBody {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		err := b.Channel.PublishWithContext(ctx,
			"",                    // exchange
			b.PublisherQueue.Name, // routing key
			false,                 // mandatory
			false,                 // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
		cancel()
		if err != nil {
			log.Printf("PublishMessage Error occured %s\n", err)
			continue
		}
		log.Printf(" [x] Sent %s\n", body)
	}
}

// ReadMessages reads messages from the stock-bot's publisher queue
func (b *Broker) ReadMessages(pool *websocket.Pool) {
	msgs, err := b.Channel.Consume(
		b.ReceiverQueue.Name, // queue
		"",                   // consumer
		true,                 // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	if err != nil {
		log.Printf("ReadMessages Error occured %s\n", err)
		return
	}

	rsvdMsgs := make(chan StockResponse)
	go messageTransformer(msgs, rsvdMsgs)
	go processResponse(rsvdMsgs, b, pool)
	select {}
}

// messageTransformer converts the message from rabbitmq to StockResponse type and passes it to the received messages channel
func messageTransformer(entries <-chan amqp.Delivery, receivedMessages chan StockResponse) {
	var sr StockResponse
	for d := range entries {
		err := utils.ParseByteArray(d.Body, &sr)
		if err != nil {
			log.Printf("Received bad response : %s ", string(d.Body))
			continue
		}
		log.Println("Received a response")
		receivedMessages <- sr
	}
}

// processResponse sends the stock response to the websocket's connection pool
func processResponse(s <-chan StockResponse, b *Broker, pool *websocket.Pool) {
	for r := range s {
		log.Println("processing stock response for ", r.RoomId)
		// body, err := json.Marshal(r)
		// errors.ErrorCheck(err)

		sr := StockResponse{
			RoomId:  r.RoomId,
			Message: r.Message,
		}

		message := websocket.Message{Type: 1, Body: websocket.Body{ChatRoomId: int32(sr.RoomId), ChatUser: "stock-bot", ChatMessage: sr.Message}}
		pool.Broadcast <- message
		log.Println("processed", sr)
	}
}
