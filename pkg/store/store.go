package store

import "github.com/sanjay/roam/pkg/models"

var (
	ChatRooms map[string]models.ChatRoom = make(map[string]models.ChatRoom)
	Users     map[string]models.User     = make(map[string]models.User)
	Messages  map[string]models.Message  = make(map[string]models.Message)
)

func AddChatRoom(chatRoom *models.ChatRoom) {
	ChatRooms[chatRoom.UID] = *chatRoom
}

func AddUser(user *models.User) {
	Users[user.UID] = *user
}

func AddMessage(message *models.Message) {
	Messages[message.UID] = *message
}
