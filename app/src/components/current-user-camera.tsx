import { useRpc } from '@/context/rpc-context'
import VideoTile from '@/components/video-tile'
import { useAuth } from '@/context/auth-context'

export function CurrentUserCamera() {
  const { localStream, camEnabled } = useRpc()
  const { user } = useAuth()
  return (
    <VideoTile
      stream={localStream}
      isLocal={true}
      camEnabled={camEnabled}
      userEmail={user?.email}
      userName={user?.name}
    />
  )
}
