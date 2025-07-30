package websocket

import (
	"errors"
	"fmt"
	"github.com/ebubekir/go-talk/api/internal/response"
	userApp "github.com/ebubekir/go-talk/api/internal/user/application"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RoomWS(hub *Hub, userService *userApp.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("Socket connected for room: %s \n", c.Param("roomId"))
		roomId := c.Param("roomId")
		tokenString := c.Query("token")
		if tokenString == "" {
			response.UnauthorizedError(c, errors.New("token is required"))
			return
		}

		token, err := hub.firebaseApp.AuthClient.VerifyIDToken(c, tokenString)
		if err != nil {
			response.UnauthorizedError(c, err)
			return
		}

		userEmail := ""

		if v, hasValue := token.Claims["email"]; hasValue {
			userEmail = v.(string)
		}

		user, err := userService.GetUserByEmail(userEmail)
		if err != nil {
			response.SystemError(c, err)
			return
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		client := &Client{
			conn:      conn,
			send:      make(chan []byte, 256),
			roomId:    roomId,
			userEmail: user.Email,
			userName:  user.Name,
		}

		hub.register <- client

		go client.writePump()
		go client.readPump(hub)
	}
}
