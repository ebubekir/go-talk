package http

import (
	"github.com/ebubekir/go-talk/api/internal/chat/application"
	roomApp "github.com/ebubekir/go-talk/api/internal/room/application"
	userApp "github.com/ebubekir/go-talk/api/internal/user/application"
)

type ChatHandler struct {
	chatService *application.ChatService
	roomService *roomApp.RoomService
	userService userApp.UserService
}

func NewChatHandler(chatService *application.ChatService, roomService *roomApp.RoomService, userService *userApp.UserService) *ChatHandler {
	return &ChatHandler{chatService: chatService, roomService: roomService, userService: *userService}
}
