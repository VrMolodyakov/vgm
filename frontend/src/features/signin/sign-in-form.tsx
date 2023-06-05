import React, { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import "bootstrap/dist/css/bootstrap.min.css";
import "./sign-in-form.css"
import { useLocation, useNavigate } from "react-router-dom";
import { useAuthStore } from "../../api/store/store";
import { useUserLogin } from "../auth/hooks/use-auth";
import { Auth, useAuth } from "../auth/context/auth";
import { Token } from "../../api/token";
import jwt_decode from "jwt-decode";


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
  let setAccessToken = useAuthStore(state => state.setAccessToken)
  let setRole = useAuthStore(state => state.setRole)
  let setRefreshToken = useAuthStore(state => state.setRefreshToken)
  const [isRegister,setIsRegister] = useState(false);
  const navigate = useNavigate();
  const location = useLocation();
  const { data, error, mutate: login, isSuccess, isError } = useUserLogin()
  const {
    handleSubmit,
    register,
    setError,
    formState: { errors },
  } = useForm<UserSubmitData>()

  useEffect(() => {
    if (isSuccess) {
      setAccessToken(data.access_token)
      setRefreshToken(data.refresh_token)
      console.log(data.refresh_token)
      const decoded:Token = jwt_decode(data.access_token)
      setRole(decoded.role) 
      navigate("/home")
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
  }

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
