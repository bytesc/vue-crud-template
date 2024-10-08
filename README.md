# vue-crud-template

✨基于和 vue3, axios 和 Go, gorm, gin 和 MySQL  的简单信息管理系统模板✨📌含完整前后端，项目在线demo：信息管理系统模板，后台管理系统模板，数据库管理系统模板。实现令牌签验，非对称加密，通过 Web 应用完成对数据库的增删改查(CRUD)，文件流的上传和下载。📌前后端分离

📌[在线演示链接](http://bytesc.top:8009)

📌[配套后端项目地址(分布式版)](https://github.com/bytesc/go-grpc-crud-template) 

📌[配套后端项目地址(普通版)](https://github.com/bytesc/go-crud-template)

[个人网站：www.bytesc.top](http://www.bytesc.top) 

🔔 如有项目相关问题，欢迎在本项目提出`issue`，我一般会在 24 小时内回复。

## 系统架构

![](./docs/readme_img/sys.png)

## 效果展示

### 文件流

文件流上传

![](./docs/readme_img/imgf1.png)

文件流下载

![](./docs/readme_img/imgf2.png)

文件列表

![](./docs/readme_img/imgfl.png)

### CRUD

![](./docs/readme_img/img1.png)

完善的查询

![](./docs/readme_img/img2.png)

多选删除

![](./docs/readme_img/img3.png)

编辑行

![](./docs/readme_img/img4.png)


### 用户登录

![](./docs/readme_img/imgu.png)

面包屑导航

![](./docs/readme_img/img7.png)
![](./docs/readme_img/img8.png)

## 项目运行方法

### 前端运行环境

- node.js `npm -v 9.5.0`
- vue3 `@vue/cli 5.0.8`


### 安装依赖
```bash
npm install

# npm run dev
# npm install -D unplugin-vue-components unplugin-auto-import

# npm install element-plus
# npm install @element-plus/icons-vue

# npm install --save nprogress

# npm install vue-router@4
# npm install vuex@next --save
# npm install axios

# npm install jsencrypt --save

```

### 运行项目
```bash
npm run dev
```

### 运行端口
`vite.config.js`
```js
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
- `port` 为前端的运行端口
- `target` 后端服务器 url

如果希望显示数据，需要用到配套的后端项目📌[配套后端项目地址](https://github.com/bytesc/go-crud-template)

### 项目打包

`vite.config.js` 配置打包相对路径(`./`)和指定部署端口号
```bash
npm run build
```
在`/dist/`文件夹下会有`index.html`.

如果本地打开需要用firefox，或是webstorm，vscode的Live Server插件等。crome内核的浏览器（google，edge）会出现无法访问本地文件的问题（不影响线上部署，线上部署后任何浏览器都可以打开的）。

### 生产环境的跨域代理

以 nigix 为例
```bash
sudo vim /etc/nginx/sites-enabled/default
```
修改配置文件，配置代理
```txt
server {
       listen 8009;
       listen [::]:8009;

       server_name bytesc.top www.bytesc.top;

       root /var/www/vue-crud;
       index index.html;

       location / {
               try_files $uri $uri/ =404;
       }
    location /api {
        proxy_pass http://bytesc.top:8008;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```
重启 nginx 服务
```bash
sudo systemctl restart nginx
```

### 配套后端项目

还需要运行配套后端项目 📌[配套后端项目地址(分布式版)](https://github.com/bytesc/go-grpc-crud-template) 或 📌[配套后端项目地址(普通版)](https://github.com/bytesc/go-crud-template)

### 官方文档

- [vue3](https://cn.vuejs.org/guide/quick-start.html)
- [vite](https://cn.vitejs.dev/guide/)

- [vue-router](https://router.vuejs.org/zh/)
- [axios](https://www.axios-http.cn/docs/intro)
- [vuex](https://vuex.vuejs.org/zh/guide/)

- [element-plus](https://element-plus.org/zh-CN/)



# 开源许可证

此翻译版本仅供参考，以 LICENSE 文件中的英文版本为准

MIT 开源许可证：

版权所有 (c) 2023 bytesc

特此授权，免费向任何获得本软件及相关文档文件（以下简称“软件”）副本的人提供使用、复制、修改、合并、出版、发行、再许可和/或销售软件的权利，但须遵守以下条件：

上述版权声明和本许可声明应包含在所有副本或实质性部分中。

本软件按“原样”提供，不作任何明示或暗示的保证，包括但不限于适销性、特定用途适用性和非侵权性。在任何情况下，作者或版权持有人均不对因使用本软件而产生的任何索赔、损害或其他责任负责，无论是在合同、侵权或其他方面。
