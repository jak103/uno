import Vue from 'vue'
import VueRouter from 'vue-router'
import Lobby from '../views/Lobby.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Lobby',
    component: Lobby,
    props: true
  },
  {
    path: '/game/:id',
    name: 'Game',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.

    // This is what you should always do except for '/' => 'Home'
    component: () => import(/* webpackChunkName: "about" */ '../views/Game.vue'),
    props: true
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
