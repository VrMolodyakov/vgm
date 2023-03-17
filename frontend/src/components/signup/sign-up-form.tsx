import React from 'react';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import {validationSchema} from "./validate-scheme"
import "bootstrap/dist/css/bootstrap.min.css";
import "./sign-up-form.css"
import { getRequest, postRequest } from '../../api/api';

type UserSubmitForm = {
  fullname: string;
  username: string;
  email: string;
  password: string;
  confirmPassword: string;
};

const SignUpForm: React.FC = () => {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors }
  } = useForm<UserSubmitForm>({
    resolver: yupResolver(validationSchema)
  });

  async function onSubmit(data: UserSubmitForm){
    console.log(JSON.stringify(data, null, 2));
    const response = await getRequest("ping")
    console.log(response)
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
            <i className="fa-solid fa-envelope"></i>
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
              className={`form-control ${
                errors.confirmPassword ? 'is-invalid' : ''
              }`}
            />
            <div className="invalid-feedback">
              {errors.confirmPassword?.message}
            </div>
          </div>
          <div className="form-group">
            <button type="submit" className="btn btn-primary">
              Register
            </button>
            <button
              type="button"
              onClick={() => reset()}
              className="btn btn-warning float-right"
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