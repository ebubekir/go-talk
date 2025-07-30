package infra

import (
	"github.com/ebubekir/go-talk/api/internal/chat/domain"
	"github.com/ebubekir/go-talk/api/pkg/mongodb"
)

type MongoDBChatRepository struct {
	db             *mongodb.MongoDb
	collectionName string
}

func NewMongoDBChatRepository(db *mongodb.MongoDb) *MongoDBChatRepository {
	return &MongoDBChatRepository{db: db, collectionName: "Chats"}
}

func (m *MongoDBChatRepository) Create(chat *domain.Chat) error {
	return mongodb.InsertOne[domain.Chat](m.db, m.collectionName, *chat)
}

func (m *MongoDBChatRepository) GetChat(roomId string) (*domain.Chat, error) {
	return mongodb.GetOneByField[domain.Chat](m.db, m.collectionName, "roomid", roomId)
}

func (m *MongoDBChatRepository) Save(chat *domain.Chat) error {
	return mongodb.UpdateOne[domain.Chat](m.db, m.collectionName, chat.Id, chat)
}
