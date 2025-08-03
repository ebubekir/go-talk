import { Tooltip, TooltipContent, TooltipTrigger } from '@/components/ui/tooltip'
import { Avatar, AvatarFallback } from '@/components/ui/avatar'
import { getAvatarName } from '@/lib/utils'

export function UserAvatar({ name, email }: { name?: string; email?: string }) {
  return (
    <Avatar>
      <Tooltip>
        <TooltipTrigger asChild>
          <Avatar>
            <AvatarFallback>{getAvatarName(name || '')}</AvatarFallback>
          </Avatar>
        </TooltipTrigger>
        <TooltipContent>
          <div className="flex flex-col items-start">
            <p className="font-semibold">{name}</p>
            <p className="text-sm ">{email}</p>
          </div>
        </TooltipContent>
      </Tooltip>
    </Avatar>
  )
}
