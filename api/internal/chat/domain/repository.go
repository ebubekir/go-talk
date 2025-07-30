package domain

type ChatRepository interface {
	Create(chat *Chat) error
	GetChat(roomId string) (*Chat, error)
	Save(chat *Chat) error
}
