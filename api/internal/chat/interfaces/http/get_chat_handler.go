package http

import (
	"github.com/ebubekir/go-talk/api/internal/httpctx"
	"github.com/ebubekir/go-talk/api/internal/response"
	userHttp "github.com/ebubekir/go-talk/api/internal/user/interfaces/http"
	"github.com/gin-gonic/gin"
	"time"
)

type ChatMessage struct {
	User          userHttp.UserResponse `json:"user" binding:"required"`
	Text          string                `json:"text" binding:"required"`
	SentAt        time.Time             `json:"sentAt" binding:"required"`
	IsCurrentUser bool                  `json:"isCurrentUser" binding:"required"`
} // @name ChatMessage

type GetChatResponse struct {
	Id      string        `json:"id" binding:"required"`
	RoomId  string        `json:"roomId" binding:"required"`
	History []ChatMessage `json:"history" binding:"required"`
} // @name Chat

// GetChat
// @Summary      GetChat
// @Description  Get chat of room
// @Security 	 ApiKeyAuth
// @Param		 roomId path string true "Room Id"
// @Tags         Room,Chat
// @Success      200     {object} Chat
// @Failure      default {object} response.ApiError
// @Router       /room/{roomId}/chat [get]
func (ch *ChatHandler) GetChat() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := httpctx.GetUserFromContext(c)
		roomId := c.Param("roomId")

		if room, err := ch.roomService.GetRoomById(roomId); err != nil {
			response.BadRequest(c, err)
			return
		} else if room == nil {
			response.BadRequestWithMessage(c, "room not found")
			return
		}

		chat, err := ch.chatService.GetChat(roomId)
		if err != nil {
			response.BadRequest(c, err)
			return
		}

		userDetails := make(map[string]userHttp.UserResponse, 0)
		chatHistory := make([]ChatMessage, 0)

		if len(chat.ChatMessages) > 0 {
			for _, message := range chat.ChatMessages {
				if _, isExists := userDetails[message.UserId]; !isExists {
					user, err := ch.userService.GetUserById(message.UserId)
					if err != nil {
						response.SystemError(c, err)
						return
					}

					userDetails[message.UserId] = userHttp.UserResponse{
						Id:    user.Id,
						Name:  user.Name,
						Email: user.Email,
					}
				}
				chatHistory = append(chatHistory, ChatMessage{
					User:          userDetails[message.UserId],
					Text:          message.Text,
					SentAt:        message.SentAt,
					IsCurrentUser: message.UserId == currentUser.Id,
				})
			}
		}

		chatResponse := GetChatResponse{
			Id:      chat.Id,
			RoomId:  chat.RoomId,
			History: chatHistory,
		}

		response.Success(c, chatResponse)
	}
}
