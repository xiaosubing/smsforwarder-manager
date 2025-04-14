import {createRouter, createWebHashHistory } from "vue-router";
import { userInfoStore } from "../store/user"



const routers = [
    {
        path: "/",
        name: "login",
        meta: {"title": "后台管理系统"},
        component: () => import('../views/Login.vue')
        // component: login
    },
    {
        path: "/index",
        name: "index",
        meta: {requiresAuth: true},
        component: ()=> import('../views/Index.vue'),
    }
]


const router = createRouter({
    history: createWebHashHistory(),
    routes: routers
})

router.beforeEach((to, _from, next) => {
    const store = userInfoStore()
    let isLogin = store.isLogin

    if (to.name !== 'Login' && to.name !== 'login' && !isLogin) {
        next({name: 'login'})
    }else if (to.name === 'login' && isLogin) {
        next({name: 'index'})
    }else {
        next()
    }
})


export default router