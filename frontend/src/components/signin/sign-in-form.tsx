import React from "react";
import { useForm } from "react-hook-form";
import "bootstrap/dist/css/bootstrap.min.css";
import "./sign-in-form.css"
import { postRequest } from "../../api/api";


type UserSubmitData = {
  username: string
  password: string
}

const SignInForm: React.FC = () => {
  const {
    handleSubmit,
    register,
    setError,
    formState: { errors },
  } = useForm<UserSubmitData>()

  //TODO:check error
  async function onSubmit(data: UserSubmitData){
    // setError("password",{ type: 'custom', message: 'custom message' })
    const response = await postRequest("auth/login",data).catch(error =>{
      console.log("inside")
      console.log(error)
    })
    console.log(response)  

  }

  return (
    <div className="wrapper">
      <div className="login">
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="form-group">
            <input id="username" type="username" required={true} {...register("username")}></input>
          </div>
          <div className="form-group">
            <label>Your password</label>
            <input  id="password" className={`form-control ${errors.password ? "is-invalid" : ""}`} {...register("password")}
            ></input>
            {errors.password && (
              <small className="text-danger">{errors.password.message}</small>
            )}
          </div>
          <button type="submit">Submit</button>
        </form>
      </div>
    </div>
  );
};

export default SignInForm;
