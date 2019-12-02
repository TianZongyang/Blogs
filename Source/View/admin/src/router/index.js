import Vue from 'vue'
import VueRouter from 'vue-router'
import Dashboard from "../views/Dashboard";
import NotFound from "../views/NotFound";
import AddArticle from "../views/AddArticle";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/dashboard'
  },
  {
    path: '/dashboard',
    name: 'DashBoard',
    component: Dashboard
  },
  {
    path: '/addArticle',
    name: 'AddArticle',
    component: AddArticle
  },
  {
    path: '*',
    name: 'NotFound',
    component: NotFound
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
