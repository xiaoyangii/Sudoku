import request from '@/utils/request.js'

export const getSudoku = (diffcult) => {
  return request.get('/sudo', {
    params: {
      diffcult,
      id: 0
    }
  })
}

export const solveSudoku = (diffcult) => {
  return request.get('/solve', {
    params: {
      id: 0
    }
  })
}
