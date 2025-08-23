import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Profile from '../views/Profile.vue'
import Login from "../views/Login.vue";

const routes = [
  { path: '/', component: Home },
  { path: '/profile', component: Profile },
  { path: '/sign-in', component: Login}
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
