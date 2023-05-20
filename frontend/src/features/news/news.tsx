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

const instance = newAxiosInstance(config.MusicServerUrl)
const refreshInstance = newAxiosInstance(config.UserServerUrl)
const days: number = 6

export const News: React.FC = () => {
  const [albums, setAlbums] = useState<AlbumView[]>([])
  const [dates, setDates] = useState<number[]>([])
  const { auth, setAuth } = useAuth();
  

  const onRequest = async (config: InternalAxiosRequestConfig): Promise<InternalAxiosRequestConfig> => {
    const decoded: Token = jwt_decode(auth.token)
    console.log(decoded)
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

  const getLatestAlbums = () => {
    let now: Date = new Date()
    let end = new Date()
    end.setDate(now.getDate() - days)
    for (let day = now; day >= end; day.setDate(day.getDate() - 1)) {
      let formattedDate = (moment(day)).format('YYYY-MM-DD')
      let url = config.ReleaseUrl.concat(formattedDate.toString())
      instance.get<AlbumView[]>(url).then(r => r.data)
      .then(
        albums => {
          if (albums.length > 0){
            setDates(prev => [...prev,albums[0].released_at])
            setAlbums(prev =>[...prev,...albums])
          }
        }
      )
      .catch(error => {
          console.log(error)
      });
      console.log(dates)
    }
  }

  useEffect(() => {
    getLatestAlbums()
  }, []);

  return (
    <>
      <Form className="list-form">
      {dates.map(date =>(
          albums.filter(album => album.released_at = date).map(album => (
            <DateRelease albums={albums} date={new Date(album.released_at)}/>
          ))
      ))}
      {}
      </Form>
    </>
  )
}

