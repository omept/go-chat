package services

import (
	"log"

	"github.com/ong-gtp/go-chat/pkg/domain/responses"
	"github.com/ong-gtp/go-chat/pkg/errors"
	"github.com/ong-gtp/go-chat/pkg/models"
)

type ChatService interface {
	CreateChatRoom(name string) (responses.ChatRoomResponse, error)
	SaveChatMessage(msg string, roomId, userId uint) bool
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
	var chtList []models.Chat
	chtMsgList := []responses.ChatMessage{}
	var chtModel models.Chat
	err := chtModel.List(roomId, &chtList)
	errors.DBErrorCheck(err)
	// transform chats
	for _, v := range chtList {
		chtMsgList = append(chtMsgList, responses.ChatMessage{
			ChatMessage:  v.Message,
			ChatUser:     v.User.Email,
			ChatRoomId:   v.ChatRoomId,
			ChatRoomName: v.ChatRoom.Name,
		})
	}
	return responses.ChatRoomMessagesResponse{Chats: chtMsgList}, nil
}

func (c *chat) SaveChatMessage(msg string, roomId, userId uint) bool {
	cht := models.Chat{
		Message:    msg,
		UserId:     userId,
		ChatRoomId: roomId,
	}
	err := cht.Add()
	if err.Error != nil {
		log.Println("error: ", err.Error)
		return false
	}

	return true
}
