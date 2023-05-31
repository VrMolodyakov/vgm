import { AxiosError, AxiosInstance, InternalAxiosRequestConfig } from "axios";
import { userInstance } from "./axiosInstance";

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

userInstance.interceptors.request.use(
    config => {
      const token = localStorage.getItem('access-token');
      if (token) {
        config.headers!['Authorization'] = token;
      }
      return config;
    },
    error => {
      return Promise.reject(error);
    }
  );
// export function createOnResponseInterceptor(base:AxiosInstance,refresh:AxiosInstance){
//     let a = (err:AxiosError) => {
//         const originalConfig = err.config;
    
//         if (originalConfig?.url !== '/user/login' && err.response) {
//           // Access Token was expired
//           if (err.response.status === 401 && !originalConfig._retry) {
//             originalConfig._retry = true;
    
//             try {
//               const rs = await refresh.post(
//                 'https://api.example.org/user/refresh',
//                 {
//                   headers: {
//                     Authorization: localStorage.getItem('refresh-token')!
//                   }
//                 }
//               );
    
//               const access = rs.data.data['X-Auth-Token'];
//               const refresh = rs.data.data['X-Refresh-Token'];
    
//               localStorage.setItem('access-token', access);
//               localStorage.setItem('refresh-token', refresh);
    
//               return base(originalConfig);
//             } catch (_error) {
  
//               // Logging out the user by removing all the tokens from local
//               localStorage.removeItem('access-token');
//               localStorage.removeItem('refresh-token');
//               // Redirecting the user to the landing page
//               window.location.href = window.location.origin;
//               return Promise.reject(_error);
//             }
//           }
//         }
    
//         return Promise.reject(err);
//       }
// }