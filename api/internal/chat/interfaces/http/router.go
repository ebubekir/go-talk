package http

import (
	"github.com/ebubekir/go-talk/api/internal/chat/application"
	roomApp "github.com/ebubekir/go-talk/api/internal/room/application"
	userApp "github.com/ebubekir/go-talk/api/internal/user/application"
	"github.com/gin-gonic/gin"
)

func RegisterChatRoutes(router *gin.RouterGroup, chatService *application.ChatService, userService *userApp.UserService, roomService *roomApp.RoomService) {
	handler := NewChatHandler(chatService, roomService, userService)

	chatGroup := router.Group("/room/:roomId/chat")
	{
		chatGroup.GET("", handler.GetChat())
		chatGroup.POST("", handler.SendChatMessage())
	}
}
