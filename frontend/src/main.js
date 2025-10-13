import { createApp } from 'vue'
import App from './App.vue'
import Router from './router'
import '@fortawesome/fontawesome-free/css/all.min.css';
import axios from "axios";

axios.interceptors.response.use(
  res => res,
  err => {
    if (err.response && err.response.status === 401) {
      const msg = err.response.data?.error;
      if (msg === 'token expired') {
        localStorage.removeItem('role');
        window.location.href = '/login';
      }
    }
    return Promise.reject(err);
  }
);

createApp(App).use(Router).mount('#app')
