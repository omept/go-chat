package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ong-gtp/go-chat/pkg/errors"
	"github.com/ong-gtp/go-chat/pkg/intetrnal/websocket"
)

var RegisterWebsocketRoute = func(router *mux.Router) {
	pool := websocket.NewPool()
	go pool.Start()
	sb := router.PathPrefix("/v1").Subrouter()

	sb.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})

}

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	errors.ErrorCheck(err)

	client := &websocket.Client{
		Connection: conn,
		Pool:       pool,
	}

	pool.Register <- client
	client.Read()
}
