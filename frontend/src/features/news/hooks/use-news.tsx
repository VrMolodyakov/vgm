
// export const useNews = ()

import { useQuery } from "@tanstack/react-query"
import { MusicService } from "../service/music";

export const useNews = (service: MusicService) => {
    return useQuery(
        ['news'],
        async () => {
            return await service.getLatest();
        }
    );
}