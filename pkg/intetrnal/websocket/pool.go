package websocket

import (
	"log"

	"github.com/ong-gtp/go-chat/pkg/errors"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			log.Println("info", "size of connection pool:", len(p.Clients))
			for c := range p.Clients {
				log.Println("info", "new client :", c)
				client.Connection.ReadJSON(Message{Body: "new user joined..."})
			}

		case client := <-p.Unregister:
			delete(p.Clients, client)
			client.Connection.ReadJSON(Message{Body: "user disconnected..."})
			log.Println("info", "size of connection pool:", len(p.Clients))

		case msg := <-p.Broadcast:
			log.Println("info", "broadcast message to clients in pool")
			for c := range p.Clients {
				err := c.Connection.WriteJSON(msg)
				errors.ErrorCheck(err)
			}
		}
	}
}
