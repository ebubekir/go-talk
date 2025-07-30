'use client'
import {Card, CardAction, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card";
import {Separator} from "@/components/ui/separator";
import {useRoom} from "@/context/room-context";
import {Avatar, AvatarFallback} from "@/components/ui/avatar";
import {getAvatarName} from "@/lib/utils";
import {ResizableHandle, ResizablePanel, ResizablePanelGroup} from "@/components/ui/resizable";
import {Tooltip, TooltipContent, TooltipTrigger} from "@/components/ui/tooltip";


export default function Room({params,}: {
    params: Promise<{ id: string }>
}) {
    const { room } = useRoom()

    return (
        <div className="flex  flex-col items-start justify-start p-8 h-screen">
            <Card className="w-full p-4 h-full">
                <CardHeader>
                    <CardTitle>room: {room?.id}</CardTitle>
                    <CardDescription>Created by {room?.owner.name}</CardDescription>
                    <CardAction>
                        <div className="flex items-center gap-2">
                            <Avatar>
                                <Tooltip>
                                    <TooltipTrigger asChild>
                                        <Avatar>
                                            <AvatarFallback>{getAvatarName(room?.owner.name || "")}</AvatarFallback>
                                        </Avatar>
                                    </TooltipTrigger>
                                    <TooltipContent>
                                        <div className="flex flex-col items-start">
                                            <p className="font-semibold">{room?.owner.name}</p>
                                            <p className="text-sm ">{room?.owner.email}</p>
                                        </div>
                                    </TooltipContent>
                                </Tooltip>
                            </Avatar>
                            {
                                room?.participants?.map(participant => (
                                    <Tooltip key={participant.email}>
                                        <TooltipTrigger asChild>
                                            <Avatar key={participant.email}>
                                                <AvatarFallback>{getAvatarName(participant.name)}</AvatarFallback>
                                            </Avatar>
                                        </TooltipTrigger>
                                        <TooltipContent>
                                            <div className="flex flex-col items-start">
                                                <p className="font-semibold">{participant.name}</p>
                                                <p className="text-sm ">{participant.email}</p>
                                            </div>
                                        </TooltipContent>
                                    </Tooltip>

                                ))
                            }
                        </div>
                        <div className="float-right">
                            <p>{room?.participants?.length} participants.</p>
                        </div>
                    </CardAction>
                </CardHeader>
                <Separator orientation="horizontal" />
                <CardContent className="h-full">
                    <ResizablePanelGroup direction="horizontal" >
                        <ResizablePanel defaultSize={85} >Cameras and participants</ResizablePanel>
                        <ResizableHandle withHandle />
                        <ResizablePanel defaultSize={15}>Chat area</ResizablePanel>
                    </ResizablePanelGroup>

                </CardContent>
            </Card>
        </div>
    )
}
