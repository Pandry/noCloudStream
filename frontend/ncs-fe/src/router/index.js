import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import SignupView from '../views/SignupView.vue'
import { store } from '@/store/store.js'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {requiresAuth: true} /*,
      beforeEnter: (to, from, next) =>{
        if (!$store.getters.userIsLoggedIn){
          next()
        }else{
          //return { name: 'login' }
          next({ name: 'login' })
        }
      }*/
    },
    {
      path: '/about',
      name: 'about',
      meta: {requiresAuth: true}, 
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/signup',
      name: 'signup',
      component: SignupView,
      meta: {isAuth: true}
    },{
      path: '/login',
      name: 'login',
      component: LoginView,
      meta: {isAuth: true}
    },
  ]
})

/*
https://router.vuejs.org/guide/advanced/navigation-guards.html#global-before-guards
router.beforeEach(async (to, from) => {
  if (
    // make sure the user is authenticated
    !isAuthenticated &&
    // ❗️ Avoid an infinite redirect
    to.name !== 'Login'
  ) {
    // redirect the user to the login page
    return { name: 'Login' }
  }
})*/


router.beforeEach((to, from, next) => {
  // Get value from somewhere to determine if the user is 
  // logged in or not
  const userIsLoggedIn = store.getters.userIsLoggedIn;
  const pageRequiresAuth = to.matched.some(record => record.meta.requiresAuth)
  const pageIsAuthPage = to.matched.some(record => record.meta.isAuth)

  let redirect = undefined

  // Determine if the route requires authentication
  if (pageRequiresAuth) {
      // If user is not logged in, navigate to the named "login" route 
      // with a query string parameter indicating where to navigate to after
      // successful login
      if (!userIsLoggedIn) {
          // Navigate to login route
          redirect = {
              name: "login",
              query: {redirect: to.fullPath}
          };
      }
  } else if (pageIsAuthPage && userIsLoggedIn) {
    redirect={name: "home"}
  }
  next(redirect)
});

export default router
