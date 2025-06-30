import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Suggestions from '../views/Suggestions.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/suggestions',
    name: 'Suggestions',
    component: Suggestions
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router

