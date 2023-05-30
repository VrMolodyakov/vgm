import axios, { AxiosInstance} from "axios";

export function newAxiosInstance(url:string) : AxiosInstance{
  return axios.create({
      baseURL: url,
      withCredentials: true,
      headers: {
        "Content-Type": "application/json",
      },
  })
}