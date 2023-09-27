import axios from 'axios'
import { showFullScreenLoading , tryHideFullScreenLoading } from '@/utils/loading.js'

// 创建 axios 实例，将来对创建出来的实例，进行自定义配置
const instance = axios.create({
  baseURL: 'http://43.136.122.18:8082',
  timeout: 5000,
  withCredentials: true,
  crossDomain: true,
  headers: {
    'Content-Type':'application/x-www-form-urlencoded'
    // 'Access-Control-Allow-Origin': '*',
  }
})

// 自定义配置 - 请求/响应 拦截器
// 添加请求拦截器
instance.interceptors.request.use(function (config) {
  // 开启loading，禁止背景点击 (节流处理，防止多次无效触发)
  showFullScreenLoading()
  return config
}, function (error) {
  // 对请求错误做些什么
  return Promise.reject(error)
})

// 添加响应拦截器
instance.interceptors.response.use(function (response) {
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
  tryHideFullScreenLoading()
  return Promise.reject(error)
})

// 导出配置好的实例
export default instance
