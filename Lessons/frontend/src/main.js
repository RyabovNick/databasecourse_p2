import Vue from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'
import vuetify from './plugins/vuetify'
import dayjs from 'dayjs'
import axios from 'axios'
// установить для axios стандартный url
axios.defaults.baseURL = 'http://localhost:80'

Vue.config.productionTip = false

Vue.filter('formatDate', function(value) {
  if (value) {
    return dayjs(String(value)).format('DD/MM/YYYY')
  }
})

// добавляем глобальный объект axios
Vue.prototype.$axios = axios

new Vue({
  router,
  vuetify,
  store,
  render: h => h(App),
}).$mount('#app')
