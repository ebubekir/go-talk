package http

import (
	"github.com/ebubekir/go-talk/api/internal/httpctx"
	"github.com/ebubekir/go-talk/api/internal/response"
	"github.com/ebubekir/go-talk/api/internal/room/domain"
	"github.com/gin-gonic/gin"
)

type SendChatMessageRequest struct {
	Text string `json:"text" binding:"required,min=1,max=1500" example:"hello world"`
} // @name SendChatMessage

type SendChatMessageResponse struct {
	IsOk bool `json:"isOk" binding:"required"`
} // @name SendChatMessageResponse

// SendChatMessage
// @Summary      SendChatMessage
// @Description  Send message to a chat.
// @Security 	 ApiKeyAuth
// @Param		 roomId path string true "Room Id"
// @Param		 request body SendChatMessageRequest true "Send chat message request"
// @Tags         Room,Chat
// @Success      200     {object} SendChatMessageResponse
// @Failure      default {object} response.ApiError
// @Router       /room/{roomId}/chat [post]
func (ch *ChatHandler) SendChatMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := httpctx.GetUserFromContext(c)
		roomId := c.Param("roomId")

		var req SendChatMessageRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			response.BadRequest(c, err)
			return
		}

		if room, err := ch.roomService.GetRoomById(roomId); err != nil {
			response.BadRequest(c, err)
			return
		} else if room == nil {
			response.BadRequestWithMessage(c, "room not found")
			return
		} else if !room.HasParticipant(domain.UserId(currentUser.Id)) {
			response.BadRequestWithMessage(c, "you are not a participant of this room")
			return
		}

		_, err := ch.chatService.GetChat(roomId)
		if err != nil {
			response.BadRequest(c, err)
			return
		}

		err = ch.chatService.SendMessage(roomId, currentUser.Id, req.Text)
		if err != nil {
			response.BadRequest(c, err)
			return
		}

		response.Success(c, &SendChatMessageResponse{IsOk: true})
	}
}
