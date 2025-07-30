package application

import (
	"errors"
	"github.com/ebubekir/go-talk/api/internal/chat/domain"
)

type ChatService struct {
	repo    domain.ChatRepository
	service *domain.ChatService
}

func NewChatService(repo domain.ChatRepository) *ChatService {
	chatService := &domain.ChatService{}
	return &ChatService{repo: repo, service: chatService}
}

func (cs *ChatService) CreateChat(roomId string) (*domain.Chat, error) {
	if isChatExists, err := cs.repo.GetChat(roomId); err != nil {
		return nil, err
	} else if isChatExists != nil {
		return nil, errors.New("chat already exists")
	}

	if chat, err := cs.service.Create(roomId); err != nil {
		return nil, err
	} else {
		if err = cs.repo.Create(chat); err != nil {
			return nil, err
		}
		return chat, nil
	}
}

func (cs *ChatService) GetChat(roomId string) (*domain.Chat, error) {
	chat, err := cs.repo.GetChat(roomId)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		chat, err = cs.CreateChat(roomId)
		if err != nil {
			return nil, err
		}
	}

	return chat, nil
}

func (cs *ChatService) SendMessage(roomId string, userId string, message string) error {
	chat, err := cs.GetChat(roomId)
	if err != nil {
		return err
	}
	chat.AddMessage(userId, message)
	return cs.repo.Save(chat)
}
