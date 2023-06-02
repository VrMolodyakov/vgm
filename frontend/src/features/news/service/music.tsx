import { AxiosInstance } from "axios";
import moment from "moment";
import config from "../../../config/config";
import { AlbumView } from "../type";
import { useNavigate } from "react-router-dom";

const days: number = 6

export class MusicService {
    navigate = useNavigate();

    client: AxiosInstance
    constructor(client: AxiosInstance) {
        this.client = client
     }

    async getLatest() {
        let now: Date = new Date()
        let end = new Date()
        end.setDate(now.getDate() - days)

        const requests = []

        for (let day = now; day >= end; day.setDate(day.getDate() - 1)) {
            let formattedDate = moment(day).format('YYYY-MM-DD');
            let url = config.ReleaseUrl.concat(formattedDate.toString());
            console.log(url)
            requests.push(
                this.client.get<AlbumView[]>(url).then(response => response.data).catch(error => {
                    console.log(error)
                    this.navigate("/auth")
                    throw error
                })
            )
        }
        return await Promise.all(requests)
    }

}