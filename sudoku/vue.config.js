const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  publicPath: './',
  transpileDependencies: true,
  devServer: {
    proxy: {
      '/sudo': {
        target: 'http://43.136.122.18:8082',
        changeOrigin: true,
        pathRewrite: {
          '^/sudo': '/'
        }
      }
    }
  },
})
