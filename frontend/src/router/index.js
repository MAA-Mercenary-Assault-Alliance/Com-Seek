import { createRouter, createWebHistory } from 'vue-router'
import AppHome from '../views/AppHome.vue'
import AppProfile from '../views/AppProfile.vue'
import AppHR from "../views/AppHR.vue"
import AppApplicants from "../views/AppApplicants.vue"
import AppAdmin from "../views/AppAdmin.vue"
import LoginForm from '../views/auth/LoginForm.vue'
import RegisterForm from '../views/auth/RegisterForm.vue'
import StudentProfilePage from '../views/StudentProfile.vue'

const routes = [
  { path: '/', component: AppHome },
  { path: '/profile', component: AppProfile },
  { path: '/hr-dashboard', component: AppHR },
  { path: '/applicants', component: AppApplicants },
  { path: '/admin', component: AppAdmin },
  { path: '/applicants/:id', name: 'Applicants', component: AppApplicants },
  { path: '/login', component: LoginForm },
  { path: '/register', component: RegisterForm },
  {
    path: '/student-profile',
    component: StudentProfilePage,
    name: 'StudentProfile',
    meta: { requiresAuth: true },
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const isAuthenticated = !!localStorage.getItem('token') // adjust based on your auth logic

  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login')
  } else {
    next()
  }
})

export default router
