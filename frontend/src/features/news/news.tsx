import { useEffect, useState } from "react";
import { Form,Row, Col} from "react-bootstrap"
import { InternalAxiosRequestConfig } from "axios";
import { Auth, useAuth } from "../../features/auth/context/auth";
import jwt_decode from 'jwt-decode'
import { Token } from "../../api/token";
import { newAxiosInstance } from "../../api/interceptors";
import config from "../../config/config";
import { AlbumView, Albums } from "./type";
import { AlbumCard } from "../../components/album/card";

const instance = newAxiosInstance(config.MusicServerUrl)
const refreshInstance = newAxiosInstance(config.UserServerUrl)

export const News: React.FC = () => {
  const [albums, setAlbums] = useState<AlbumView[]>([])
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
    return instance.get<AlbumView[]>(config.NewsUrl).then(r => r.data).then(
      albums => setAlbums(() => albums)
    )
      .catch(error => {
        console.log(error)
      });
  }

  useEffect(() => {
    const albums = getLatestAlbums()
  }, []);

  return (
    <>
      <Form className="list-form">
        <Row xs={5} sm={5} lg={5} xl={5} className="g-3">
          {albums.map(album => (
            <Col key={album.album_id}>
              <AlbumCard title={album.title} id={album.album_id} publisher={album.publisher} imageSrc={album.small_image_src}/>
            </Col>
          ))}
        </Row>
      </Form>
    </>
  )
}