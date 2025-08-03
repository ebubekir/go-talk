export interface UserResponse {
    id: string;
    email: string;
    name: string;
}

export interface Room {
    id: string;
    isPrivate: boolean;
    ownerId: string;
}

export interface RoomDetails {
    id: string;
    owner: UserResponse;
    isPrivate: boolean;
    participants: UserResponse[];
}

export interface ChatMessage {
    user: UserResponse,
    text: string;
    sentAt: string;
    isCurrentUser: boolean;
}

export interface Chat {
    history: ChatMessage[];
    id: string;
    roomId: string;
}