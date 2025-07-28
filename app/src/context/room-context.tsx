import {RoomDetails} from "@/api/types";
import {createContext, useContext, useEffect, useState} from "react";
import {api} from "@/api/api";
import {useParams} from "next/navigation";
import {LoadingSplash} from "@/components/loading-splash";

interface RoomContext {
    room: RoomDetails | null;
    isLoading: boolean;
    apiError?: string | null;
}

const RoomContext = createContext<RoomContext | null>(null);

export function RoomContextProvider({children}: { children: React.ReactNode }) {
    const [room, setRoom] = useState<RoomDetails | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const [apiError, setApiError] = useState<string | null>(null);
    const params = useParams()
    const roomId = params.id as string;

    const getRoomDetails = async  () => {
        try {
            setIsLoading(true)
            const roomDetail = await api.rooms.getRoomDetails(roomId)
            setRoom(roomDetail)
        } catch (e: any) {
            setApiError(e?.message || "An error occurred while fetching room details.")
        } finally {
            setIsLoading(false)
        }
    }

    useEffect(() => {
        getRoomDetails()
    }, [])

    if (isLoading) {
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
        <RoomContext.Provider value={{room: room, isLoading: isLoading, apiError: apiError}}>
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