package websocket

import (
	"log"

	"github.com/ong-gtp/go-chat/pkg/errors"
	"github.com/ong-gtp/go-chat/pkg/utils"
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
	defer utils.Revive()
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			log.Println("info:", "New client. Size of connection pool:", len(p.Clients))
			for c := range p.Clients {
				err := c.Connection.WriteJSON(Message{Type: 1, Body: "new user joined..."})
				errors.ErrorCheck(err)
			}

		case client := <-p.Unregister:
			delete(p.Clients, client)
			log.Println("info:", "disconnected a client. size of connection pool:", len(p.Clients))
			for c := range p.Clients {
				err := c.Connection.WriteJSON(Message{Type: 1, Body: "user disconnected..."})
				errors.ErrorCheck(err)
			}

		case msg := <-p.Broadcast:
			log.Println("info", "broadcast message to clients in pool")
			for c := range p.Clients {
				err := c.Connection.WriteJSON(msg)
				errors.ErrorCheck(err)
			}
		}
	}
}
