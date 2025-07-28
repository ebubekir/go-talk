import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import { getAuth } from 'firebase/auth';

interface ApiConfig extends AxiosRequestConfig {
    requiresAuth?: boolean;
}

export class BaseApi {
    protected instance: AxiosInstance;
    protected baseURL: string;

    constructor() {
        this.baseURL = process.env.NEXT_PUBLIC_API_URL + '/v1';
        this.instance = axios.create({
            baseURL: this.baseURL,
            headers: {
                'Content-Type': 'application/json',
            },
        });

        this.setupInterceptors();
    }

    private setupInterceptors() {
        this.instance.interceptors.request.use(
            async (config: any) => {
                const auth = getAuth();
                const user = auth.currentUser;

                if (config.requiresAuth && user) {
                    const token = await user.getIdToken();
                    config.headers.Authorization = `Bearer ${token}`;
                }
                config.headers['Content-Type'] = 'application/json';

                return config;
            },
            (error) => {
                console.log('error', error);
                Promise.reject(error);
            }
        );

        this.instance.interceptors.response.use(
            (response) => response.data,
            (error) => Promise.reject<Error>(error.response?.data)
        );
    }

    protected async get<RequestType, ResponseType>(
        url: string,
        config?: ApiConfig
    ): Promise<ResponseType> {
        return await this.instance.get(url, config);
    }

    protected async post<RequestType, ResponseType>(
        url: string,
        data?: RequestType,
        config?: ApiConfig
    ): Promise<ResponseType> {
        return await this.instance.post(url, data, config);
    }

    protected async put<RequestType, ResponseType>(
        url: string,
        data?: RequestType,
        config?: ApiConfig
    ): Promise<ResponseType> {
        return await this.instance.put(url, data, config);
    }

    protected async delete<RequestType, ResponseType>(
        url: string,
        config?: ApiConfig
    ): Promise<ResponseType> {
        return await this.instance.delete(url, config);
    }
}
