import axios, { HeadersDefaults } from 'axios';

const axiosClient = axios.create({
    baseURL:"http://localhost:8080",
    withCredentials: true,
    headers: {
      "Content-Type": "application/json",
    },
})

axiosClient.interceptors.request.use(
  config => {
    const token = localStorage.getItem('access-token');
    if (token) {
      // Configure this as per your backend requirements
      config.headers!['Authorization'] = token;
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

export default axiosClient;
