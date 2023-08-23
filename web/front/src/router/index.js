import Vue from 'vue'
import VueRouter from 'vue-router'

import Front from '../views/Front'

// 子页面组件
import FrontIndex from '../components/front/FrontIndex'
import FrontPigeonhole from '../components/front/FrontPigeonhole'
import FrontWork from '../components/front/FrontWork'
import FrontTool from '../components/front/FrontTool'
import FrontAbout from '../components/front/FrontAbout'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        component: Front,
        redirect: '/index',
        children: [
            {path: 'index', component: FrontIndex},
            {path: 'pigeonhole', component: FrontPigeonhole},
            {path: 'work', component: FrontWork},
            {path: 'tool', component: FrontTool},
            {path: 'about', component: FrontAbout},
        ]
    },
]

const router = new VueRouter({
    routes
})

export default router