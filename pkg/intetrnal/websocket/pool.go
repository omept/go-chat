package websocket

import (
	"log"
	"os"
	"runtime/debug"

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
	defer p.ReviveWebsocket()
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			log.Println("info:", "New client. Size of connection pool:", len(p.Clients))
			for c := range p.Clients {
				err := c.Connection.WriteJSON(Message{Type: 1, Body: Body{ChatMessage: "new user joined..."}})
				errors.ErrorCheck(err)
			}

		case client := <-p.Unregister:
			delete(p.Clients, client)
			log.Println("info:", "disconnected a client. size of connection pool:", len(p.Clients))
			for c := range p.Clients {

				err := c.Connection.WriteJSON(Message{Type: 1, Body: Body{ChatMessage: "user disconnected..."}})
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

func (p *Pool) ReviveWebsocket() {
	if err := recover(); err != nil {
		if os.Getenv("LOG_PANIC_TRACE") == "true" {
			log.Println(
				"level:", "error",
				"err: ", err,
				"trace", string(debug.Stack()),
			)
		} else {
			log.Println(
				"level", "error",
				"err", err,
			)
		}
		go p.Start()
	}
}
