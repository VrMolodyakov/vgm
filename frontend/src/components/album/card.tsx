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
            <div className={styles.container}>
                <div className={styles.imageContainer}>
                    <Image src={imageSrc} alt="Image" className={styles.image} />
                </div>
                <div className={styles.content}>
                    <div className={styles.title}>{title}</div>
                    <div className={styles.publisher}>{publisher}</div>
                </div>
            </div>
        </Card.Body>
    </Card>
}