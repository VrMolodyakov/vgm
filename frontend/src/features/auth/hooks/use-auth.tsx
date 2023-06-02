import { useMutation } from "@tanstack/react-query";
import config from "../../../config/config";
import { AxiosError } from "axios";
import { postRequest } from "../../../api/axios/axiosInstance";

type LoginData = {
    username: string
    password: string
}

type RegisterData = {
    username:string
    email:string
    password:string
    role:string
}
  
type TokenResponse = {
    access_token:string 
    refresh_token:string
    logged_in:string
}

const login = async (loginData:LoginData) => {

    const res = await postRequest<TokenResponse>(config.SignInUrl,loginData)
                .then(r => r.data)
                .catch((err: Error | AxiosError) => {
                    return Promise.reject(err)
                  })
    return res
}

const reg = async (regData:RegisterData) => {
    const response = await postRequest(config.SignUpUrl,regData)
    return response
}

export const useUserLogin = () => {
    return useMutation<TokenResponse,AxiosError,LoginData>((data: LoginData) => login(data), {});
};

export const useUserRegister = () => {
    return useMutation<any,AxiosError,RegisterData>((data: RegisterData) => reg(data), {});
};