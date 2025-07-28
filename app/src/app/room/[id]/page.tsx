'use client'
import {use, useEffect, useRef, useState} from 'react'
import {RoomDetails} from "@/api/types";
import {api} from "@/api/api";
import {LoadingSplash} from "@/components/loading-splash";
import {Card, CardDescription, CardHeader, CardTitle} from "@/components/ui/card";
import {Separator} from "@/components/ui/separator";
import {closeSocket, connectToSocket} from "@/api/socket.api";
import {useAuth} from "@/context/auth-context";
import {useRoom} from "@/context/room-context";
import {EventPayload, webSocketEventHandlers} from "@/lib/room-events";


export default function Room({params,}: {
    params: Promise<{ id: string }>
}) {
    const { id } = use(params)
    const { room } = useRoom()
    const { authToken } = useAuth()
    // const socketRef = useRef<WebSocket | null>(null);
    //
    //
    // useEffect(() => {
    //     if (!id || !authToken) return;
    //
    //     let ws: WebSocket;
    //
    //     const init = async () => {
    //         try {
    //             ws = await connectToSocket(id, authToken);
    //             socketRef.current = ws;
    //             ws.onmessage = (event) => {
    //                 try {
    //                     const data = JSON.parse(event.data) as EventPayload ;
    //                     webSocketEventHandlers[data.type].handle(data);
    //                 }  catch (error) {
    //                     console.error("Socket error", error)
    //                 }
    //             }
    //
    //         } catch (e) {
    //             console.error("Socket init error")
    //         }
    //     }
    //     init();
    //     return () => {
    //         if (ws && ws.readyState === WebSocket.OPEN) {
    //             ws.close();
    //         }
    //     };
    // }, [id, authToken]);


    // if (isLoading) {
    //     return (
    //         <div className="flex min-h-screen flex-col items-center justify-center ">
    //             <LoadingSplash />
    //         </div>
    //     )
    // }
    //
    //
    // if (apiError) {
    //     return (
    //         <div className="flex min-h-screen flex-col items-center justify-center ">
    //             <p className="text-red-500">{apiError}</p>
    //         </div>
    //     )
    // }

    return (
        <div className="flex min-h-screen flex-col items-center justify-center p-8 ">
            <Card className="w-full p-4">
                <CardHeader>
                    <CardTitle>room: {room?.id}</CardTitle>
                    <CardDescription>Created by {room?.owner.name}</CardDescription>
                </CardHeader>
                <Separator orientation="horizontal" />
            </Card>
        </div>
    )
}
