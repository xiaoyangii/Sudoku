import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'homepage',
    component: () => import('@/views/homepage.vue')
  },
  {
    path: '/select',
    name: 'Select',
    component: () => import('@/views/select_view.vue')
  },
  {
    path: '/game/:lever',
    name: 'Game',
    component: () => import('@/views/game_view.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
