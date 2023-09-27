const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    port: port,
    open: false,//是否自动打开浏览器
    overlay: {
      warnings: false,
      errors: false
    },
    proxy: {
      '/sudo': {
        target: 'http://43.136.122.18:8082/sudo',
        changeOrigin: true,
        pathRewrite: {
          '^/sudo': '/'
        }
      }
    }
  },
})
