import Vue from 'vue'
import VueRouter from 'vue-router'
import Lobby from '../views/Lobby.vue'

Vue.use(VueRouter)

const routes = [
  { path: '/', name: 'Lobby', component: Lobby, props: true },
  { path: '/game/:id', name: 'Game', component: () => import('../views/Game.vue'), props: true },
  { path: '/help', name: 'Help', component: () => import('../views/Help.vue'), /* props: true */ },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  
  // for anchor scrolling
  scrollBehavior (to) {
    if (to.hash) {
        return {selector: to.hash}
    } else {
        return { x: 0, y: 0 }
    }
  },
  routes
})

export default router
