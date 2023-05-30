import React, { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import "bootstrap/dist/css/bootstrap.min.css";
import "./sign-in-form.css"
import { postRequest } from "../../api/api";
import { useLocation, useNavigate } from "react-router-dom";
import { Auth, useAuth } from "../../features/auth/context/auth";
import jwt_decode from "jwt-decode";
import { Token } from "../../api/token";
import config from "../../config/config";
import { useAuthStore } from "../../api/store/store";
import { useUserLogin } from "../../features/auth/hooks/use-auth";
import { AxiosError } from "axios";


type UserSubmitData = {
  username: string
  password: string
}

type TokenResponse = {
  access_token:string 
  refresh_token:string
  logged_in:string
}

const SignInForm: React.FC = () => {
  
  let setToken = useAuthStore(state => state.setToken)

  const { auth,setAuth } = useAuth();
  const [isRegister,setIsRegister] = useState(false);
  const navigate = useNavigate();
  const location = useLocation();
  const { data, error, mutate: login, isSuccess, isError } = useUserLogin();
  const {
    handleSubmit,
    register,
    setError,
    formState: { errors },
  } = useForm<UserSubmitData>()

  // const getToken = async (userData:UserSubmitData) =>{
  //   return postRequest<TokenResponse>(config.SignInUrl,userData).then(r => r.data)
  //   .catch(error => {   
  //     if (error.response.status === 400){
  //       setError("root",{type:'custom',message:"wrong username or password"})
  //     }else{
  //       setError("root",{type:'custom',message:"internal server error"})
  //     }
  //   });
  // }

  useEffect(() => {
    if (isSuccess) {
      setToken(data.access_token)
    } else if (isError) {
      if (error){
        if (error.response?.status === 400){
          setError("root",{type:'custom',message:"wrong username or password"})
        }else{
          setError("root",{type:'custom',message:"internal server error"})
        }
      }
    } else return;
  }, [isSuccess, isError]);

  function onSubmit(userData: UserSubmitData){
    login(userData)
    navigate("/home")
    // (async() => {
    //   const response = await getToken(data);
    //   if (response) {
    //     setToken(response.access_token)
    //     const accessToken = response.access_token
    //     const decoded:Token = jwt_decode(accessToken);
    //     const auth:Auth = {
    //       token:accessToken,
    //       role:decoded.role
    //     } 
    //     setAuth(() => auth)
    //     navigate("/home");
    //   }
    // })();
  }

  // useEffect(() => {
  //   if (auth !== "" && auth !== null) {
  //     LocalStorage.set("access_token", auth)

  //   }
  // }, [auth]);

  useEffect(() => {
    if (location.state?.previousUrl === "/reg"){
      setIsRegister(true)
    }
  }, []);

  return (
    <div className="wrapper">
      <div className="login">
        <form onSubmit={handleSubmit(onSubmit)}>
          {isRegister && (
                          <div className="alert alert-success" role="alert">
                          {"you have been successfully registered"}
                          </div>
                         )
          }
          <div className="form-group">
            <input id="username" type="username" required={true} {...register("username")}></input>
          </div>
          <div className="form-group">
            <label>Your password</label>
            <input  id="password" className={`form-control ${errors.password ? "is-invalid" : ""}`} {...register("password")}
            ></input>
          </div>
          <button type="submit">Submit</button>
          {errors.root && (
              <small className="text-danger">{errors.root.message}</small>
          )}
        </form>
      </div>
    </div>
  );
};

export default SignInForm;
