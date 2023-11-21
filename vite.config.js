import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server:{
    port : 8009, //指定部署端口号
    proxy:{
      "/api":{  //代理解决跨域
        target:"http://bytesc.top:8008/"
      }
    },
  },
  base: "./" //打包相对路径
})


