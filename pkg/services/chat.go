package services

import "github.com/ong-gtp/go-chat/pkg/domain/responses"

type ChatService interface {
	ChatRooms() (responses.ChatRoomsResponse, error)
	ChatRoomMessages() (responses.ChatRoomMessagesResponse, error)
}

type chat struct{}

func NewChatService() *chat {
	return &chat{}
}

func (c *chat) ChatRooms() (responses.ChatRoomsResponse, error) {
	return responses.ChatRoomsResponse{}, nil
}

func (c *chat) ChatRoomMessages() (responses.ChatRoomMessagesResponse, error) {
	return responses.ChatRoomMessagesResponse{}, nil
}
