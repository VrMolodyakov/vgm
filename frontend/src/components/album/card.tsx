import { Stack, Col, Card, Row, Image } from "react-bootstrap"
import { Link } from "react-router-dom"
import styles from "./card.module.css"

type AlbumProps = {
    title: string
    id: string
    publisher: string
    imageSrc: string
}

export function AlbumCard({ title, id, publisher, imageSrc }: AlbumProps) {
    return <Card className={`${styles.base}`}>
        <Card.Body as={Link} to={`/${id}`} className={`h-100 text-reset text-decoration-none ${styles.card}`}>
            <Row className="align-items-center justify-content-center h-100">
                <Col xs={2} className="d-flex align-items-center justify-content-center">
                    <Image src={imageSrc} alt="Image" style={{ width: '60px', height: '60px' }} />
                </Col>
                <Col>
                    <Row>                        
                        <span className={`${styles.title}`}>{title}</span>
                        <span className={`${styles.publisher}`}>{publisher}</span>
                    </Row>
                </Col>
            </Row>
        </Card.Body>
    </Card>
}