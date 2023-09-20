import axios from 'axios'
import { showFullScreenLoading , tryHideFullScreenLoading } from '@/utils/loading.js'

// 创建 axios 实例，将来对创建出来的实例，进行自定义配置
const instance = axios.create({
  baseURL: '',
  timeout: 5000
})

// 自定义配置 - 请求/响应 拦截器
// 添加请求拦截器
instance.interceptors.request.use(function (config) {
  // 在发送请求之前做些什么
  // 开启loading，禁止背景点击 (节流处理，防止多次无效触发)
  showFullScreenLoading()
  return config
}, function (error) {
  // 对请求错误做些什么
  return Promise.reject(error)
})

// 添加响应拦截器
instance.interceptors.response.use(function (response) {
  // 2xx 范围内的状态码都会触发该函数。
  // 对响应数据做点什么 (默认axios会多包装一层data，需要响应拦截器中处理一下)
  const res = response.data
  if (res.status !== 200) {
    tryHideFullScreenLoading()
    Vue.prototype.$message({
      message: res.message,
      type: 'warning'
    });
    // 抛出一个错误的promise
    return Promise.reject(res.message)
  } else {
    tryHideFullScreenLoading()
  }
  return res
}, function (error) {
  // 超出 2xx 范围的状态码都会触发该函数。
  // 对响应错误做点什么
  return Promise.reject(error)
})

// 导出配置好的实例
export default instance
