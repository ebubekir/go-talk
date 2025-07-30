package websocket

type Broadcaster interface {
	Broadcast(evt *Event)
}
