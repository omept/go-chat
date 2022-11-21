package websocket

import (
	"encoding/json"
	"log"
	"strings"

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
	Type int `json:"Type,omitempty"`
	Body Body
}
type Body struct {
	ChatRoomName string `json:"chatRoomName,omitempty"`
	ChatRoomId   int32  `json:"chatRoomId,omitempty"`
	ChatMessage  string `json:"chatMessage,omitempty"`
	ChatUser     string `json:"chatUser,omitempty"`
}

func (c *Client) Read(bodyChan chan []byte) {
	defer func() {
		c.Pool.Unregister <- c
		c.Connection.Close()
	}()
	defer c.Pool.ReviveWebsocket()

	for {
		messageType, p, err := c.Connection.ReadMessage()
		errors.ErrorCheck(err)
		var body Body
		err = json.Unmarshal(p, &body)
		errors.ErrorCheck(err)
		body.ChatUser = c.Email
		message := Message{Type: messageType, Body: body}
		c.Pool.Broadcast <- message
		log.Println("info:", "Message received: ", body, "messageType: ", messageType)

		if strings.Index(body.ChatMessage, "/stock=") == 0 {
			bodyChan <- p
		}
	}
}
