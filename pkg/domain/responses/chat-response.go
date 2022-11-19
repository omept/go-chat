package responses

import "github.com/ong-gtp/go-chat/pkg/models"

type ChatRoomsResponse struct {
	ChatRooms []models.ChatRoom `json:"ChatRooms"`
}
type ChatRoomResponse struct {
	ChatRoom models.ChatRoom `json:"ChatRoom"`
}

type ChatRoomMessagesResponse struct {
	Chats []models.Chat `json:"Chats"`
}
