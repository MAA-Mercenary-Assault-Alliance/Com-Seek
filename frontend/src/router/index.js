import { createRouter, createWebHistory } from 'vue-router'
import AppHome from '../views/AppHome.vue'
import AppHR from "../views/AppHR.vue"
import AppApplicants from "../views/AppApplicants.vue"
import AppAdmin from "../views/AppAdmin.vue"
import LoginForm from '../views/auth/LoginForm.vue'
import RegisterForm from '../views/auth/RegisterForm.vue'
import StudentProfilePage from '../views/StudentProfile.vue'
import LandingPage from '../views/LandingPage.vue'
import AppJobCreation from "../views/AppJobCreation.vue";
import TermsPage from '../views/docs/Terms.vue'
import PrivacyPage from '../views/docs/Privacy.vue'
import CookiesPage from '../views/docs/Cookies.vue'
import NotFound from '../views/NotFound.vue'
import CompanyRegisterForm from '../views/auth/CompanyRegisterForm.vue'

const routes = [
  { path: '/', redirect: '/landing-page' },

  { path: '/home', component: AppHome },
  { path: '/landing-page', component: LandingPage, meta: { layout: 'blank' } },
  { path: '/hr-dashboard',
    component: AppHR,
    meta: { requiresAuth: true, role: 'company' }
  },
  { path: '/admin',
    component: AppAdmin,
    meta: { requiresAuth: true, role: 'admin' }
  },
  { path: '/applicants/:id',
    name: 'Applicants',
    component: AppApplicants,
    meta: { requiresAuth: true, role: 'company' }
  },
  { path: '/login', component: LoginForm },
  { path: '/logout', component: LoginForm }, //TODO: Implement logout functionality
  { path: '/register', component: RegisterForm },
  { path: '/register-company', component: CompanyRegisterForm },
  {
    path: '/student-profile',
    component: StudentProfilePage,
    name: 'StudentProfile',
    meta: { requiresAuth: true, role: 'student' },
  },
  { path: '/create-job', name: "CreateJob", component: AppJobCreation},

  { path: '/docs/terms', component: TermsPage, meta: { layout: 'blank' } },
  { path: '/docs/privacy', component: PrivacyPage, meta: { layout: 'blank' } },
  { path: '/docs/cookies', component: CookiesPage, meta: { layout: 'blank' } },

  { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound, meta: { layout: 'blank' } },
]

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
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
