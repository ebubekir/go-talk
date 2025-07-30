'use client'
import {createContext, useCallback, useContext, useEffect, useRef, useState} from "react";
import {connectToSocket} from "@/api/socket.api";
import {EventPayload, webSocketEventHandlers} from "@/lib/room-events";
import {useAuth} from "@/context/auth-context";
import {RoomContextProvider, useRoom} from "@/context/room-context";

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

export const RoomSocketProvider = ({ children }: { children: React.ReactNode }) => {
    const { authToken } = useAuth();
    const socketRef = useRef<WebSocket | null>(null);
    const [socket, setSocket] = useState<WebSocket | null>(null);
    const connectedRef = useRef(false)
    const { getRoomDetails, room, roomId } = useRoom()



    useEffect(() => {
        console.log('roomSocket useEffect')
        if (!room?.id || !authToken || connectedRef.current) return;

        const initSocket = async () => {
            try {
                const ws = await connectToSocket(room?.id, authToken);
                ws.onmessage = (event) => {
                    try {
                        const data = JSON.parse(event.data) as EventPayload;
                        if (webSocketEventHandlers[data.type]) {
                            webSocketEventHandlers[data.type]?.handle(data);
                            console.log(data)
                            // actionRef.current(data)
                            getRoomDetails()
                        }
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
    }, [room?.id, authToken]);

    return (
        <RoomSocketContext.Provider value={{ socket }}>
            {children}
        </RoomSocketContext.Provider>
    );
};