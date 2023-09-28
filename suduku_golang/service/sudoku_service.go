package service

import (
	"log"
	"math/rand"
	"sudu/utils"
)

var (
	diffMap = map[int]int{
		1: 60,
		2: 49,
		3: 35,
		4: 31,
		5: 27,
	}
)

type SudokuSerice struct {
	matrixs [][utils.MATRIX_SIZE][utils.MATRIX_SIZE]int
	matrix  [utils.MATRIX_SIZE][utils.MATRIX_SIZE]int
}

func NewSudokuService() *SudokuSerice {
	return new(SudokuSerice)
}

func GetSudoku(diffcult int) [][utils.MATRIX_SIZE][utils.MATRIX_SIZE]int {
	log.Println("------------------")
	s := NewSudokuService()
	s.CreateSudoku()
	s.DigHole(diffcult)
	// for _, matrix := range s.matrixs {
	// 	utils.Print(matrix)
	// }
	return s.matrixs
}

func (s *SudokuSerice) CreateSudoku() {
	for i := 0; i < utils.MATRIX_SIZE; i++ {

		s.matrix = [utils.MATRIX_SIZE][utils.MATRIX_SIZE]int{}
		s.fillDiagonalBlocks()

		s.fillReamining(0, 0)
		s.matrixs = append(s.matrixs, s.matrix)
	}

}

func (s *SudokuSerice) DigHole(diffcult int) {
	for idx, matrix := range s.matrixs {
		if diffcult < 4 {
			s.matrixs[idx] = digHoleNormal(matrix, diffMap[diffcult])
		} else {
			s.matrixs[idx] = digHoleHard(matrix, diffMap[diffcult])
		}
	}
}

func (s *SudokuSerice) fillReamining(row, col int) bool {
	if row == utils.MATRIX_SIZE-1 && col == utils.MATRIX_SIZE-1 {
		return true
	}
	if col == utils.MATRIX_SIZE {
		col = 0
		row++
	}
	if row < utils.SRN {
		if col == 0 {
			col += utils.SRN
		}
	} else if row >= utils.SRN && row < utils.MATRIX_SIZE-utils.SRN {
		if col == utils.SRN {
			col += utils.SRN
		}
	} else if row >= utils.MATRIX_SIZE-utils.SRN {
		if col == utils.MATRIX_SIZE-utils.SRN {
			row++
			col = 0
			if row == utils.MATRIX_SIZE {
				return true
			}
		}
	}
	for i := 1; i <= utils.MATRIX_SIZE; i++ {
		if utils.CheckSafe(s.matrix, row, col, i) {
			s.matrix[row][col] = i
			if s.fillReamining(row, col+1) {
				return true
			}
		}
		s.matrix[row][col] = 0
	}
	return false
}

func (s *SudokuSerice) fillDiagonalBlocks() {
	for i := 0; i < utils.MATRIX_SIZE; i += 3 {
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				random := 0
				for {
					random = getRandomNumber(utils.MATRIX_SIZE)
					if utils.CheckBlock(s.matrix, x+i, y+i, random) {
						break
					}
				}
				s.matrix[x+i][y+i] = random
			}
		}
	}
}

func getRandomNumber(max int) int {
	num := rand.Intn(max) + 1
	return num
}

func digHoleHard(matrix [9][9]int, lowbound int) [9][9]int {

	visited := [9][9]int{}
	totalGiven := 9 * 9
	flag := false
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			flag = true
			n := matrix[row][col]
			if visited[row][col] != 1 {
				visited[row][col] = 1
				for i := 1; i <= 9; i++ {
					if i != n {
						matrix[row][col] = 0
						if utils.CheckSafe(matrix, row, col, i) {
							matrix[row][col] = i
							// copyMatrixUtil.copy
							if utils.Solve(matrix, 0, 0) {
								matrix[row][col] = n
								flag = false
								break
							}
						}
					}
				}
				if flag {
					matrix[row][col] = 0
					totalGiven--
					if totalGiven <= lowbound {
						return matrix
					}
				}
			}

		}
	}
	return matrix
}

func digHoleNormal(matrix [9][9]int, lowbound int) [9][9]int {
	visited := [9][9]int{}
	totalGiven := 9 * 9
	flag := false
	for {
		if totalGiven < lowbound {
			break
		}
		flag = true
		row := getRandomRowCol(9)
		col := getRandomRowCol(9)
		n := matrix[row][col]
		if visited[row][col] != 1 {
			visited[row][col] = 1
			for i := 1; i <= 9; i++ {
				if i != n {
					matrix[row][col] = 0
					if utils.CheckSafe(matrix, row, col, i) {
						matrix[row][col] = i
						if utils.Solve(matrix, 0, 0) {
							matrix[row][col] = n
							flag = false
							break
						}
					}
				}
			}
			if flag {
				matrix[row][col] = 0
				totalGiven--
			}
		}
	}
	return matrix
}

func getRandomRowCol(max int) int {
	return rand.Intn(max)
}
