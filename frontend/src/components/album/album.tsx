import { Form, Stack, Row, Col, Button, Card, Badge } from "react-bootstrap"
import { Link } from "react-router-dom"
import styles from "./card.module.css"

type AlbumProps = {
    title:string
    id:string
    organization:string
}

function AlbumCard({ title, id ,organization }: AlbumProps) {
    return <Card>
        <Card.Body as={Link} to={`/${id}`} className={`h-100 text-reset text-decoration-none ${styles.card}`}>
            <Stack gap={2} className="align-items-center justify-content-center h-100">
                <span className="fs-5">{title}</span>
            </Stack>
        </Card.Body>
    </Card>
}