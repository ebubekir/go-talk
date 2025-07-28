package websocket

import (
	"errors"
	"fmt"
	"github.com/ebubekir/go-talk/api/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RoomWS(hub *Hub) gin.HandlerFunc {
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

		userName := ""

		if v, hasValue := token.Claims["name"]; hasValue {
			userName = v.(string)
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		client := &Client{
			conn:      conn,
			send:      make(chan []byte, 256),
			roomId:    roomId,
			userEmail: userEmail,
			userName:  userName,
		}

		hub.register <- client

		go client.writePump()
		go client.readPump(hub)
	}
}
