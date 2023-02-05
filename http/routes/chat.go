package routes

import (
	"github.com/gorilla/mux"
	"github.com/ong-gtp/go-chat/controllers"
	"github.com/ong-gtp/go-chat/http/middlewares"
	"github.com/ong-gtp/go-chat/services"
)

var RegisterChatRoutes = func(router *mux.Router) {

	sb := router.PathPrefix("/v1/api/chat").Subrouter()
	sb.Use(middlewares.HeaderMiddleware)
	sb.Use(middlewares.Authenticated)

	var chat controllers.ChatController
	chat.RegisterService(services.NewChatService())

	sb.HandleFunc("/create", chat.Create).Methods("POST")
	sb.HandleFunc("/rooms", chat.ChatRooms).Methods("POST")
	sb.HandleFunc("/room-messages", chat.ChatRoomMessages).Methods("POST")
}
