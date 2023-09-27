import request from '@/utils/request.js'

export const getSudoku = (diffcult) => {
  return request.get('/sudo', {
    params: {
      diffcult
    }
  })
}
