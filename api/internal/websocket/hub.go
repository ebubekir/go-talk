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
	dispatcher  *EventDispatcher
}

func NewHub(firebaseApp *firebase.App, dispatcher *EventDispatcher) *Hub {
	return &Hub{
		clients:     make(map[string]map[*Client]bool),
		broadcast:   make(chan Event),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		firebaseApp: firebaseApp,
		dispatcher:  dispatcher,
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
			go h.dispatcher.Dispatch(&Event{
				RoomId: client.roomId,
				Type:   EventParticipantJoined,
				Payload: &ParticipantJoinedPayload{
					UserName:  client.userName,
					UserEmail: client.userEmail,
					JoinedAt:  time.Now(),
				},
				Timestamp: time.Now(),
			})

		case client := <-h.unregister:
			if roomClients, ok := h.clients[client.roomId]; ok {
				if _, exists := roomClients[client]; exists {
					delete(roomClients, client)
					go h.dispatcher.Dispatch(&Event{
						RoomId: client.roomId,
						Type:   EventParticipantLeft,
						Payload: &ParticipantLeftPayload{
							UserName:  client.userName,
							UserEmail: client.userEmail,
							LeftAt:    time.Now(),
						},
						Timestamp: time.Now(),
					})
					close(client.send)
				}
			}

		case message := <-h.broadcast:
			fmt.Printf("Broadcasting event: %s\n", message.Type)
			if roomClients, ok := h.clients[message.RoomId]; ok {
				for client := range roomClients {
					select {
					case client.send <- message.ToJSON():
					default:
						print("Close client'da misin la")
						close(client.send)
						delete(roomClients, client)
					}
				}
			}
		}
	}
}

func (h *Hub) Broadcast(evt *Event) {
	h.broadcast <- *evt
}
