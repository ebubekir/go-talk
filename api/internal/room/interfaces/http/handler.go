package http

import (
	"github.com/ebubekir/go-talk/api/internal/room/application"
	userApp "github.com/ebubekir/go-talk/api/internal/user/application"
)

type RoomHandler struct {
	roomService *application.RoomService
	userService *userApp.UserService
}

func NewRoomHandler(roomService *application.RoomService, userService *userApp.UserService) *RoomHandler {
	return &RoomHandler{roomService: roomService, userService: userService}
}
