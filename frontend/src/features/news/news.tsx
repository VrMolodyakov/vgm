import { useEffect, useState } from "react";
import { Form, Row, Col } from "react-bootstrap"
import { InternalAxiosRequestConfig } from "axios";
import { Auth, useAuth } from "../../features/auth/context/auth";
import jwt_decode from 'jwt-decode'
import { Token } from "../../api/token";
import { newAxiosInstance } from "../../api/interceptors";
import config from "../../config/config";
import { AlbumView} from "./type";
import { AlbumCard } from "../../components/album/card";
import moment from 'moment';
import "./new.css"
import { DateRelease } from "./date-release";
import { date } from "yup";

const instance = newAxiosInstance(config.MusicServerUrl)
const refreshInstance = newAxiosInstance(config.UserServerUrl)
const days: number = 6

export const News: React.FC = () => {
  const [albums, setAlbums] = useState<AlbumView[]>([])
  const [dates, setDates] = useState<number[]>([])
  const { auth, setAuth } = useAuth();
  

  const onRequest = async (config: InternalAxiosRequestConfig): Promise<InternalAxiosRequestConfig> => {
    const decoded: Token = jwt_decode(auth.token)
    const expireTime = decoded.exp * 1000;
    const now = + new Date();
    if (expireTime > now) {
      config.headers["Authorization"] = 'Bearer ' + auth.token;
    } else {
      const response = await refreshAccessToken();
      const data = response.data;
      const accessToken = data.access_token;
      const auth: Auth = {
        token: accessToken,
        role: decoded.role
      }
      setAuth(() => auth)
      config.headers["Authorization"] = 'Bearer ' + accessToken;
    }
    return config;
  };
  instance.interceptors.request.use(onRequest)

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
      {/* <button onClick={() => {
        console.log(albums)
        console.log(dates)
      }}></button> */}
      <Form className="list-form">
      {dates.map(date =>(
          <DateRelease key={date} albums={albums.filter(album => album.released_at === date)} date={new Date(date)}/>
      ))}
      {}
      </Form>
    </>
  )
}

