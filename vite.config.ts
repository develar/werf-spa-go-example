// eslint-disable-next-line @typescript-eslint/ban-ts-comment
import vue from "@vitejs/plugin-vue"
import { defineConfig } from "vite"

// https://vitejs.dev/config/
// noinspection SpellCheckingInspection,TypeScriptUnresolvedVariable
// @ts-ignore
export default defineConfig({
  plugins: [
    vue(),
  ],
  build: {
    sourcemap: true,
    emptyOutDir: true,
  },
  server: {
    proxy: {
      "/graphql": "http://localhost:8080",
    }
  }
})