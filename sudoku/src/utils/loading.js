import { Loading } from 'element-ui';

let loading;
function startLoading () {
  loading = Loading.service({
      lock: true,
      text: '拼命加载中……',
      background: 'rgba(0, 0, 0, 0.8)',
      spinner: 'el-icon-loading'
  })
}
// 关闭loading事件
function endLoading () {
    loading.close()
}

// 默认状态
let needLoadingRequestCount = 0

// 开启loading
export function showFullScreenLoading() {
  if (needLoadingRequestCount === 0) {
      startLoading()
  }
  needLoadingRequestCount++
}
// 取消loading
export function tryHideFullScreenLoading() {
  if (needLoadingRequestCount <= 0) return
      needLoadingRequestCount--
  if (needLoadingRequestCount === 0) {
      endLoading()
  }
}
