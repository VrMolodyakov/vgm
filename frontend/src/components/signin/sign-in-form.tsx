import React, { useContext, useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import "bootstrap/dist/css/bootstrap.min.css";
import "./sign-in-form.css"
import { postRequest } from "../../api/api";
import { AuthContextType } from "../../features/auth/types/auth-context-type";
import AuthContext from "../../features/auth/context/auth";
import { useLocation, useNavigate } from "react-router-dom";
import { LocalStorage } from "../../features/local-storage/service/service";
import useAuth from "../../features/auth/hooks/use-auth";


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
  const { auth,setAuth } = useAuth();
  const [isRegister,setIsRegister] = useState(false);
  const navigate = useNavigate();
  const location = useLocation();
  const {
    handleSubmit,
    register,
    setError,
    formState: { errors },
  } = useForm<UserSubmitData>()

  const getToken = async (userData:UserSubmitData) =>{
    return postRequest<TokenResponse>("auth/login",userData).then(r => r.data)
    .catch(error => {   
      if (error.response.status === 400){
        setError("root",{type:'custom',message:"wrong username or password"})
      }else{
        setError("root",{type:'custom',message:"internal server error"})
      }
    });
  }

  function onSubmit(data: UserSubmitData){
    (async() => {
      const response = await getToken(data);
      console.log("response: ",response)
      if (response) {
        const accessToken = response.access_token
        setAuth("accessToken")
      }
    })();
  }

  useEffect(() => {
    if (auth !== "" && auth !== null) {
      LocalStorage.set("access_token", auth)
      console.log("access")
      // navigate("/home");
    }
  }, [auth]);

  useEffect(() => {
    setAuth("abc")
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
