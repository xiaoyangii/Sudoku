import Vue from 'vue'
import App from './App.vue'
import router from './router'
// 全局默认样式
import '@/assets/style/reset.css'
import '@/assets/style/global.css'
import '@/assets/style/variable.less'
import './utils/rem.js'
// 按需引入 element-ui 组件
import '@/helper/registerElementComponents.js'


Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
