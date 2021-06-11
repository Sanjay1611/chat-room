package models

import "github.com/google/uuid"

type User struct {
	UID        string
	Name       string
	Age        string
	ChatRoomID string
}

func NewUser() *User {
	return &User{
		UID: uuid.New().String(),
	}
}

func (user *User) SetName(name string) *User {
	user.Name = name
	return user
}

func (user *User) SetAge(age string) *User {
	user.Age = age
	return user
}

func (user *User) SetChatRoomID(chatRoomID string) *User {
	user.ChatRoomID = chatRoomID
	return user
}
