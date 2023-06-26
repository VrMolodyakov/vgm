import { useQuery } from "@tanstack/react-query";
import "./playlist.css"
import axios from "axios";
import config from "../../config/config";

const baseAxiosClient = axios.create({
    baseURL: "http://localhost:8082",
    withCredentials: true,
    headers: {
        "Content-Type": "application/json",
    },
})

type URL = {
    url: string
}

const PlaylistForm: React.FC = () => {
    const { data } = useQuery<URL>(
        ['playlist'],
        () =>
            baseAxiosClient.get(config.GetPlaylistLink + "/2").then((response) => response.data).catch(e => Promise.reject(e))
    );
    return (

        <div className="playlist-container">
            <div className="playlist-box">
                <p>
                    A link to a random youtube playlist compiled from a collection on the site.
                </p>
                {data ? (
                    <a className="playlist-link" href={data?.url} target="_blank" rel="noopener noreferrer">
                        <span className="playlist-icon"></span>
                        Playlist link
                    </a>
                ) : (
                    <p>Please wait, the link is being created</p>
                )}
            </div>
        </div>
    )
}

export default PlaylistForm