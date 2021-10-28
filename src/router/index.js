import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/LoginComponent.vue'
import Register from '../views/RegisterComponent.vue'
import Home from '../views/Home.vue'

const routes = [
	{
		path: '/',
		name: 'Home',
		component: Home //component: () => import('../views/About.vue')
	},
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
		path: '/register',
		name: 'register',
		// route level code-splitting
		// this generates a separate chunk (about.[hash].js) for this route
		// which is lazy-loaded when the route is visited.
		component: Register
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

router.beforeEach((to, from, next) => {
	const publicPages = ['/login', '/register'];
	const authRequired = !publicPages.includes(to.path);
	const loggedIn = localStorage.getItem('user');

	// trying to access a restricted page + not logged in
	// redirect to login page
	if (authRequired && !loggedIn) {
		next('/login');
	} else {
		next();
	}
});

export default router
