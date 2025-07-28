package websocket

import (
	"fmt"
	"github.com/ebubekir/go-talk/api/pkg/firebase"
	"time"
)

type Hub struct {
	clients     map[string]map[*Client]bool // roomId ->clients
	broadcast   chan Event
	register    chan *Client
	unregister  chan *Client
	firebaseApp *firebase.App // For verify websocket connections
}

func NewHub(firebaseApp *firebase.App) *Hub {
	return &Hub{
		clients:     make(map[string]map[*Client]bool),
		broadcast:   make(chan Event),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		firebaseApp: firebaseApp,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			if h.clients[client.roomId] == nil {
				h.clients[client.roomId] = make(map[*Client]bool)
			}
			h.clients[client.roomId][client] = true

		case client := <-h.unregister:
			if roomClients, ok := h.clients[client.roomId]; ok {
				if _, exists := roomClients[client]; exists {
					delete(roomClients, client)
					close(client.send)
				}
			}

		case message := <-h.broadcast:
			if roomClients, ok := h.clients[message.RoomId]; ok {
				for client := range roomClients {
					select {
					case client.send <- message.ToJSON():
					default:
						close(client.send)
						delete(roomClients, client)
					}
				}
			}
		}
	}
}

func (h *Hub) Broadcast(roomId string, event EventType, payload interface{}) {
	fmt.Printf("Broadcasting event: %s\n", event)
	h.broadcast <- Event{
		RoomId:    roomId,
		Type:      event,
		Payload:   payload,
		Timestamp: time.Now(),
	}
}
