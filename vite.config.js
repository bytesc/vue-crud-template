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


