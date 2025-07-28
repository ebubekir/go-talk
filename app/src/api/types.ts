export interface UserResponse {
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