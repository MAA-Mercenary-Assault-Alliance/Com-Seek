import { createRouter, createWebHistory } from 'vue-router'
import AppHome from '../views/AppHome.vue'
import AppProfile from '../views/AppProfile.vue'
import AppHR from "../views/AppHR.vue"
import AppApplicants from "../views/AppApplicants.vue"
import AppAdmin from "../views/AppAdmin.vue"
import LoginForm from '../views/auth/LoginForm.vue'
import RegisterForm from '../views/auth/RegisterForm.vue'
import StudentProfilePage from '../views/StudentProfile.vue'
import LandingPage from '../views/LandingPage.vue'

const routes = [
  { path: '/', redirect: '/landing-page' },

  { path: '/home', component: AppHome },
  { path: '/landing-page', component: LandingPage, meta: { layout: 'blank' } },
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
    meta: { requiresAuth: true, role: 'student' },
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const email = localStorage.getItem('email');
  const role = localStorage.getItem('role');
  const isAuthenticated = !!(email && role);

  if (to.meta.requiresAuth && !isAuthenticated) {
    return next('/login');
  }

  if (to.meta.role && isAuthenticated) {
    const required = to.meta.role; // can be string or array
    const allowed =
      Array.isArray(required) ? required.includes(role) : role === required;

    if (!allowed) {
      return next('/');
    }
  }

  next();
});


export default router
