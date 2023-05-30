

// // const axiosInstance: AxiosInstance = 

// const instance = axios.create({
//     baseURL: "http://localhost:8080",
//     withCredentials: true,
//     headers: {
//       "Content-Type": "application/json",
//     },
// });

// const refreshInstance = axios.create({
//     baseURL: "http://localhost:8080",
//     withCredentials: true,
//     headers: {
//       "Content-Type": "application/json",
//     },
// });

// instance.interceptors.request.use(
//     async (config) => {
//       const accessToken = localStorage.getItem("access_token");
//       const auth = jwt_decode(accessToken);
//       const expireTime = auth.exp * 1000;
//       const now = + new Date();
//       if (expireTime > now) {
//         config.headers["Authorization"] = 'Bearer ' + accessToken;
//       } else {
//           const response = await refreshAccessToken();
//           const data = response.data;
//           const accessToken = data.access_token;
//           setAuth({token: accessToken});
//           localStorage.removeItem("access_token");
//           localStorage.setItem("access_token", accessToken);
//           config.headers["Authorization"] = 'Bearer ' + accessToken;
//       }
//       console.log("exist from interceptors")
//       return config;
//     },
//     (error) => {
//       console.log(error)
//       console.log("token is expired")
//       }
//   );
