import axios, { AxiosInstance} from "axios";
import config from "../../config/config";

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


export function newAxiosInstance(url:string) : AxiosInstance{
  return axios.create({
      baseURL: url,
      withCredentials: true,
      headers: {
        "Content-Type": "application/json",
      },
  })
}

export function createUserClient(baseURL:string){
  let client = newAxiosInstance(baseURL)
  return client
}

export function createMusicClient(
    baseURL:string,
    refreshURL:string,
    getToken:() => string,
    removeAccessToken:() => void,
    removeRefreshToken:() => void,
    setAccessToken:(token:string) => void,
  ){
  let client = newAxiosInstance(baseURL)
  client.interceptors.request.use(
    config => {
      const token = getToken();
      if (token) {
        config.headers!['Authorization'] = token;
      }
      return config;
    },
    error => {
      return Promise.reject(error);
    }
  )

  client.interceptors.response.use(
    res => {
      console.log("response inter")
      return res;
    },
    async err => {
      const originalConfig = err.config;
      console.log("response error inter")  
      if (originalConfig.url !== config.SignInUrl && err.response) {
        if (err.response.status === 401 && !originalConfig._retry) {
          originalConfig._retry = true;
  
          try {
            const rs = await axios.post(
              refreshURL,
              {
                headers: {
                  Authorization: getToken()
                }
              }
            );
  
            const data = rs.data;
            setAccessToken(data.access_token)
            return client(originalConfig);
          } catch (_error) {

            removeAccessToken()
            removeRefreshToken()
            // Redirecting the user to the landing page
            return Promise.reject(_error);
          }
        }
      }
  
      return Promise.reject(err);
    }
  )
  return client
}

const userInstance = newAxiosInstance(config.UserServerUrl)
export {userInstance}