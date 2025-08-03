'use client'
import { createContext, useContext, useEffect, useRef, useState } from 'react'
import { connectToSocket } from '@/api/socket.api'
import { EventPayload, RoomEventName, webSocketEventHandlers } from '@/lib/room-events'
import { useAuth } from '@/context/auth-context'
import { useRoom } from '@/context/room-context'
import { WebRTCManager } from '@/lib/webrtc'

interface RoomSocketContextType {
  socket: WebSocket | null
  camEnabled?: boolean
  micEnabled?: boolean
  toggleCam?: () => void
  toggleMic?: () => void
  leaveCall?: () => void
  rtcManager?: WebRTCManager | null
  localStream?: MediaStream | null
  // participantStreams?: ParticipantStream[]
}

type ParticipantStream = {
  id: string
  stream: MediaStream | null
}

const RoomSocketContext = createContext<RoomSocketContextType>({ socket: null })

export const useSocket = () => {
  const context = useContext(RoomSocketContext)
  if (!context) {
    throw new Error('useSocket must be used within a SocketProvider')
  }
  return context
}

export const RoomSocketProvider = ({ children }: { children: React.ReactNode }) => {
  const { authToken } = useAuth()
  const socketRef = useRef<WebSocket | null>(null)
  const [socket, setSocket] = useState<WebSocket | null>(null)
  const connectedRef = useRef(false)
  const { getRoomDetails, roomId, getChat, room } = useRoom()

  useEffect(() => {
    if (!roomId || !authToken || connectedRef.current) return
    const initSocket = async () => {
      try {
        const ws = await connectToSocket(roomId, authToken)
        ws.onmessage = (event) => {
          try {
            console.log('room-socket-context received message')
            const data = JSON.parse(event.data) as EventPayload
            if (webSocketEventHandlers[data.type]) {
              webSocketEventHandlers[data.type].handle(data)
              getRoomDetails()
              getChat()
            }
          } catch (e) {
            console.error('Socket message error', e)
          }
        }

        connectedRef.current = true
        socketRef.current = ws
        setSocket(ws)
      } catch (e) {
        console.error('Failed to connect to socket', e)
      }
    }

    initSocket()

    return () => {
      socketRef.current?.close()
    }
  }, [roomId, authToken])

  return (
    <RoomSocketContext.Provider value={{ socket }}>{children}</RoomSocketContext.Provider>
  )
}
