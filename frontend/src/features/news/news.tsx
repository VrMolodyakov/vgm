import { useEffect, useState } from "react";
import { Form} from "react-bootstrap"
import axios, { InternalAxiosRequestConfig } from "axios";
import { Auth, useAuth } from "../../features/auth/context/auth";
import jwt_decode from 'jwt-decode'
import { Token } from "../../api/token";
import config from "../../config/config";
import { AlbumView } from "./type";
import moment from 'moment';
import "./news.css"
import { DateRelease } from "./date-release";
import { useAuthStore } from "../../api/store/store";
import { createMusicClient, newAxiosInstance } from "../../api/axios/axiosInstance";
import { createOnRequestInterceptor, onErrorRequest } from "../../api/axios/interceptors";

const days: number = 6

export const News: React.FC = () => {
  const instance = newAxiosInstance(config.MusicServerUrl)
  const refreshInstance = newAxiosInstance(config.UserServerUrl)
  let getToken = useAuthStore(state => state.getAccessToken)
  let removeAccessToken = useAuthStore(state => state.removeAccessToken)
  let removeRefreshToken = useAuthStore(state => state.removeRefreshToken)
  let setAccessToken = useAuthStore(state => state.setAccessToken)
  const [albums, setAlbums] = useState<AlbumView[]>([])
  const [dates, setDates] = useState<number[]>([])

  let onRequest = createOnRequestInterceptor(getToken)
  instance.interceptors.request.use(onRequest,onErrorRequest)

  let client = createMusicClient(
    config.MusicServerUrl,
    config.UserServerUrl + "/" + config.ReleaseUrl,
    getToken,
    removeAccessToken,
    removeRefreshToken,
    setAccessToken
  )

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

