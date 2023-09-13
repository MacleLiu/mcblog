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
Vue.filter('timeFormat', (dt) => {
  return dayjs(dt).format('YYYY-MM-DD HH:mm:ss')
})

// 修改网页title的指令
Vue.directive('title', {
  inserted: function (el, binding) {
    document.title = el.dataset.title + ' - Macle'
  }
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
