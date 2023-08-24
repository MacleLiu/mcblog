import Vue from 'vue'
import App from './App.vue'
import router from './router'
import dayjs from 'dayjs'

import './plugins/http'
import './plugins/antui'

Vue.config.productionTip = false

Vue.filter('dateFormat', (dt) => {
  return dayjs(dt).format('YYYY-MM-DD')
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
