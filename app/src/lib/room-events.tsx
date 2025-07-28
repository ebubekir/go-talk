import {showInfoToastMessage} from "@/components/ui/toast";
import {useEffect} from "react";

enum RoomEventName {
    ParticipantJoined = 'participant-joined',
}

export interface EventPayload {
    payload: any;
    roomId: string;
    timestamp: string;
    type: RoomEventName;
}

abstract class RoomEvent {
    name: RoomEventName;

    constructor(name: RoomEventName) {
        this.name = name;
    }

    /**
     * Handle the event data.
     * @param data - The data associated with the event.
     */
    abstract handle(data: EventPayload): void;
}

class ParticipantJoinedEvent extends RoomEvent {
    constructor() {
        super(RoomEventName.ParticipantJoined);
    }

    handle(data: EventPayload) {
        const payload = data.payload as { userName: string, userId: string, joinedAt: string };
        showInfoToastMessage(`${payload?.userName} joined the room!`);
    }
}

const participantJoinedEvent = new ParticipantJoinedEvent();

export const webSocketEventHandlers: Record<RoomEventName, RoomEvent> = {
    [RoomEventName.ParticipantJoined]: participantJoinedEvent,
}

export const useRoomEvents = (id: string, authToken: string) => {
    useEffect(() => {

    }, [id, authToken])
}
