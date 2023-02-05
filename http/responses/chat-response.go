package responses

import "github.com/ong-gtp/go-chat/models"

type ChatRoomsResponse struct {
	ChatRooms []models.ChatRoom `json:"ChatRooms"`
}
type ChatRoomResponse struct {
	ChatRoom models.ChatRoom `json:"ChatRoom"`
}

type ChatMessage struct {
	ChatMessage  string `json:"chatMessage"`
	ChatUser     string `json:"chatUser"`
	ChatRoomId   uint   `json:"chatRoomId"`
	ChatRoomName string `json:"chatRoomName"`
}

type ChatRoomMessagesResponse struct {
	Chats []ChatMessage `json:"Chats"`
}
