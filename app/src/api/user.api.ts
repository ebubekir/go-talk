import { BaseApi } from '@/api/base.abi';
import { UserResponse } from '@/api/types';

export class UserApi extends BaseApi {
    private readonly path = '/user';

    async create() {
        return this.post<null, UserResponse>(
            this.path + '/register',
            null,
            {
                requiresAuth: true,
            }
        );
    }

    async getUser()  {
        return this.get<null, UserResponse>(
            this.path + "/me",
            {
                requiresAuth: true,
            }
        )
    }
}
