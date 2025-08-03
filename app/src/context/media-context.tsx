import React, { createContext, useContext, useEffect, useRef, useState } from 'react'
import { useSocket } from './room-socket-context'
import { useRoom } from './room-context'
import { useAuth } from '@/context/auth-context'
import { UserResponse } from '@/api/types'

const GOOGLE_STUN = [{ urls: 'stun:stun.l.google.com:19302' }]

type RemoteStreams = { [peerId: string]: MediaStream }
type PeerConnections = { [peerId: string]: RTCPeerConnection }

type MediaContextType = {
  stream: MediaStream | null
  remoteStreams: RemoteStreams
  startMedia: () => Promise<void>
  stopMedia: () => void
  isEnabled: boolean
}

const MediaContext = createContext<MediaContextType | undefined>(undefined)

export const MediaProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [stream, setStream] = useState<MediaStream | null>(null)
  const [remoteStreams, setRemoteStreams] = useState<RemoteStreams>({})
  const [isEnabled, setIsEnabled] = useState(false)
  const peerConnections = useRef<PeerConnections>({})
  const { socket } = useSocket()
  const { room } = useRoom()
  const { user } = useAuth()
  const userId = user?.id

  // Helper: check if socket is open
  const isSocketOpen = () => socket && socket.readyState === WebSocket.OPEN

  // Start/stop local media
  const startMedia = async () => {
    if (stream) return
    try {
      const userStream = await navigator.mediaDevices.getUserMedia({
        video: true,
        audio: true,
      })
      setStream(userStream)
      setIsEnabled(true)
    } catch (err) {
      alert('Could not access camera/mic: ' + (err as Error).message)
    }
  }

  const stopMedia = () => {
    stream?.getTracks().forEach((track) => track.stop())
    setStream(null)
    setIsEnabled(false)
    Object.values(peerConnections.current).forEach((pc) => pc.close())
    peerConnections.current = {}
    setRemoteStreams({})
  }

  // Helper: create peer connection
  const createPeerConnection = (peerId: string) => {
    if (!isSocketOpen()) {
      console.warn('Socket is not open, cannot create peer connection')
      return null
    }
    const pc = new RTCPeerConnection({ iceServers: GOOGLE_STUN })
    if (stream) {
      stream.getTracks().forEach((track) => pc.addTrack(track, stream))
    }
    pc.ontrack = (event) => {
      setRemoteStreams((prev) => {
        // Always use a single MediaStream per peer
        let remoteStream = prev[peerId]
        if (!remoteStream) {
          remoteStream = new MediaStream()
        }
        // Add all tracks from the event, if not already present
        event.streams[0]?.getTracks().forEach((track) => {
          if (!remoteStream.getTrackById(track.id)) {
            remoteStream.addTrack(track)
          }
        })
        // Also add the event.track directly (for browsers that fire ontrack per track)
        if (!remoteStream.getTrackById(event.track.id)) {
          remoteStream.addTrack(event.track)
        }
        return { ...prev, [peerId]: remoteStream }
      })
    }
    pc.onicecandidate = (event) => {
      if (event.candidate && isSocketOpen()) {
        socket!.send(
          JSON.stringify({
            type: 'ice-candidate',
            candidate: event.candidate,
            to: peerId,
          }),
        )
      }
    }
    peerConnections.current[peerId] = pc
    return pc
  }

  // Automate connection to all peers
  useEffect(() => {
    if (!isSocketOpen() || !room?.participants || !userId || !isEnabled) return
    console.log('peerConnections', peerConnections.current)

    // Connect to all other users
    room.participants.forEach(async (participant: UserResponse) => {
      if (participant.id === userId) return // Don't connect to self
      if (peerConnections.current[participant.id]) return // Already connected
      console.log('participant', participant.id, userId)

      const pc = createPeerConnection(participant.id)
      if (!pc) return // Only proceed if pc is not null

      try {
        const offer = await pc.createOffer()
        await pc.setLocalDescription(offer)
        if (isSocketOpen()) {
          socket!.send(
            JSON.stringify({ type: 'offer', offer, to: participant.id, from: userId }),
          )
        }
      } catch (err) {
        console.error('Error creating/sending offer:', err)
      }
    })

    // Clean up on unmount
    return () => {
      Object.values(peerConnections.current).forEach((pc) => pc.close())
      peerConnections.current = {}
      setRemoteStreams({})
    }
    // eslint-disable-next-line
  }, [socket, room?.participants, isEnabled, userId, stream])

  // Handle signaling messages
  useEffect(() => {
    if (!socket) return

    const handleMessage = async (event: MessageEvent) => {
      const data = JSON.parse(event.data)?.payload
      const peerId = data.from
      if (!peerId || peerId === userId) return

      if (data.type === 'offer') {
        const pc = createPeerConnection(peerId)
        if (!pc) return
        try {
          await pc.setRemoteDescription(new RTCSessionDescription(data.offer))
          const answer = await pc.createAnswer()
          await pc.setLocalDescription(answer)
          if (isSocketOpen()) {
            socket!.send(JSON.stringify({ type: 'answer', answer, to: peerId }))
          }
        } catch (err) {
          console.error('Error handling offer:', err)
        }
      }
      if (data.type === 'answer' && peerConnections.current[peerId]) {
        try {
          await peerConnections.current[peerId].setRemoteDescription(
            new RTCSessionDescription(data.answer),
          )
        } catch (err) {
          console.error('Error setting remote description for answer:', err)
        }
      }
      if (data.type === 'ice-candidate' && peerConnections.current[peerId]) {
        try {
          await peerConnections.current[peerId].addIceCandidate(
            new RTCIceCandidate(data.candidate),
          )
        } catch (e) {
          console.error('Error adding received ice candidate', e)
        }
      }
    }

    socket.addEventListener('message', handleMessage)
    // Handle socket close: clean up everything
    const handleSocketClose = () => {
      Object.values(peerConnections.current).forEach((pc) => pc.close())
      peerConnections.current = {}
      setRemoteStreams({})
      // stopMedia()
    }
    socket.addEventListener('close', handleSocketClose)

    return () => {
      socket.removeEventListener('message', handleMessage)
      // socket.removeEventListener('close', handleSocketClose)
    }
    // eslint-disable-next-line
  }, [socket, stream, userId])

  return (
    <MediaContext.Provider
      value={{ stream, remoteStreams, startMedia, stopMedia, isEnabled }}
    >
      {children}
    </MediaContext.Provider>
  )
}

export const useMedia = () => {
  const ctx = useContext(MediaContext)
  if (!ctx) throw new Error('useMedia must be used within a MediaProvider')
  return ctx
}
