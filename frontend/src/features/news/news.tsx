import { useEffect, useState } from "react";
import { Form} from "react-bootstrap"
import config from "../../config/config";
import { AlbumView } from "./type";
import "./news.css"
import { DateRelease } from "./date-release";
import { useAuthStore } from "../../api/store/store";
import { createMusicClient, newAxiosInstance } from "../../api/axios/axiosInstance";
import { MusicService } from "./service/music";
import { useNews } from "./hooks/use-news";


export const News: React.FC = () => {
  let getToken = useAuthStore(state => state.getAccessToken)
  let removeAccessToken = useAuthStore(state => state.removeAccessToken)
  let removeRefreshToken = useAuthStore(state => state.removeRefreshToken)
  let setAccessToken = useAuthStore(state => state.setAccessToken)
  const [albums, setAlbums] = useState<AlbumView[]>([])
  const [dates, setDates] = useState<number[]>([])

  let client = createMusicClient(
    config.MusicServerUrl,
    config.UserServerUrl + "/" + config.ReleaseUrl,
    getToken,
    removeAccessToken,
    removeRefreshToken,
    setAccessToken
  )
  
  let musicService = new MusicService(client)
  const { data, isLoading, isError } = useNews(musicService);

  useEffect(() => {
    if (!isLoading && !isError && data) {
      const albums = data.flat()
      console.log(albums)
      setAlbums(albums)
      setDates(albums.map(album => album.released_at))
    }    
  }, [data]);

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

