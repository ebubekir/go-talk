import { useRpc } from '@/context/rpc-context'
import { CameraIcon, CameraOffIcon, MicIcon, MicOffIcon } from 'lucide-react'
import { Button } from '@/components/ui/button'

export function ToggleControls() {
  return (
    <div className="flex items-center gap-2 m-auto">
      <CameraToggleButton />
      <MicToggleButton />
    </div>
  )
}

function CameraToggleButton() {
  const { toggleCam, camEnabled } = useRpc()

  return (
    <Button
      variant={camEnabled ? 'default' : 'destructive'}
      onClick={toggleCam}
      className="cursor-pointer"
    >
      {camEnabled ? (
        <CameraIcon className="w-10 h-10" />
      ) : (
        <CameraOffIcon className="w-10 h-10" />
      )}
    </Button>
  )
}

function MicToggleButton() {
  const { toggleMic, micEnabled } = useRpc()

  return (
    <Button
      variant={micEnabled ? 'default' : 'destructive'}
      onClick={toggleMic}
      className="cursor-pointer"
    >
      {micEnabled ? (
        <MicIcon className={`w-10 h-10`} />
      ) : (
        <MicOffIcon className="w-10 h-10" />
      )}
    </Button>
  )
}
