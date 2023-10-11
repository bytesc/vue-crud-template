import { createRouter,createWebHistory,createWebHashHistory } from 'vue-router'
import File1 from "../components/file/File1.vue";
import File2 from "../components/file/File2.vue";
import App from "../App.vue";
import Form from "../components/form/form.vue"


const routes = [
    {
        path:"/",
        component: App,
        redirect:"/form",
    },
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

