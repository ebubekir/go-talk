package http

import (
	"github.com/ebubekir/go-talk/api/internal/httpctx"
	"github.com/ebubekir/go-talk/api/internal/response"
	"github.com/gin-gonic/gin"
)

type RoomResponse struct {
	Id        string `json:"id" binding:"required"`
	OwnerId   string `json:"ownerId" binding:"required"`
	IsPrivate bool   `json:"isPrivate" binding:"required"`
} // @name Room

// CreateRoom
// @Summary      CreateRoom
// @Description  Create room
// @Security 	 ApiKeyAuth
// @Tags         Room
// @Success      200     {object} Room
// @Failure      default {object} response.ApiError
// @Router       /room/create [post]
func (r *RoomHandler) CreateRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := httpctx.GetUserFromContext(c)

		room, err := r.roomService.Create(currentUser.Id, false)
		if err != nil {
			response.BadRequestWithMessage(c, err.Error())
			return
		}

		response.Success(c, &RoomResponse{
			Id:        string(room.Id),
			IsPrivate: room.IsPrivate,
			OwnerId:   string(room.OwnerId),
		})
	}
}
