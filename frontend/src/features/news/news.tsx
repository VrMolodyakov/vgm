import { useEffect, useState } from "react";
import { getRequest } from "../../api/api";
import axios, { InternalAxiosRequestConfig } from "axios";
import { Auth, useAuth } from "../../features/auth/context/auth";
import jwt_decode from 'jwt-decode'
import { Token } from "../../api/token";
import { newAxiosInstance } from "../../api/interceptors";
import config from "../../config/config";

type AlbumsResponse = {
  access_token: string
  refresh_token: string
  logged_in: string
}

const instance = newAxiosInstance(config.MusicServerUrl)
const refreshInstance = newAxiosInstance(config.UserServerUrl)

export const News: React.FC = () => {
  const [albums, setAlbums] = useState()
  const { auth, setAuth } = useAuth();

  const onRequest = async (config: InternalAxiosRequestConfig): Promise<InternalAxiosRequestConfig> => {
    console.log("inside")
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

  const refreshAccessToken = async () => {
    return refreshInstance.get(config.RefreshTokenUrl);
  };

  const getLatestAlbums = async (url: string) => {
    return instance.get(config.NewsUrl).then(r => r.data)
      .catch(error => {
        console.log(error)
      });
  }

  useEffect(() => {
    const albums = getLatestAlbums("")
    console.log(albums)
  }, []);

  return (
    <>
    </>
  )
}