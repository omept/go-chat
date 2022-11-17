package services

import (
	"github.com/ong-gtp/go-chat/pkg/domain/responses"
	"github.com/ong-gtp/go-chat/pkg/models"
)

type ChatService interface {
	ChatRooms() (responses.ChatRoomsResponse, error)
	ChatRoomMessages(roomId uint) (responses.ChatRoomMessagesResponse, error)
}

type chat struct{}

func NewChatService() *chat {
	return &chat{}
}

func (c *chat) ChatRooms() (responses.ChatRoomsResponse, error) {
	return responses.ChatRoomsResponse{ChatRooms: []models.ChatRoom{}}, nil
}

func (c *chat) ChatRoomMessages(roomId uint) (responses.ChatRoomMessagesResponse, error) {
	return responses.ChatRoomMessagesResponse{Chats: []models.Chat{}}, nil
}
