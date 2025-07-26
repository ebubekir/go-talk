package domain

type RoomRepository interface {
	Create(room *Room) error
	Save(room *Room) error
	Delete(roomId string) error
	GetRoomById(roomId string) (*Room, error)
}
