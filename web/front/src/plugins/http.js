import Vue from 'vue'
import axios from 'axios'

// let Url = '后端接口地址'
let Url = 'https://macleliu.com:8888/api/v1'

axios.defaults.baseURL = Url

Vue.prototype.$http = axios

export { Url }