import { UserApi } from '@/api/user.api';
import {RoomApi} from "@/api/room.api";

export class Api {
    public readonly users: UserApi;
    public readonly rooms: RoomApi;

    constructor() {
        this.users = new UserApi();
        this.rooms = new RoomApi()
    }
}

// Singleton instance
export const api = new Api();
