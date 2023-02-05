package requests

type ChatRoomMessagesPayload struct {
	RoomId uint `json:"RoomId"`
}
type ChatRoomCreatePayload struct {
	Name string `json:"Name"`
}
