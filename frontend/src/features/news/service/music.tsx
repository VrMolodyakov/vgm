import { AxiosInstance } from "axios";
import moment from "moment";
import config from "../../../config/config";
import { AlbumView } from "../type";
import { useNavigate } from "react-router-dom";

const days: number = 6

type DatesResponse = {
    dates: number[]
}

export class MusicService {
    navigate = useNavigate();

    client: AxiosInstance
    constructor(client: AxiosInstance) {
        this.client = client
     }

    async getLatest() {
        let res = await this.client.get<DatesResponse>(config.LastDaysUrl + days).then(r => r.data)
        let dates = res.dates.map(date => new Date(date))
        const requests = []
        for (let i = 0; i < dates.length; i++){
            let formattedDate = moment(dates[i]).format('YYYY-MM-DD');
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