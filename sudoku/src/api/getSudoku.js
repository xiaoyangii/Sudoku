import request from '@/utils/request.js'

export const getSudoku = (difficulty) => {
  return request.get('/newSudoku/createTerminalMatrix', {
    params: {
      difficulty
    }
  })
}