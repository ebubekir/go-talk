package infra

import (
	"github.com/ebubekir/go-talk/api/internal/user/domain"
	"github.com/ebubekir/go-talk/api/pkg/mongodb"
)

type MongoDbUserRepository struct {
	db             *mongodb.MongoDb
	collectionName string
}

func NewMongoDbUserRepository(db *mongodb.MongoDb) *MongoDbUserRepository {
	return &MongoDbUserRepository{db: db, collectionName: "Users"}
}

func (m *MongoDbUserRepository) Create(user *domain.User) error {
	return mongodb.InsertOne[domain.User](m.db, m.collectionName, *user)
}

func (m *MongoDbUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	return mongodb.GetOneByField[domain.User](m.db, m.collectionName, "email", email)
}

func (m *MongoDbUserRepository) GetUserById(id string) (*domain.User, error) {
	return mongodb.GetOneByField[domain.User](m.db, m.collectionName, "_id", id)
}
