'use client'

import {RoomContextProvider, useRoom} from "@/context/room-context";
import {use} from "react";
import {RoomSocketProvider} from "@/context/room-socket-context";

export default function RoomLayout({
    children,
    params,
}: Readonly<{
    children: React.ReactNode;
    params: Promise<{ id: string}>
}>) {
    const { id } = use(params)
    return (
        <RoomContextProvider>
            <RoomSocketProvider>
                {children}
            </RoomSocketProvider>
        </RoomContextProvider>
    );
}