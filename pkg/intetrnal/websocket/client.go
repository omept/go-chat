package websocket

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/ong-gtp/go-chat/pkg/errors"
)

type Client struct {
	ID         string
	Connection *websocket.Conn
	Pool       *Pool
	// mutex      sync.Mutex
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Connection.Close()
	}()

	for {
		messageType, p, err := c.Connection.ReadMessage()
		body := string(p)

		errors.ErrorCheck(err)
		message := Message{Type: messageType, Body: body}
		c.Pool.Broadcast <- message
		log.Println("info", "Message received: ", body)
	}
}
