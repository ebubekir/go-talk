package domain

type User struct {
	Id         string `bson:"_id"`
	FirebaseId string
	Name       string
	Email      string
	IsDeleted  bool
}
