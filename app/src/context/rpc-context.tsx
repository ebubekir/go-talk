import { createContext, useContext, useEffect, useRef, useState } from 'react'
import { useSocket } from '@/context/room-socket-context'
import { showErrorToastMessage } from '@/components/ui/toast'

type RpcContextType = {
  localStream?: MediaStream | null
  remoteStreams?: Record<string, MediaStream | null>
  toggleCam: () => void
  toggleMic: () => void
  camEnabled: boolean
  micEnabled: boolean
}

const RPCContext = createContext<RpcContextType | undefined>(undefined)

export const RpcProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [localStream, setLocalStream] = useState<MediaStream | null>(null)
  const [camEnabled, setCamEnabled] = useState(false)
  const [micEnabled, setMicEnabled] = useState(false)
  const { socket } = useSocket()

  useEffect(() => {
    if (!socket) return
    if (socket.readyState !== socket.OPEN) return

    async function setupMedia() {
      try {
        const localMediaStream = await navigator.mediaDevices.getUserMedia({
          video: true,
          audio: true,
        })
        setLocalStream(localMediaStream)
        localMediaStream.getTracks().forEach((track) => (track.enabled = false))
      } catch (e) {
        showErrorToastMessage(
          'Failed to access media devices. Please check your permissions.',
        )
      }
    }

    setupMedia().then(() => {
      socket.addEventListener('message', () => {
        console.log('rpcProvider received message')
      })
    })
  }, [socket, socket?.readyState])

  const toggleCam = () => {
    if (!localStream) return
    const videoTrack = localStream.getVideoTracks()[0]
    videoTrack.enabled = !camEnabled
    setCamEnabled((prev) => !prev)
  }

  const toggleMic = () => {
    if (!localStream) return
    const audioTrack = localStream.getAudioTracks()[0]
    audioTrack.enabled = !micEnabled
    setMicEnabled((prev) => !prev)
  }

  return (
    <RPCContext.Provider
      value={{
        remoteStreams: {},
        camEnabled,
        micEnabled,
        toggleCam,
        toggleMic,
        localStream,
      }}
    >
      {children}
    </RPCContext.Provider>
  )
}

export const useRpc = (): RpcContextType => {
  const ctx = useContext(RPCContext)
  if (!ctx) {
    throw new Error('useRpc must be used within a RpcProvider')
  }
  return ctx
}
