package http

import (
	"github.com/ebubekir/go-talk/api/internal/room/application"
	"github.com/gin-gonic/gin"
)

func RegisterRoomRoutes(router *gin.RouterGroup, roomService *application.RoomService) {
	handler := NewRoomHandler(roomService)

	roomGroup := router.Group("/room")
	{
		roomGroup.POST("/create", handler.CreateRoom())
	}
}
