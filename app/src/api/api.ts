import { UserApi } from '@/api/user.api';
import {RoomApi} from "@/api/room.api";
import {ChatApi} from "@/api/chat.api";

export class Api {
    public readonly users: UserApi;
    public readonly rooms: RoomApi;
    public readonly chat: ChatApi

    constructor() {
        this.users = new UserApi();
        this.rooms = new RoomApi()
        this.chat = new ChatApi()
    }
}

// Singleton instance
export const api = new Api();
