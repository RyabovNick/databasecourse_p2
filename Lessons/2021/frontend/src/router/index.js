import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Menu from '../views/Menu.vue'
import SignIn from '../views/SignIn.vue'
import UserOrder from '../views/UserOrder.vue'
import store from '../store'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
  {
    path: '/sign_in',
    name: 'SignIn',
    component: SignIn,
    meta: {
      noAuth: true,
    },
  },
  {
    path: '/menu',
    name: 'Menu',
    component: Menu,
  },
  {
    path: '/user_order',
    name: 'UserOrder',
    component: UserOrder,
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
})

/**
 * 1. Проверить, что пользователь авторизован (ищем в store.state.auth.isAuth)
 *    а) Если !isAuth - ищем токен в localStorage.
 *        Если найден - устанавливаем axios default header, isAuth = true.
 *    1. Пользователь не авторизован:
 *       а) Редирект на страницу /sign_in
 *       б) Пользователь успешно вводит логин и пароль
 *       в) Сохранение токена в localStorage
 *       г) Установка для axios header по умолчанию
 *       д) Редирект на страницу (либо нашу, либо лучший вариант, на страницу,
 *         на которую пользователь пытался попасть изначально)
 *    2. Пользователь авторизован: заходит на страницу
 */
router.beforeEach((to, from, next) => {
  // TODO: не давать возможность авторизованным пользователям
  // переходит на страницу sign_in (sign_up)
  console.log('to.meta.noAuth', to.meta.noAuth)
  // если не требует авторизации
  if (to.meta.noAuth) {
    // и если пользователь авторизован
    console.log('store.state.auth.isAuth: ', store.state.auth.isAuth)
    if (store.state.auth.isAuth) {
      // и пытается попасть на sign_in
      console.log('to.name: ', to.name)
      if (to.name === 'SignIn') {
        // редирект на главную
        next({ name: 'Home' })
        return
      }
    }
    // если не авторизован, либо пытается попасть не на sign_in
    next()
    return
  }

  const token = localStorage.getItem('token')
  console.log('token: ', token)
  if (token) {
    store.dispatch('auth/setAuth', token)
    next()
  } else {
    next({ name: 'SignIn' })
  }
})

export default router
