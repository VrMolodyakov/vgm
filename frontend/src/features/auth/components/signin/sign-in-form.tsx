import React from "react";
import { useForm } from "react-hook-form";
import "bootstrap/dist/css/bootstrap.min.css";
import "./sign-in-form.css"


type UserSubmitData = {
  username: string;
  password: string;
};

const Login: React.FC = () => {
  const {
    handleSubmit,
    register,
    setError,
    formState: { errors },
  } = useForm<UserSubmitData>();

  const onSubmit = (data: UserSubmitData) => {
    setError("password",{ type: 'custom', message: 'custom message' })
    console.log(data);
  };

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

export default Login;
