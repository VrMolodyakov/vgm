import React from 'react';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import { validationSchema } from "./validate-scheme"
import "bootstrap/dist/css/bootstrap.min.css";
import "./sign-up-form.css"
import { useUserRegister } from '../auth/hooks/use-auth';

type UserSubmitForm = {
  username: string
  email: string
  password: string
  confirmPassword: string
}

type UserRequest = {
  username: string
  email: string
  password: string
  role: string
}

const SignUpForm: React.FC = () => {
  const { mutate: reg, isSuccess, isError } = useUserRegister()
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors }
  } = useForm<UserSubmitForm>({
    resolver: yupResolver(validationSchema)
  });
  //TODO: react-query
  async function onSubmit(data: UserSubmitForm) {
    console.log(JSON.stringify(data, null, 2));

    let newUser: UserRequest = {
      username: data.username,
      email: data.email,
      password: data.password,
      role: "user"
    }
    reg(newUser)
  };

  return (
    <div className="wrapper">
      <div className="register">
        <form onSubmit={handleSubmit(onSubmit)} className="register-form">
          <div className="form-group">
            <label>Username</label>
            <input
              type="text"
              {...register('username')}
              className={`form-control ${errors.username ? 'is-invalid' : ''}`}
            />
            <div className="invalid-feedback">{errors.username?.message}</div>
          </div>

          <div className="form-group">
            <label>Email</label>
            <input
              type="text"
              {...register('email')}
              className={`form-control ${errors.email ? 'is-invalid' : ''}`}
            />
            <div className="invalid-feedback">{errors.email?.message}</div>
          </div>

          <div className="form-group">
            <label>Password</label>
            <input
              type="password"
              {...register('password')}
              className={`form-control ${errors.password ? 'is-invalid' : ''}`}
            />
            <div className="invalid-feedback">{errors.password?.message}</div>
          </div>

          <div className="form-group">
            <label>Confirm Password</label>
            <input
              type="password"
              {...register('confirmPassword')}
              className={`form-control ${errors.confirmPassword ? 'is-invalid' : ''
                }`}
            />
            <div className="invalid-feedback">
              {errors.confirmPassword?.message}
            </div>
          </div>

          <div className="button-group">
            <button type="submit" className="btn btn-primary btn-wide">
              Register
            </button>
            <button
              type="button"
              onClick={() => reset()}
              className="btn btn-warning btn-wide"
            >
              Reset
            </button>
          </div>
        </form>
      </div>
    </div>

  );
};

export default SignUpForm;