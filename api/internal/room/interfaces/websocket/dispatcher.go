package websocket

//import (
//	"github.com/ebubekir/go-talk/api/internal/room/domain"
//	"golang.org/x/net/websocket"
//)
//
//type Publisher interface {
//	Publish(event domain.RoomEvent)
//}
//
//type GinWebSocketPublisher struct {
//	connections map[domain.RoomId][]*websocket.Conn
//}
//
//func (p *GinWebSocketPublisher) Publish(event domain.RoomEvent) {
//	for _, conn := range p.connections[event.RoomId] {
//		_ = conn.WriteJSON(event)
//	}
//}
