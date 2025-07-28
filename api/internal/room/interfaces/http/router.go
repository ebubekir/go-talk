package http

import (
	"github.com/ebubekir/go-talk/api/internal/room/application"
	userApp "github.com/ebubekir/go-talk/api/internal/user/application"
	"github.com/gin-gonic/gin"
)

func RegisterRoomRoutes(router *gin.RouterGroup, roomService *application.RoomService, userService *userApp.UserService) {
	handler := NewRoomHandler(roomService, userService)

	roomGroup := router.Group("/room")
	{
		roomGroup.POST("/create", handler.CreateRoom())
		roomGroup.GET("/:roomId", handler.GetRoomDetail())
	}
}
