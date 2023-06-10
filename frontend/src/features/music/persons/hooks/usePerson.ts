import { AxiosError } from "axios";
import { MusicService } from "../../service/music";
import { Person } from "../types";
import { useMutation } from "@tanstack/react-query";

export const usePerson = (service: MusicService) => {
    return useMutation<any,AxiosError,Person>((data: Person) => service.createPerson(data), {});
};