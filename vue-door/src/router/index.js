import { createRouter, createWebHashHistory } from 'vue-router'
import Login from '../views/Login/Login.vue'
import Dashboard from '../views/Dashboard/Dashboard.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/', name: "Login", component: Login },
    { path: '/dashborad',name: "Dashboard", component: Dashboard },
  ]
})

export default router
