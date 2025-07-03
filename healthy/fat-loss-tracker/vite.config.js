import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'
// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  server: {
    proxy: {
      '/api': {
        target: 'http://192.168.9.191:5011',
        // target: 'http://192.168.31.21:5011',
        changeOrigin: true,
        secure: false
      }
    }
  }
})
