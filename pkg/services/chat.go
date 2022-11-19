package services

import (
	"github.com/ong-gtp/go-chat/pkg/domain/responses"
	"github.com/ong-gtp/go-chat/pkg/errors"
	"github.com/ong-gtp/go-chat/pkg/models"
)

type ChatService interface {
	CreateChatRoom(name string) (responses.ChatRoomResponse, error)
	ChatRooms() (responses.ChatRoomsResponse, error)
	ChatRoomMessages(roomId uint) (responses.ChatRoomMessagesResponse, error)
}

type chat struct{}

func NewChatService() *chat {
	return &chat{}
}

func (c *chat) ChatRooms() (responses.ChatRoomsResponse, error) {
	var chtList []models.ChatRoom
	var chtRoomModel models.ChatRoom
	err := chtRoomModel.List(&chtList)
	errors.DBErrorCheck(err)
	return responses.ChatRoomsResponse{ChatRooms: chtList}, nil
}

func (c *chat) CreateChatRoom(name string) (responses.ChatRoomResponse, error) {
	cht := models.ChatRoom{
		Name: name,
	}
	err := cht.Add()
	errors.DBErrorCheck(err)
	return responses.ChatRoomResponse{ChatRoom: cht}, nil
}

func (c *chat) ChatRoomMessages(roomId uint) (responses.ChatRoomMessagesResponse, error) {
	return responses.ChatRoomMessagesResponse{Chats: []models.Chat{}}, nil
}
