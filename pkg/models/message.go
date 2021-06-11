package models

import "github.com/google/uuid"

type Message struct {
	UID     string
	Content string
	UserID  string
}

func NewMessage() *Message {
	return &Message{
		UID: uuid.New().String(),
	}
}

func (message *Message) SetContent(content string) *Message {
	message.Content = content
	return message
}

func (message *Message) SetUserID(userID string) *Message {
	message.UserID = userID
	return message
}
