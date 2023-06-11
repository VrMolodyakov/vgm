import { Col, Row } from "react-bootstrap";
import { AlbumView} from "./type";
import "./date-release.css";

type DateReleaseProps = {
    albums:AlbumView[]
    date:Date
}

export function DateRelease({ albums, date }: DateReleaseProps) {
  return (
    <Row xs={5} sm={5} lg={5} xl={5} className="g-3">
      <div className="dateblock">
        <Row className="date">
          <span>{date.toLocaleString('en-US', { month: 'short' })}</span>
        </Row>
        <Row className="day">
          <span>{date.getDay()}</span>
        </Row>
      </div>
      {albums.map(album => (
        <Col key={album.album_id} className="album-col">
          <div className="album-card">
            <div className="album-image">
              <img src={album.small_image_src} alt={album.title} />
            </div>
            <div className="album-details">
              <div className="album-title">{album.title}</div>
              <div className="album-publisher">{album.publisher}</div>
            </div>
          </div>
        </Col>
      ))}
    </Row>
  )
}