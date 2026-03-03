import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import UserView from '../views/UserView.vue'
import ProductsView from '@/views/ProductsView.vue'
import OrdersView from '@/views/OrdersView.vue'
import BasketView from '../views/BasketView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/login',
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue'),
    },
    {
      path: '/user',
      name: 'user',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/UserView.vue'),
    },
    {
      path: '/admin/create-user',
      name: 'create-user',
      component: () => import('../views/CreateUserView.vue'),
    },
    {
      path: '/admin/orders',
      name: 'admin-orders',
      component: () => import('../views/AdminOrdersView.vue'),
    },
    {
      path: '/products',
      name: 'products',
      component: () => import('../views/ProductsView.vue'),
    },
    {
      path: '/products/:id',
      name: 'product-details',
      component: () => import('../views/ProductView.vue'),
    },
    {
      path: '/orders',
      name: 'orders',
      component: () => import('../views/OrdersView.vue'),
    },
    {
      path: '/basket',
      name: 'basket',
      component: () => import('../views/BasketView.vue'),
    },
    {
      path: '/listings',
      name: 'listings',
      component: () => import('../views/ListingsView.vue'),
    },
    {
      path: '/seller/orders',
      name: 'seller-orders',
      component: () => import('../views/SellerOrdersView.vue'),
    },
  ],
})

export default router
