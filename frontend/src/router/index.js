import { createRouter, createWebHistory } from 'vue-router'
import AppHome from '../views/AppHome.vue'
import AppProfile from '../views/AppProfile.vue'
import AppLogin from "../views/AppLogin.vue"
import AppHR from "../views/AppHR.vue";
import AppApplicants from "../views/AppApplicants.vue";
import AppAdmin from "../views/AppAdmin.vue";

const routes = [
  { path: '/', component: AppHome },
  { path: '/profile', component: AppProfile },
  { path: '/sign-in', component: AppLogin },
  { path: '/hr-dashboard', component: AppHR },
  { path: '/applicants', component: AppApplicants },
  { path: '/admin', component: AppAdmin },

  { path: '/applicants/:id', name: 'Applicants', component: AppApplicants },
]


const router = createRouter({
  history: createWebHistory(),
  routes: routes,
})

export default router
