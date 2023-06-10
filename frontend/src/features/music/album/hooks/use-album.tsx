import { AxiosError } from "axios"
import { useMutation } from "@tanstack/react-query"
import { MusicService } from "../../service/music";
import { FullAlbum } from "../types";

export const useAlbum = (service: MusicService) => {
    return useMutation<any,AxiosError,FullAlbum>((data: FullAlbum) => service.createAlbum(data), {});
};