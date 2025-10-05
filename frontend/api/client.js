import axios from "axios";

export const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || "http://localhost:8080",
  withCredentials: true,
});


api.interceptors.response.use(
  (res) => res,
  (err) => {
    // you can toast here based on err.response?.data?.error
    return Promise.reject(err);
  }
);
