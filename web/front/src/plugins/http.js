import Vue from 'vue'
import axios from 'axios'

let Url = 'http://224bceab.r16.cpolar.top/api/v1'
axios.defaults.baseURL = Url

Vue.prototype.$http = axios

export { Url }