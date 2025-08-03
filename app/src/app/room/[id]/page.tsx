'use client'
import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Separator } from '@/components/ui/separator'
import { useRoom } from '@/context/room-context'
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from '@/components/ui/resizable'
import RoomChat from '@/components/chat/room-chat'
import { UserAvatar } from '@/components/user-avatar'
import { Participants } from '@/components/participants'
import { ToggleControls } from '@/components/toggle-controls'

export default function Room({ params }: { params: Promise<{ id: string }> }) {
  const { room } = useRoom()

  return (
    <div className="flex  flex-col items-start justify-start p-8 h-screen">
      <Card className="w-full  h-full">
        <CardHeader>
          <CardTitle>room: {room?.id}</CardTitle>
          <CardDescription>Created by {room?.owner.name}</CardDescription>
          <CardAction>
            <div className="flex items-center gap-2">
              {room?.participants?.map((participant, key) => (
                <UserAvatar key={key} name={participant.name} email={participant.email} />
              ))}
            </div>
            <div className="float-right">
              <p>{room?.participants?.length} participants.</p>
            </div>
          </CardAction>
        </CardHeader>
        <Separator orientation="horizontal" />
        <CardContent className="h-full">
          <ResizablePanelGroup direction="horizontal">
            <ResizablePanel defaultSize={85} className="p-4">
              <Participants />
            </ResizablePanel>
            <ResizableHandle withHandle />
            <ResizablePanel defaultSize={15}>
              <RoomChat />
            </ResizablePanel>
          </ResizablePanelGroup>
        </CardContent>
        <CardFooter>
          <ToggleControls />
        </CardFooter>
      </Card>
    </div>
  )
}
