import Vue from 'vue'
import VueRouter from 'vue-router'

import Front from '../views/Front'
import Index from '../views/Index'

// 子页面组件
import FrontPigeonhole from '../components/pigeonhole/FrontPigeonhole'
import FrontWork from '../components/works/FrontWork'
import FrontTool from '../components/tools/FrontTool'
import FrontAbout from '../components/about/FrontAbout'

import ArticleRead from '../components/articles/ArticleRead'
import CateList from '../components/articles/CateList'
import TagList from '../components/articles/TagList'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        component: Index,
    },
    {
        path: '/blog',
        component: Front,
        redirect: '/',
        children: [
            {path: 'pigeonhole', component: FrontPigeonhole},
            {path: 'work', component: FrontWork},
            {path: 'tool', component: FrontTool},
            {path: 'about', component: FrontAbout},
            {path: 'cate/:cid', component: CateList, props: true},
            {path: 'tag/:tid', component: TagList, props: true},
        ]
    },
    {
        path: '/blog/article/:id',
        component: ArticleRead,
        props: true,
    },
]

const router = new VueRouter({
    // mode: 'history',
    routes
})

export default router