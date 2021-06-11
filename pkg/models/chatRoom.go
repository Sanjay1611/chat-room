package models

import "github.com/google/uuid"

type ChatRoom struct {
	UID  string
	Name string
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		UID: uuid.New().String(),
	}
}

func (chatRoom *ChatRoom) SetName(name string) *ChatRoom {
	chatRoom.Name = name
	return chatRoom
}
