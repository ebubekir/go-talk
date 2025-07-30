import {RoomDetails} from "@/api/types";
import {createContext, useCallback, useContext, useEffect, useState} from "react";
import {api} from "@/api/api";
import {useParams} from "next/navigation";
import {LoadingSplash} from "@/components/loading-splash";
import {EventPayload} from "@/lib/room-events";

interface RoomContext {
    room: RoomDetails | null;
    isLoading: boolean;
    apiError?: string | null;
    getRoomDetails:  () => void;
    roomId?: string;
}

const RoomContext = createContext<RoomContext | null>(null);

export function RoomContextProvider({children}: { children: React.ReactNode }) {
    const [room, setRoom] = useState<RoomDetails | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const [initialLoading, setInitialLoading] = useState(true); // NEW
    const [apiError, setApiError] = useState<string | null>(null);
    const params = useParams()
    const roomId = params.id as string;

    const getRoomDetails = async  () => {
        try {
            // Only set isLoading if it's the initial load
            if (initialLoading) setIsLoading(true);
            const roomDetail = await api.rooms.getRoomDetails(roomId)
            setRoom(roomDetail)
        } catch (e: any) {
            setApiError(e?.message || "An error occurred while fetching room details.")
        } finally {
            if (initialLoading) setIsLoading(false);
            setInitialLoading(false); // NEW: after first load, never show splash again
        }
    }

    useEffect(() => {
        setInitialLoading(true); // NEW: reset for new roomId
        getRoomDetails()
    }, [roomId])

    if (isLoading && initialLoading) { // Only show splash on first load
        return (
            <div className="flex min-h-screen flex-col items-center justify-center ">
                <LoadingSplash />
            </div>
        )
    }


    if (apiError) {
        return (
            <div className="flex min-h-screen flex-col items-center justify-center ">
                <p className="text-red-500">{apiError}</p>
            </div>
        )
    }

    return (
        <RoomContext.Provider value={{room: room, roomId: roomId, isLoading: isLoading, apiError: apiError, getRoomDetails}}>
            {children}
        </RoomContext.Provider>
    );
}

export function useRoom() {
    const context = useContext(RoomContext);
    if (!context) {
        throw new Error("useRoom must be used within a RoomContextProvider");
    }
    return context
}