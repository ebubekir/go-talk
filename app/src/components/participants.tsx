import { CurrentUserCamera } from '@/components/current-user-camera'

export function Participants() {
  return (
    <div className="w-full grid grid-cols-4 gap-4 ">
      <CurrentUserCamera />
    </div>
  )
}
