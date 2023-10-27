import { createRouter,createWebHistory,createWebHashHistory } from 'vue-router'

// 注册页面路由
import File1 from "../components/file/File1.vue";
import File2 from "../components/file/File2.vue";
import App from "../App.vue";
import Form from "../components/form/form.vue"

// 注册页面路由
const routes = [
    //这个是主路由
    {
        path:"/",
        component: App,
        redirect:"/form",  // 重定向，相当于默认打开form子路由，否则会空白
    },
    //以下三个是子路由
    {
        path: '/form',
        component:Form,
    },
    {
        path: '/file1',
        component:File1,
    },
    {
        path: '/file2',
        component:File2
    },
]

const router = createRouter({
    // history: createWebHistory(),
    history:createWebHashHistory(),
    routes,
})

export default router;

