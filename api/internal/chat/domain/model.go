package domain

import "time"

type ChatMessage struct {
	UserId string
	Text   string
	SentAt time.Time
}

type Chat struct {
	Id           string `bson:"_id"`
	RoomId       string
	ChatMessages []ChatMessage
}

func (c *Chat) AddMessage(userId string, text string) {
	c.ChatMessages = append(c.ChatMessages, ChatMessage{
		UserId: userId,
		Text:   text,
		SentAt: time.Now(),
	})
}
