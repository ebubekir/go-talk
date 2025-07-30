package domain

type ChatService struct{}

func (cs *ChatService) Create(roomId string) (*Chat, error) {
	return &Chat{
		Id:           roomId,
		RoomId:       roomId,
		ChatMessages: make([]ChatMessage, 0),
	}, nil
}
