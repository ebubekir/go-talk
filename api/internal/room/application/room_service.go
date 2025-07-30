package application

import (
	"github.com/ebubekir/go-talk/api/internal/room/domain"
	userApp "github.com/ebubekir/go-talk/api/internal/user/application"
)

type RoomService struct {
	repo        domain.RoomRepository
	service     *domain.RoomService
	userService userApp.UserService
}

func NewRoomService(repo domain.RoomRepository, userService *userApp.UserService) *RoomService {
	roomService := &domain.RoomService{}
	return &RoomService{repo: repo, service: roomService, userService: *userService}
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

func (rs *RoomService) JoinRoom(roomId string, userEmail string) error {
	room, err := rs.repo.GetRoomById(roomId)
	if err != nil {
		return err
	}

	user, err := rs.userService.GetUserByEmail(userEmail)
	if err != nil {
		return err
	}

	err = room.AddParticipant(domain.UserId(user.Id))
	if err != nil {
		return err
	}
	_ = rs.repo.Save(room)
	return nil
}

func (rs *RoomService) GetRoomById(roomId string) (*domain.Room, error) {
	return rs.repo.GetRoomById(roomId)
}

func (rs *RoomService) LeaveRoom(roomId string, userEmail string) error {
	room, err := rs.repo.GetRoomById(roomId)
	if err != nil {
		return nil
	}

	user, err := rs.userService.GetUserByEmail(userEmail)
	if err != nil {
		return err
	}

	err = room.RemoveParticipant(domain.UserId(user.Id))
	if err != nil {
		return err
	}

	err = rs.repo.Save(room)

	if err != nil {
		return err
	}

	return nil
}
