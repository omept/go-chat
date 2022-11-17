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
	Email      string
}

type Message struct {
	Type   int    `json:"Type,omitempty"`
	Body   string `json:"Body,omitempty"`
	RoomId string `json:"RoomId,omitempty"`
	Email  string `json:"Email,omitempty"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Connection.Close()
	}()

	for {
		messageType, p, err := c.Connection.ReadMessage()
		errors.ErrorCheck(err)
		body := string(p)
		message := Message{Type: messageType, Body: body}
		c.Pool.Broadcast <- message
		log.Println("info:", "Message received: ", body, "messageType: ", messageType)
	}
}
