package application

import (
	"github.com/ebubekir/go-talk/api/internal/room/domain"
	userApp "github.com/ebubekir/go-talk/api/internal/user/application"
	"github.com/ebubekir/go-talk/api/internal/websocket"
)

type RoomService struct {
	repo        domain.RoomRepository
	service     *domain.RoomService
	userService userApp.UserService
	roomHub     *websocket.Hub
}

func NewRoomService(repo domain.RoomRepository, hub *websocket.Hub, userService *userApp.UserService) *RoomService {
	roomService := &domain.RoomService{}
	return &RoomService{repo: repo, service: roomService, roomHub: hub, userService: *userService}
}

func (rs *RoomService) Create(createdByUserId string, isPrivateRoom bool) (*domain.Room, error) {
	if room, err := rs.service.Create(createdByUserId, isPrivateRoom); err != nil {
		return nil, err
	} else {
		if err = rs.repo.Create(room); err != nil {
			return nil, err
		} else {
			return room, nil
		}
	}

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

	user, err := rs.userService.GetUserById(userId)
	if err != nil {
		return err
	}

	rs.roomHub.Broadcast(
		roomId,
		websocket.EventParticipantJoined,
		&websocket.ParticipantJoinedPayload{
			UserId:   user.Id,
			UserName: user.Name})

	return nil
}

func (rs *RoomService) GetRoomById(roomId string) (*domain.Room, error) {
	return rs.repo.GetRoomById(roomId)
}
