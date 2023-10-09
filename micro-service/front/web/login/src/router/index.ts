import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import("../views/login.vue")
    },
    {
      path: '/register',
      name: 'register',
      component: () => import("../views/login.vue")
    },
    {
      path: '/test',
      name: 'test',
      component: () => import("../views/test.vue")
    },
    {
      path: '/mail',
      name: 'mail',
      component: () => import("../views/mail.vue"),
    },
    {
      path: '/writeEmail',
      component: () => import("../components/writeEmail.vue"),
    },
  ]
})

export default router
