// context/socket-context.tsx
'use client'
import {createContext, useContext, useEffect, useRef, useState} from "react";
import {connectToSocket} from "@/api/socket.api";
import {EventPayload, webSocketEventHandlers} from "@/lib/room-events";
import {useAuth} from "@/context/auth-context";
import {RoomContextProvider} from "@/context/room-context";

interface RoomSocketContextType {
    socket: WebSocket | null;
}

const RoomSocketContext = createContext<RoomSocketContextType>({ socket: null });

export const useSocket = () => {
    const context = useContext(RoomSocketContext);
    if (!context) {
        throw new Error("useSocket must be used within a SocketProvider");
    }
    return context;
}

export const RoomSocketProvider = ({ roomId, children }: { roomId: string, children: React.ReactNode }) => {
    const { authToken } = useAuth();
    const socketRef = useRef<WebSocket | null>(null);
    const [socket, setSocket] = useState<WebSocket | null>(null);
    const connectedRef = useRef(false)

    useEffect(() => {
        if (!roomId || !authToken || connectedRef.current) return;

        const initSocket = async () => {
            try {
                const ws = await connectToSocket(roomId, authToken);
                ws.onmessage = (event) => {
                    try {
                        const data = JSON.parse(event.data) as EventPayload;
                        webSocketEventHandlers[data.type]?.handle(data);
                    } catch (e) {
                        console.error("Socket message error", e);
                    }
                };
                connectedRef.current = true;
                socketRef.current = ws;
                setSocket(ws);
            } catch (e) {
                console.error("Failed to connect to socket", e);
            }
        };

        initSocket();

        return () => {
            socketRef.current?.close();
        };
    }, [roomId, authToken]);

    return (
        <RoomSocketContext.Provider value={{ socket }}>
            {children}
        </RoomSocketContext.Provider>
    );
};