import { createRouter,createWebHistory,createWebHashHistory } from 'vue-router'

// 注册页面路由
import App from "../App.vue";

import File1 from "../components/pages/Page1.vue";
import File2 from "../components/pages/Page2.vue";

import Form from "../components/form/form.vue"

import Info1 from "../components/pages/page_info/Info1.vue";
import Detail1 from "../components/pages/page_info/detail/Detail1.vue";

import Login from "../components/user/Login.vue";
import Signup from "../components/user/Signup.vue";
import ChangePwd from "../components/user/ChangePwd.vue";

import FileUpload from "../components/files/FileUpload.vue";
import BigFileUpload from "../components/files/BigFileUpload.vue";
import FileList from "../components/files/FileList.vue";

// 注册页面路由
const routes = [
    //这个是主路由
    {
        path:"/",
        component: App,
        redirect:"/form",  // 重定向，相当于默认打开form子路由，否则会空白
    },
    {
        path:"/user",
        name:"user",
        children: [
            {
                path:"login",
                component: Login,
                name: "login",
                meta:{
                    title:"登录",
                }
            },
            {
                path:"signup",
                component: Signup,
                name: "signup",
                meta:{
                    title:"注册",
                }
            },
            {
                path:"change_pwd",
                component: ChangePwd,
                name: "change_pwd",
                meta:{
                    title:"修改密码",
                }
            },
        ],
    },
    {
        path:"/files",
        name: "files",
        component: FileList,
        meta:{
            title:"文件",
        },
    },
    {
        path: '/form',
        component:Form,
        name: "form",
        meta:{
            title:"表单",
        }
    },
    {
        path: '/page1',
        component:File1,
        name: "page1",
        meta:{
            title:"页面1",
        },
        children: [   // 在<router-view></router-view>显示
            { path: 'info1', component: Info1, name:"info1" , // 子路由path不要加‘ / ’
                meta:{
                    title:"信息一",
                },
                children:[{
                    path: 'detail1', component: Detail1, name:"detail1" , // 子路由path不要加‘ / ’
                    meta:{
                        title:"详情一",
                    },
                },
                ]
            },
        ],
    },
    {
        path: '/page2',
        component:File2,
        name: "page2",
        meta:{
            title:"页面2",
        },
    },
]

const router = createRouter({
    // history: createWebHistory(),
    history:createWebHashHistory(),
    routes,
})

export default router;

