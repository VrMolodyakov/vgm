import { useEffect, useState } from "react";
import { Form} from "react-bootstrap"
import { AlbumView } from "./type";
import "./news.css"
import { DateRelease } from "./date-release";
import { MusicService } from "../service/music";
import { useNews } from "./hooks/use-news";
import { useMusicClient } from "../client-provider/context/context";


export const News: React.FC = () => {
  const [albums, setAlbums] = useState<AlbumView[]>([])
  const [dates, setDates] = useState<number[]>([])
  let client = useMusicClient()
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
    <div className="album-container">
      <Form className="list-form">
        <div className="centered-block">
          <div className="release">
          <h2>New's release</h2>
          </div>
          {dates.map(date => (
            <DateRelease key={date} albums={albums.filter(album => album.released_at === date)} date={new Date(date)} />
          ))}
        </div>
      </Form>
    </div>
  )
}

