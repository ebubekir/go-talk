import {showInfoToastMessage} from "@/components/ui/toast";
import {useEffect} from "react";

enum RoomEventName {
    ParticipantJoined = 'participant-joined',
    ParticipantLeft = 'participant-left',
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
        const payload = data.payload as { userName: string, userEmail: string, joinedAt: string };
        showInfoToastMessage(`${payload?.userName} joined the room!`);
    }
}

class ParticipantLeftEvent extends RoomEvent {
    constructor() {
        super(RoomEventName.ParticipantJoined);
    }

    handle(data: EventPayload) {
        const payload = data.payload as { userName: string, userEmail: string, leftAt: string };
        showInfoToastMessage(`${payload?.userName} left the room!`);
    }
}

export const webSocketEventHandlers: Record<RoomEventName, RoomEvent> = {
    [RoomEventName.ParticipantJoined]: new ParticipantJoinedEvent(),
    [RoomEventName.ParticipantLeft]: new ParticipantLeftEvent(),
}
