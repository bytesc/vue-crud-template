环境
```txt
go version go1.20.5 windows/amd64
npm -v 9.5.0
@vue/cli 5.0.8
```

新建项目
```bash
vue create vue-crud
>vite
>js
```

安装依赖
```bash
npm install

npm run dev
npm install -D unplugin-vue-components unplugin-auto-import
npm install element-plus
npm install @element-plus/icons-vue

npm install vue-router@4
```

`main.js`中创建app，注册组件
`app.vue`根组件
`router/router.js`中注册`app.vue`的根路由`/`

按照项目逻辑创建`*****/*****.vue`的子组件，导入到根组件`app.vue`直接显示。或者在`router/router.js`导入子页面路由，通过`app.vue`中的`<router-view></router-view>`组件显示。
