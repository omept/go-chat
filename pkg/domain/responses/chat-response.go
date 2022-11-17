package responses

import "github.com/ong-gtp/go-chat/pkg/models"

type ChatRoomsResponse struct {
	ChatRooms []models.ChatRoom `json:"ChatRooms"`
}

type ChatRoomMessagesResponse struct {
	Chats []models.Chat `json:"Chats"`
}
