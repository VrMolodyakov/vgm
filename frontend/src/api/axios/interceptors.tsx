import { AxiosError, InternalAxiosRequestConfig } from "axios";

export const onErrorRequest = (error: AxiosError | Error): Promise<AxiosError> => {
    return Promise.reject(error);
}
  
export function createOnRequestInterceptor(getAccessToken:() => string){
    let i = async (config: InternalAxiosRequestConfig): Promise<InternalAxiosRequestConfig> => {
        console.log("inside req interce")
        const token = getAccessToken()
        if (token) {
          config.headers!['Authorization'] = token;
        }
        return config;
    }
    return i
}