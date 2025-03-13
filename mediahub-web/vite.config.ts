import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vite.dev/config/
export default defineConfig({
  base:"/mediahub-web",
  plugins: [
    vue()
  ],
  server: {
    proxy: {
      "/mediahub-api": { // “/api” 以及前置字符串会被替换为真正域名 
        target: "http://mediahub/", // 请求域名 
        secure: false, // 请求是否为https 
        changeOrigin: true  // 是否跨域 
      }
    }
  }
})
