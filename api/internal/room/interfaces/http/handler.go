package http

import "github.com/ebubekir/go-talk/api/internal/room/application"

type RoomHandler struct {
	roomService *application.RoomService
}

func NewRoomHandler(roomService *application.RoomService) *RoomHandler {
	return &RoomHandler{roomService: roomService}
}
