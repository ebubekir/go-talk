package http

import (
	"github.com/ebubekir/go-talk/api/internal/httpctx"
	"github.com/ebubekir/go-talk/api/internal/response"
	userHttp "github.com/ebubekir/go-talk/api/internal/user/interfaces/http"
	"github.com/gin-gonic/gin"
)

type RoomDetailResponse struct {
	Id           string                  `json:"id" binding:"required"`
	Owner        userHttp.UserResponse   `json:"owner" binding:"required"`
	IsPrivate    bool                    `json:"isPrivate" binding:"required"`
	Participants []userHttp.UserResponse `json:"participants" binding:"required"`
} // @name RoomDetail

// GetRoomDetail
// @Summary      GetRoomDetail
// @Description  Get detail of room
// @Security 	 ApiKeyAuth
// @Param		 roomId path string true "Room Id"
// @Tags         Room
// @Success      200     {object} RoomDetail
// @Failure      default {object} response.ApiError
// @Router       /room/{roomId} [get]
func (r *RoomHandler) GetRoomDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := httpctx.GetUserFromContext(c)
		roomId := c.Param("roomId")

		room, err := r.roomService.GetRoomById(roomId)
		if err != nil {
			response.BadRequestWithMessage(c, err.Error())
			return
		}

		if room == nil {
			response.BadRequestWithMessage(c, "room not found")
			return
		}

		if err = r.roomService.JoinRoom(currentUser.Id, roomId); err != nil {
			response.SystemError(c, err)
			return
		}

		res := &RoomDetailResponse{
			Id:        string(room.Id),
			IsPrivate: room.IsPrivate,
		}

		userIds := make([]string, 0)
		userIds = append(userIds, currentUser.Id)

		ownerUser, err := r.userService.GetUserById(string(room.OwnerId))
		if err != nil {
			response.SystemError(c, err)
			return
		}
		res.Owner = userHttp.UserResponse{
			Id:    ownerUser.Id,
			Name:  ownerUser.Name,
			Email: ownerUser.Email,
		}

		for _, participant := range room.Participants {
			participantUser, err := r.userService.GetUserById(string(participant.UserId))
			if err != nil {
				response.SystemError(c, err)
				return
			}
			res.Participants = append(res.Participants, userHttp.UserResponse{
				Id:    participantUser.Id,
				Name:  participantUser.Name,
				Email: participantUser.Email,
			})
		}

		response.Success(c, res)

	}
}
