import { AlbumCard } from "./card"
import { Row, Col } from "react-bootstrap"
type AlbumListProps = {
    albums:AlbumView[]
}

type AlbumView = {
    id :string
    title:string
    organization:string
}

export function AlbumList({albums}:AlbumListProps){
    return (
        <>
             <Row xs={1} sm={2} lg={3} xl={4} className="g-3">
                {albums.map(album => (
                    <Col key={album.id}>
                    </Col>
                ))}
            </Row>
        </>
    )
}