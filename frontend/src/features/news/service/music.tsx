import { AxiosInstance } from "axios";
import moment from "moment";
import config from "../../../config/config";

const days: number = 6

export class MusicService {
    constructor(private client: AxiosInstance) { }

    getLatest() {
        let now: Date = new Date();
        let end = new Date();
        end.setDate(now.getDate() - days);

        const requests = [];

        for (let day = now; day >= end; day.setDate(day.getDate() - 1)) {
            let formattedDate = moment(day).format('YYYY-MM-DD');
            let url = config.ReleaseUrl.concat(formattedDate.toString());
            requests.push(
                client.get<AlbumView[]>(url).then(response => response.data).catch(error => {
                    console.log(error);
                    return [];
                })
            );
        }
        const req = await Promise.all(requests);
        const albumsArray = req.flat();
        const newDates = albumsArray.map(album => album.released_at);
    }

}