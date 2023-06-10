import { useQuery } from "@tanstack/react-query";
import { MusicService } from "../../service/music";

export const usePersonList = (service: MusicService) => {
    return useQuery(
        ['persons'],
        async () => {
            return await service.getPersons();
        }
    );
}