package websocket

import (
	"encoding/json"
	"time"
)

type EventType string

const (
	EventRoomCreated       EventType = "room-created"
	EventParticipantJoined EventType = "participant-joined"
	EventParticipantLeft   EventType = "participant-left"
	EventMessageSent       EventType = "message-sent"
	EventRoomClosed        EventType = "room-closed"
	EventTypeOffer         EventType = "offer"
	EventTypeAnswer        EventType = "answer"
	EventTypeIceCandidate  EventType = "ice-candidate"
	EventTypeCandidate     EventType = "candidate"
	EventTypeIce           EventType = "ice"
)

type ParticipantJoinedPayload struct {
	UserEmail string    `json:"userEmail"`
	UserName  string    `json:"userName"`
	JoinedAt  time.Time `json:"joinedAt"`
}

type ParticipantLeftPayload struct {
	UserEmail string    `json:"userEmail"`
	UserName  string    `json:"userName"`
	LeftAt    time.Time `json:"leftAt"`
}

type SendMessagePayload struct {
	UserEmail string    `json:"userEmail"`
	UserName  string    `json:"userName"`
	Text      string    `json:"text"`
	SentAt    time.Time `json:"sentAt"`
}
type Event struct {
	From      string      `json:"from"`
	To        string      `json:"to"`
	Type      EventType   `json:"type"`
	RoomId    string      `json:"roomId"`
	Payload   interface{} `json:"payload"`
	Timestamp time.Time   `json:"timestamp"`
}

func (e *Event) ToJSON() []byte {
	bytes, _ := json.Marshal(e)
	return bytes
}

type EventListener interface {
	HandleEvents() []EventType
	HandleEvent(evt *Event)
}

type EventDispatcher struct {
	listeners []EventListener
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		listeners: make([]EventListener, 0),
	}
}

func (e *EventDispatcher) Register(listener EventListener) {
	e.listeners = append(e.listeners, listener)
}

func (e *EventDispatcher) Dispatch(evt *Event) {
	for _, listener := range e.listeners {
		for _, eventType := range listener.HandleEvents() {
			if eventType == evt.Type {
				listener.HandleEvent(evt)
				break
			}
		}
	}
}
