// import './style.css'

import { createApp } from 'vue'

//创建vue app
import App from './App.vue'
const app = createApp(App)

// 注册ElementPlus前端组件
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.use(ElementPlus)

// 把路由配置注册到应用
import router from './router/router'
app.use(router)

// 把vuex册到应用
import {store} from "./utils/store.js";
app.use(store)

app.mount('#app')





