import { UserApi } from '@/api/user.api';

export class Api {
    public readonly users: UserApi;

    constructor() {
        this.users = new UserApi();
    }
}

// Singleton instance
export const api = new Api();
