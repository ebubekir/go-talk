import {BaseApi} from "@/api/base.api";
import {Room, RoomDetails} from "@/api/types";

export class RoomApi extends BaseApi {
    private readonly path ="/room";

    async create() {
        return this.post<null, Room>(
            this.path + "/create",
            null,
            {
                requiresAuth: true,
            }
        )
    }

    async getRoomDetails(id: string) {
        return this.get<string, RoomDetails>(
            this.path + `/${id}`,
            {
                requiresAuth: true,
            }
        )
    }
}