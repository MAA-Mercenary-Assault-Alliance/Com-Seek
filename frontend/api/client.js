import axios from "axios";
import router from "../src/router";
const baseURL = import.meta.env.VITE_API_URL || "http://localhost:8000";

export const api = axios.create({
  baseURL,
  withCredentials: true,
});

api.interceptors.response.use(
  (res) => res,
  (err) => {
    if (err.response && err.response.status === 401) {
      console.log("Unauthorized! Redirecting to login...");
      const msg = err.response.data?.error;
      if (msg === "token expired") {
        localStorage.removeItem("role");
      }
      router.push("/login");
    }
    return Promise.reject(err);
  },
);
