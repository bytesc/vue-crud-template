import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server:{
    port : 8080 //指定部署端口号
  },
  base: "./" //打包相对路径
})
