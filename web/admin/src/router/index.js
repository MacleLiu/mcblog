import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from '../views/Login'
import Admin from '../views/Admin'

// 子页面组件
import AdminIndex from '../components/admin/AdminIndex'
import AddArt from '../components/article/AddArt'
import ArtList from '../components/article/ArtList'
import CateList from '../components/category/CateList'
import UserList from '../components/user/UserList'
import ToolList from '../components/tool/ToolList'
import WishList from '../components/wish/WishList'

Vue.use(VueRouter)

const routes = [
    {
        path: '/login',
        component: Login,
    },
    {
        path: '/admin',
        component: Admin,
        children: [
            {path: 'index', component: AdminIndex},
            {path: 'addArt', component: AddArt},
            {path: 'addArt/:id', component: AddArt, props: true},
            {path: 'articles', component: ArtList},
            {path: 'categories', component: CateList},
            {path: 'users', component: UserList},
            {path: 'tools', component: ToolList},
            {path: 'wishs', component: WishList},
        ]
    }
]

const router = new VueRouter({
    routes
})

router.beforeEach((to, from, next) => {
    const token = window.sessionStorage.getItem('token')
    if (to.path === '/login') return next()
    if (!token) {
        next('/login')
    }else{
        next()
    }
})

export default router