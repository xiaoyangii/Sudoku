import request from '@/utils/request.js'

export const getSudoku = (difficulty) => {
  return request.get('/sudo', {
    params: {
      difficulty
    }
  })
}
