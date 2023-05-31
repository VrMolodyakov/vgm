import axios, { AxiosInstance} from "axios";
import config from "../../config/config";

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
      return res;
    },
    async err => {
      const originalConfig = err.config;
  
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
            window.location.href = window.location.origin;
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