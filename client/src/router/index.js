import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    props: true
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.

    // This is what you should always do except for '/' => 'Home'
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
    props: true
  },
  {
    path: '/help',
    name: 'Help',
    component: () => import(/* webpackChunkName: "about" */ '../views/Help.vue'),
<<<<<<< HEAD
<<<<<<< HEAD
    // props: true
=======
    props: true
>>>>>>> Created rules page, and drop down menu -> need to fix the appearence of text on the help page.
=======
    // props: true
>>>>>>> fixed the text display issue
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
