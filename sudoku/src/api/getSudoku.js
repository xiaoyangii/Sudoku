import request from '@/utils/request.js'

export const getSudoku = (level) => {
  return request.post('', {
    params: {
      level
    }
  })
}