import { useMutation } from "@tanstack/react-query";
import { postRequest } from "../../../api/api";
import config from "../../../config/config";
import axios, { AxiosError } from "axios";

type LoginData = {
    username: string
    password: string
}

type TokenResponse = {
    access_token:string 
    refresh_token:string
    logged_in:string
}

type Code = {
    code:number
}

const login = async (loginData:LoginData) =>{
    const res = await postRequest<TokenResponse>(config.SignInUrl,loginData)
                .then(r => r.data)
                .catch((err: Error | AxiosError) => {
                    if (axios.isAxiosError(err))  {
                        return Promise.reject(err)
                    } else {
                        return Promise.reject(err)                     
                    }
                  })
    return res
}

export const useUserLogin = () => {
    return useMutation<TokenResponse,AxiosError,LoginData>((data: LoginData) => login(data), {});
};