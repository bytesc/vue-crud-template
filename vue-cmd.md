## 基本配置

### 环境
```txt
go version go1.20.5 windows/amd64
npm -v 9.5.0
@vue/cli 5.0.8
```

### 新建项目
```bash
vue create vue-crud
>vite
>js
```
注意，框架会自己创建一个.gitignore文件忽略一些文件夹

### 安装依赖
```bash
npm install

npm run dev
npm install -D unplugin-vue-components unplugin-auto-import
npm install element-plus
npm install @element-plus/icons-vue

npm install vue-router@4
# npm install vuex

npm install axios
```

### 运行项目
```bash
npm run dev
```

## 项目结构

### router 路由
`main.js`中创建app，注册组件
`app.vue`根组件
`router/router.js`中注册`app.vue`的根路由`/`

router.js
```js
import { createRouter,createWebHistory,createWebHashHistory } from 'vue-router'

// 注册页面路由
import File1 from "../components/file/File1.vue";
import File2 from "../components/file/File2.vue";
import App from "../App.vue";
import Form from "../components/form/form.vue"

import Info1 from "../components/file/file_info/Info1.vue";
import Detail1 from "../components/file/file_info/detail/Detail1.vue";

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
        name: "form",
        meta:{
            title:"表单",
        }
    },
    {
        path: '/file1',
        component:File1,
        name: "file1",
        meta:{
            title:"文件1",
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
        path: '/file2',
        component:File2,
        name: "file2",
        meta:{
            title:"文件2",
        },
    },
]

const router = createRouter({
    // history: createWebHistory(),
    history:createWebHashHistory(),
    routes,
})

export default router;
```

### 组件结构
按照项目逻辑创建`*****/*****.vue`的子组件，`import`子组件导入到其它组件中直接显示（比如，nav，footer）。

或者在`router/router.js`导入子页面路由，通过`app.vue`中的`<router-view></router-view>`组件显示（比如各种增删改查页面）。

如果个别页面（比如login）不需要显示导航等子组件，可以用`<el-dialog>`设置 `:fullscreen="true"`



## Axios

新建 request.js
```js
//axios 封装

import axios from "axios";
import {ElMessage} from "element-plus";


// 全局配置
const service = axios.create({
    baseURL:"/api",
    timeout:300,  //请求超时
})

// 响应拦截
service.interceptors.response.use(res=>{
    // res为获取的所有数据
    // console.log(res)
    const {code,data,msg} = res.data
    // code,data,msg是后端返回json里的三个字段（后端随便约定）
    // code是后端随便约定的状态标识码（不是http状态码）
    if (code === "200"){
        // 后端返回成功
        ElMessage.success(msg)
        return data
    }else if(code === "400"){
        // 后端返回失败
        ElMessage.error(msg)
    }
})

// 请求方法配置
function request(options){
    if(options.method.toLowerCase() === "get"){
        options.params = options.data
    }
    return service(options)
}

// 给request添加request.post()和.get(),从而给request传参（这一步作用是封装）
["get","post"].forEach(item=>{
    request[item] = (url,data) =>{
        return request({
            url,
            data,
            method:item
        })
    }
})

// 导出才能用
export {request}
```

form.vue 中
```ts
import {request} from "../../utils/requests.js";
const getTableData = async (cur = 1)=>{
  // let res= request.get("/list",{
  //   pageSize:10,
  //   pageNum:cur
  // })
  let res= await request.get(`user/list/?pageSize=${10}&pageNum=${cur}`)
  console.log(res)
}
// getTableData()
```

解决跨域问题 vite.config.js
```js
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server:{
    port : 8086, //指定部署端口号
    proxy:{
      "/api":{  //代理解决跨域
        target:"http://127.0.0.1:8088/"
      }
    }
  },
  base: "./" //打包相对路径
})
```

## 打包
`vite.config.js` 配置打包相对路径(`./`)和指定部署端口号
```bash
npm run build
```
在`/dist/`文件夹下会有`index.html`.

如果本地打开需要用firefox，或是webstorm，vscode的Live Server插件等。crome内核的浏览器（google，edge）会出现无法访问本地文件的问题（不影响线上部署，线上部署后任何浏览器都可以打开的）。

一种解决方法：
```txt
对打包后的index.html进行处理

去除<script type=module>元素
除去其他<script>的nomodule属性
移除<script id=vite-legacy-entry>元素的内容，并将data-src属性名改为src
移除 <script id=vite-legacy-entry></script>内的代码
将所有资源地址修改为相对地址（例如/assets/index-legacy.xxxx.js改为./assets/index-legacy.xxxx.js注意，还有 CSS 文件）

```