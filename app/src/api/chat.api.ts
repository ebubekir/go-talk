import { BaseApi} from "@/api/base.api";
import { ChatMessage, Chat } from "@/api/types";

export class ChatApi extends  BaseApi {
    private readonly path = '/room';

    async getChatHistory(roomId: string) {
        return this.get<string, Chat>(
            `${this.path}/${roomId}/chat`,
            {
                requiresAuth: true,
            }
        );
    }

    async sendMessage(roomId: string, text: string) {
        return this.post<{ text: string }, any>(
            `${this.path}/${roomId}/chat`,
            { text },
            {
                requiresAuth: true,
            }
        );
    }

}