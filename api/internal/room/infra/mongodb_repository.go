package infra

import (
	"github.com/ebubekir/go-talk/api/internal/room/domain"
	"github.com/ebubekir/go-talk/api/pkg/mongodb"
)

type MongoDbRoomRepository struct {
	db             *mongodb.MongoDb
	collectionName string
}

func NewMongoDbRoomRepository(db *mongodb.MongoDb) *MongoDbRoomRepository {
	return &MongoDbRoomRepository{db: db, collectionName: "Rooms"}
}

func (m *MongoDbRoomRepository) Create(room *domain.Room) error {
	return mongodb.InsertOne[domain.Room](m.db, m.collectionName, *room)
}

func (m *MongoDbRoomRepository) GetRoomById(id string) (*domain.Room, error) {
	return mongodb.GetOneByField[domain.Room](m.db, m.collectionName, "_id", id)
}

func (m *MongoDbRoomRepository) Save(room *domain.Room) error {
	return mongodb.UpdateOne[domain.Room](m.db, m.collectionName, string(room.Id), room)
}

func (m *MongoDbRoomRepository) Delete(roomId string) error {
	return mongodb.DeleteOne(m.db, m.collectionName, roomId)
}
