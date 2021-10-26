import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/LoginComponent.vue'
import Home from '../views/Home.vue'

const routes = [
//   {
//     path: '/',
//     name: 'Home',
//     component: Home //component: () => import('../views/About.vue')
//   },
// Match all paths vue2 Use * vue3 Use /:pathMatch(.*)* or /:pathMatch(.*) or /:catchAll(.*)
  {
    path: '/:catchAll(.*)',
    name: 'signin',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: Login
  },
  {
    path: '/home',
    name: 'home',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: Home
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
