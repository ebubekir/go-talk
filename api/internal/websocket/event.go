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
)

type ParticipantJoinedPayload struct {
	UserId   string    `json:"userId"`
	UserName string    `json:"userName"`
	JoinedAt time.Time `json:"joinedAt"`
}

type Event struct {
	Type      EventType   `json:"type"`
	RoomId    string      `json:"roomId"`
	Payload   interface{} `json:"payload"`
	Timestamp time.Time   `json:"timestamp"`
}

func (e *Event) ToJSON() []byte {
	bytes, _ := json.Marshal(e)
	return bytes
}
