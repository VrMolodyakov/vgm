import axios, { AxiosInstance } from "axios";

//TODO:for docker -> process.env.REACT_APP_GATEWAY_URL
const baseAxiosClient = axios.create({
    baseURL:"http://localhost:8080",
    withCredentials: true,
    headers: {
      "Content-Type": "application/json",
    },
})

export function getRequest(URL:string) {
    return baseAxiosClient.get(`/${URL}`).then(response => response);
}
  
export function postRequest<T>(URL:string, payload:any) {
    return baseAxiosClient.post<T>(`/${URL}`, payload).then(response => response);
}
  
export function patchRequest(URL:string, payload:any) {
    return baseAxiosClient.patch(`/${URL}`, payload).then(response => response);
}
  
export function deleteRequest(URL:string) {
    return baseAxiosClient.delete(`/${URL}`).then(response => response);
}
