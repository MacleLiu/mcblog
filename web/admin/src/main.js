import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'

import './plugins/antui'

Vue.config.productionTip = false

axios.defaults.baseURL = 'http://localhost:8888/api/v1'
axios.interceptors.request.use(
  (config) => {
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)
// 添加响应拦截器
axios.interceptors.response.use(
  (response) => {
    if (response.data.status === 1004 || response.data.status === 1005 || response.data.status === 1006 || response.data.status === 1007) {
      router.replace({
        path: '/login'
      })
    }
    return response;
  }, 
  (error) => {
    return Promise.reject(error)
  }
);

Vue.prototype.$http = axios


new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
