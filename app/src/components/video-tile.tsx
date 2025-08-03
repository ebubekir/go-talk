import { useEffect, useRef, useState } from 'react'
import { UserAvatar } from '@/components/user-avatar'
import { MessageCircleIcon, MicVocalIcon } from 'lucide-react'

export default function VideoTile({
  stream,
  isLocal,
  camEnabled = true,
  userName,
  userEmail,
}: {
  stream?: MediaStream | null
  isLocal?: boolean
  camEnabled?: boolean
  userName?: string
  userEmail?: string
}) {
  const videoRef = useRef<HTMLVideoElement>(null)
  const [isTalking, setIsTalking] = useState(false)

  useEffect(() => {
    if (videoRef.current && stream) {
      videoRef.current.srcObject = stream
    }

    if (stream) {
      const audioContext = new AudioContext()
      const source = audioContext.createMediaStreamSource(stream)
      const analyser = audioContext.createAnalyser()
      const dataArray = new Uint8Array(analyser.frequencyBinCount)

      source.connect(analyser)

      function detectVolume() {
        analyser.getByteFrequencyData(dataArray)
        const volume = dataArray.reduce((a, b) => a + b) / dataArray.length

        if (volume > 10) {
          setIsTalking(true)
        } else {
          setIsTalking(false)
        }

        requestAnimationFrame(detectVolume)
      }

      detectVolume()
    }
  }, [stream, camEnabled])

  if (!camEnabled || !stream) {
    // If the camera is not enabled or the stream is not available, show a placeholder
    return (
      <div className="w-full h-full relative aspect-video rounded-lg shadow-lg bg-gray-400 flex items-center justify-center overflow-hidden">
        <video
          ref={videoRef}
          autoPlay
          playsInline
          muted={isLocal}
          style={{ display: 'none' }}
          className="rounded-lg shadow-lg"
        />
        <div className="m-auto">
          <UserAvatar name={userName} email={userEmail} />
        </div>
        {isTalking && (
          <div className="bg-blue-300 w-full h-1 absolute bottom-0 pulse"></div>
        )}
      </div>
    )
  }

  return (
    <div className="w-full h-full relative aspect-video flex items-center justify-center overflow-hidden rounded-lg ">
      <video
        ref={videoRef}
        autoPlay
        playsInline
        muted={isLocal}
        className="rounded-lg shadow-lg w-full "
      />
      {isTalking && (
        <div className="bg-blue-300 rounded-full absolute top-0 left-0 w-10 h-10 flex items-center justify-center animate-pulse text-white mt-2 ml-2">
          <MicVocalIcon className="w-5 h-5 font-extrabold" />
        </div>
      )}
    </div>
  )
}
