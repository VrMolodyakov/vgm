import { useEffect, useState } from "react";
import { Form} from "react-bootstrap"
import { InternalAxiosRequestConfig } from "axios";
import { Auth, useAuth } from "../../features/auth/context/auth";
import jwt_decode from 'jwt-decode'
import { Token } from "../../api/token";
import config from "../../config/config";
import { AlbumView } from "./type";
import moment from 'moment';
import "./news.css"
import { DateRelease } from "./date-release";
import { useAuthStore } from "../../api/store/store";
import { newAxiosInstance } from "../../api/axios/axiosInstance";
import { createOnRequestInterceptor, onErrorRequest } from "../../api/axios/interceptors";

const days: number = 6

export const News: React.FC = () => {
  const instance = newAxiosInstance(config.MusicServerUrl)
  const refreshInstance = newAxiosInstance(config.UserServerUrl)
  let getToken = useAuthStore(state => state.getToken)
  const [albums, setAlbums] = useState<AlbumView[]>([])
  const [dates, setDates] = useState<number[]>([])


  // const onRequest = async (config: InternalAxiosRequestConfig): Promise<InternalAxiosRequestConfig> => {
  //   const token = getToken()
  //   const decoded: Token = jwt_decode(token);
  //   const expireTime = decoded.exp * 1000;
  //   const now = +new Date();
  //   if (expireTime > now) {
  //     config.headers["Authorization"] = 'Bearer ' + token;
  //   } else {
  //     console.log("refresh");
  //     const response = await refreshAccessToken();
  //     const data = response.data;
  //     const accessToken = data.access_token;
  //     const newAuth: Auth = {
  //       token: accessToken,
  //       role: decoded.role
  //     };
  //     setAuth(() => newAuth);
  //     config.headers["Authorization"] = 'Bearer ' + accessToken;
  //   }
  //   return config;
  // }
  let onRequest = createOnRequestInterceptor(getToken)
  instance.interceptors.request.use(onRequest,onErrorRequest)


  instance.interceptors.response.use(
    res => {
      return res;
    },
    async err => {
      const originalConfig = err.config;
  
      if (originalConfig.url !== '/user/login' && err.response) {
        // Access Token was expired
        if (err.response.status === 401 && !originalConfig._retry) {
          originalConfig._retry = true;
  
          try {
            const rs = await axios.post(
              'https://api.example.org/user/refresh',
              {
                headers: {
                  Authorization: localStorage.getItem('refresh-token')!
                }
              }
            );
  
            const access = rs.data.data['X-Auth-Token'];
            const refresh = rs.data.data['X-Refresh-Token'];
  
            localStorage.setItem('access-token', access);
            localStorage.setItem('refresh-token', refresh);
  
            return axiosClient(originalConfig);
          } catch (_error) {
            toast.error('Session time out. Please login again.', {
              id: 'sessionTimeOut'
            });
            // Logging out the user by removing all the tokens from local
            localStorage.removeItem('access-token');
            localStorage.removeItem('refresh-token');
            // Redirecting the user to the landing page
            window.location.href = window.location.origin;
            return Promise.reject(_error);
          }
        }
      }
  
      return Promise.reject(err);
    }
  );
  



  const refreshAccessToken = () => {
    return refreshInstance.get(config.RefreshTokenUrl);
  };

  const getLatestAlbums = async () => {
    let now: Date = new Date();
    let end = new Date();
    end.setDate(now.getDate() - days);

    const requests = [];

    for (let day = now; day >= end; day.setDate(day.getDate() - 1)) {
      let formattedDate = moment(day).format('YYYY-MM-DD');
      let url = config.ReleaseUrl.concat(formattedDate.toString());
      requests.push(
        instance.get<AlbumView[]>(url).then(response => response.data).catch(error => {
          console.log(error);
          return [];
        })
      );
    }
    const req = await Promise.all(requests);
    const albumsArray = req.flat();
    const newDates = albumsArray.map(album => album.released_at);
    setAlbums(prev => [...prev, ...albumsArray]);
    setDates(prev => [...prev, ...newDates]);
  };

  useEffect(() => {
    getLatestAlbums()
  }, []);

  return (
    <>
      <Form className="list-form">
        <div className="centered-block">
          <h2>New's release</h2>
          {dates.map(date => (
            <DateRelease key={date} albums={albums.filter(album => album.released_at === date)} date={new Date(date)} />
          ))}
        </div>
      </Form>
    </>
  )
}

