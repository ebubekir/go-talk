package application

import (
	"fmt"
	"github.com/ebubekir/go-talk/api/internal/websocket"
)

type RoomEventListener struct {
	roomService *RoomService
	broadCaster websocket.Broadcaster
}

func NewRoomEventListener(roomService *RoomService, broadcaster websocket.Broadcaster) *RoomEventListener {
	return &RoomEventListener{roomService: roomService, broadCaster: broadcaster}
}

func (e *RoomEventListener) HandleEvents() []websocket.EventType {
	return []websocket.EventType{
		websocket.EventParticipantJoined,
		websocket.EventParticipantLeft,
	}
}

func (e *RoomEventListener) HandleEvent(evt *websocket.Event) {
	if evt.Type == websocket.EventParticipantJoined {
		payload := evt.Payload.(*websocket.ParticipantJoinedPayload)
		err := e.roomService.JoinRoom(evt.RoomId, payload.UserEmail)
		if err != nil {
			fmt.Println(err)
		}
		e.broadCaster.Broadcast(evt)
	} else if evt.Type == websocket.EventParticipantLeft {
		payload := evt.Payload.(*websocket.ParticipantLeftPayload)
		err := e.roomService.LeaveRoom(evt.RoomId, payload.UserEmail)
		if err != nil {
			fmt.Println(err)
		}
		e.broadCaster.Broadcast(evt)
	} else {
		// Wrong event.
	}
}
