import { createRouter, createWebHistory } from 'vue-router'
import AuthPage from '@/components/auth/AuthPage.vue'
// import Dashboard from '@/components/Dashboard.vue'
// import Home from '@/components/Home.vue'
// import JobList from '@/components/JobList.vue'
// import Profile from '@/components/Profile.vue'
// import NotFound from '@/components/NotFound.vue'

const routes = [
  // {
  //   path: '/',
  //   name: 'Home',
  //   component: Home
  // },
  {
    path: '/auth',
    name: 'Auth',
    component: AuthPage,
    meta: { requiresGuest: true }
  },
  {
    path: '/login',
    redirect: '/auth'
  },
  {
    path: '/register',
    redirect: '/auth'
  },
  // {
  //   path: '/dashboard',
  //   name: 'Dashboard',
  //   component: Dashboard,
  //   meta: { requiresAuth: true }
  // },
  // {
  //   path: '/jobs',
  //   name: 'Jobs',
  //   component: JobList,
  //   meta: { requiresAuth: true }
  // },
  // {
  //   path: '/profile',
  //   name: 'Profile',
  //   component: Profile,
  //   meta: { requiresAuth: true }
  // },
  // {
  //   path: '/:pathMatch(.*)*',
  //   name: 'NotFound',
  //   component: NotFound
  // }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Navigation guards
router.beforeEach((to, from, next) => {
  const isAuthenticated = localStorage.getItem('userToken')
  
  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/auth')

  // else if (to.meta.requiresGuest && isAuthenticated) {
  //   next('/dashboard')
  // }
  
  } else {
    next()
  }
})

export default router