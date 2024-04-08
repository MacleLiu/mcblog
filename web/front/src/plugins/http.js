import Vue from 'vue'
import axios from 'axios'

// let Url = '后端接口地址'

axios.defaults.baseURL = Url

Vue.prototype.$http = axios

export { Url }