package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"time"
)

type MongoDb struct {
	ConnectionString string
	DbName           string
	Client           *mongo.Client
}

func New(connectionString, DbName string) *MongoDb {
	return &MongoDb{
		ConnectionString: connectionString,
		DbName:           DbName,
		Client:           nil,
	}
}

func (m *MongoDb) CheckConnection() error {
	_, err := m.getClient()
	return err
}

func (m *MongoDb) getClient() (*mongo.Client, error) {
	if m.Client != nil {
		return m.Client, nil
	}

	var err error
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(m.ConnectionString).SetServerAPIOptions(serverAPI)
	m.Client, err = mongo.Connect(opts)
	return m.Client, err
}

func (m *MongoDb) Close() error {
	if m.Client == nil {
		return nil
	}

	err := m.Client.Disconnect(context.Background())

	if err != nil {
		log.Println(err)
	}

	m.Client = nil
	return err
}

func getCollection(name string, db *MongoDb) (*mongo.Collection, error) {
	client, err := db.getClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(db.DbName).Collection(name)
	return collection, nil
}

func SetValue(db *MongoDb, collectionName string, id string, fieldName string, value any) error {
	client, err := db.getClient()
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{fieldName: value}}

	collection := client.Database(db.DbName).Collection(collectionName)

	// Update one document
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	if err != nil {
		return err
	}

	return err
}

func GetOneByField[T any](db *MongoDb, collectionName, searchField string, searchValue any) (*T, error) {
	collection, err := getCollection(collectionName, db)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.D{{searchField, searchValue}}

	var result *T
	err = collection.FindOne(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	return result, err
}

func GetOneById[T any](db *MongoDb, collectionName string, id string) (*T, error) {
	collection, err := getCollection(collectionName, db)
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.D{{"_id", id}}

	var result *T
	err = collection.FindOne(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	return result, err
}

func InsertOne[T any](db *MongoDb, collectionName string, record T) error {
	collection, err := getCollection(collectionName, db)
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err = collection.InsertOne(ctx, record)
	return err
}

func UpdateOne[T any](db *MongoDb, collectionName string, id string, record *T) error {
	collection, err := getCollection(collectionName, db)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", record}}
	opts := options.UpdateOne()

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	_, err = collection.UpdateOne(ctx, filter, update, opts) //result
	if err != nil {
		return err
	}

	return nil
}

func UpsertOne[T any](db *MongoDb, collectionName string, id string, record *T) error {
	collection, err := getCollection(collectionName, db)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", record}}
	opts := options.UpdateOne().SetUpsert(true)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err = collection.UpdateOne(ctx, filter, update, opts) //result
	if err != nil {
		return err
	}

	return nil
}

func Query[T any](db *MongoDb, collectionName string, filter interface{}, opts *options.FindOptionsBuilder) ([]T, error) {
	collection, err := getCollection(collectionName, db)
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)

	if cursor, err := collection.Find(ctx, filter, opts); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	} else {
		var result []T
		if err = cursor.All(ctx, &result); err != nil {
			return nil, err
		}

		if result == nil {
			return []T{}, nil
		}

		return result, nil
	}
}

func GetAll[T any](db *MongoDb, collectionName string) ([]T, error) {
	collection, err := getCollection(collectionName, db)
	if err != nil {
		return nil, err
	}

	filter := bson.D{}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var result []T
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteOne(db *MongoDb, collectionName string, id string) error {
	collection, err := getCollection(collectionName, db)
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//defer client.Disconnect(ctx)

	filter := bson.D{{"_id", id}}

	_, err = collection.DeleteOne(ctx, filter)
	return err
}

func DeleteAll(db *MongoDb, collectionName string, filter interface{}) error {
	collection, err := getCollection(collectionName, db)
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	_, err = collection.DeleteMany(ctx, filter)
	return err
}

func Count(db *MongoDb, collectionName string, filter interface{}, opts *options.CountOptionsBuilder) (int64, error) {
	collection, err := getCollection(collectionName, db)
	if err != nil {
		return 0, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if count, err := collection.CountDocuments(ctx, filter, opts); err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}
		return -1, err
	} else {
		return count, nil
	}
}
