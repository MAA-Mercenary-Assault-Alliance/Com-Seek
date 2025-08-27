import { createRouter, createWebHistory } from 'vue-router'
import AppHome from '../views/AppHome.vue'
import AppProfile from '../views/AppProfile.vue'
import AppLogin from "../views/AppLogin.vue";

const routes = [
  { path: '/', component: AppHome },
  { path: '/profile', component: AppProfile },
  { path: '/sign-in', component: AppLogin}
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
