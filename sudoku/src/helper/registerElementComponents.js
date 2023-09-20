import Vue from 'vue'
import 'element-ui/lib/theme-chalk/index.css'
import 'element-ui/lib/theme-chalk/base.css'
import {
  Message,
  MessageBox,
  Loading
} from 'element-ui'


Vue.prototype.$message = Message
Vue.prototype.$messageBox = MessageBox
Vue.prototype.$loading = Loading.service
Vue.prototype.$confirm = MessageBox.confirm
