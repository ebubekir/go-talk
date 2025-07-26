package application

import (
	"github.com/ebubekir/go-talk/api/internal/room/domain"
)

type RoomService struct {
	repo    domain.RoomRepository
	service *domain.RoomService
}

func NewRoomService(repo domain.RoomRepository) *RoomService {
	roomService := &domain.RoomService{}
	return &RoomService{repo: repo, service: roomService}
}

func (rs *RoomService) Create(createdByUserId string, isPrivateRoom bool) (*domain.Room, error) {
	return rs.service.Create(createdByUserId, isPrivateRoom)
}

func (rs *RoomService) JoinRoom(userId string, roomId string) error {
	room, err := rs.repo.GetRoomById(roomId)
	if err != nil {
		return err
	}
	err = room.AddParticipant(domain.UserId(userId))
	if err != nil {
		return err
	}

	_ = rs.repo.Save(room)

	// Emit event
	//event := domain.RoomEvent{
	//	Type:      domain.EventParticipantJoined,
	//	RoomId:    domain.RoomId(roomId),
	//	Timestamp: time.Now(),
	//	Payload: map[string]any{
	//		"user_id": userId,
	//	},
	//}

	return nil
}
