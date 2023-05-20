import { Col, Row } from "react-bootstrap";
import { AlbumCard } from "../../components/album/card";
import { AlbumView} from "./type";
import "./new.css";

type DateReleaseProps = {
    albums:AlbumView[]
    date:Date
}

export function DateRelease({albums,date}:DateReleaseProps){
    return  <Row xs={5} sm={5} lg={5} xl={5} className="g-3">
    <div className="dateblock">
      <Row className="date">
        <span>{date.getMonth()}</span>
      </Row>
      <Row className="day">
        <span>{date.getDay()}</span>
      </Row>
    </div>
    {albums.map(album => (
      <Col key={album.album_id} className="album-col">
        <AlbumCard title={album.title} id={album.album_id} publisher={album.publisher} imageSrc={album.small_image_src} />
      </Col>
    ))}
  </Row>
}