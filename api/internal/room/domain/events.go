package domain

import "time"

type RoomEventType string

const (
	EventRoomCreated       RoomEventType = "room-created"
	EventParticipantJoined RoomEventType = "participant-joined"
	EventParticipantLeft   RoomEventType = "participant-left"
	EventMessageSent       RoomEventType = "message-sent"
	EventRoomClosed        RoomEventType = "room-closed"
)

type RoomEvent struct {
	Type      RoomEventType
	RoomId    RoomId
	Timestamp time.Time
	Payload   any
}
