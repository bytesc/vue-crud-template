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
```

### 运行项目
```bash
npm run dev
```

### 基本结构和配置
`main.js`中创建app，注册组件
`app.vue`根组件
`router/router.js`中注册`app.vue`的根路由`/`

按照项目逻辑创建`*****/*****.vue`的子组件，导入到根组件`app.vue`直接显示（比如，nav，footer）。
或者在`router/router.js`导入子页面路由，通过`app.vue`中的`<router-view></router-view>`组件显示（比如各种增删改查页面）。

如果个别页面（比如login）不需要显示导航等子组件，可以用`<el-dialog>`设置 `:fullscreen="true"`

### 打包
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
