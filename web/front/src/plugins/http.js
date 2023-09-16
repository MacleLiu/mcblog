import Vue from 'vue'
import axios from 'axios'

let Url = 'https://localhost:8888/api/v1'
axios.defaults.baseURL = Url

Vue.prototype.$http = axios

export { Url }