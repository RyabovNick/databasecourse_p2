import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import dayjs from 'dayjs'

Vue.config.productionTip = false

Vue.filter('formatDate', function(value) {
  if (value) {
    return dayjs(String(value)).format('DD/MM/YYYY')
  }
})

new Vue({
  router,
  vuetify,
  render: h => h(App),
}).$mount('#app')
