package utils

import (
	"fmt"
	"math"
)

const (
	NUM         = 9
	MATRIX_SIZE = 9
)

var (
	SRN int = int(math.Sqrt(float64(MATRIX_SIZE)))
)

func CheckSafe(sudoku [9][9]int, row, col, n int) bool {
	check := [10]int{}
	for i := 0; i < NUM; i++ {
		check[sudoku[row][i]]++
		if sudoku[row][i] == n {
			return false
		}
		if check[sudoku[row][i]] > 1 && sudoku[row][i] != 0 {
			return false
		}
	}
	check = [10]int{}
	for i := 0; i < NUM; i++ {
		if sudoku[i][col] == n {
			return false
		}
		if check[sudoku[row][i]] > 1 && sudoku[row][i] != 0 {
			return false
		}
	}

	return CheckBlock(sudoku, row, col, n)
}

func CheckBlock(sudoku [9][9]int, row, col, n int) bool {
	check := [10]int{}
	startRow := row - (row % 3)
	startCol := col - (col % 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			if sudoku[i+startRow][j+startCol] == n {
				// log.Println("false 1")
				return false
			}
			if check[sudoku[row][i]] > 1 && sudoku[row][i] != 0 {
				// log.Println("false 2")
				return false
			}
		}

	}
	return true
}

func CreateSudoku() [][]int {
	matrix := make([][]int, MATRIX_SIZE)

	return matrix
}

func Solve(matrix [9][9]int, row, col int) bool {
	if row == NUM-1 && col == NUM {
		return true
	}
	if col == NUM {
		col = 0
		row++
	}

	if matrix[row][col] != 0 {
		return Solve(matrix, row, col+1)
	}
	for i := 1; i <= NUM; i++ {
		if CheckSafe(matrix, row, col, i) {
			matrix[row][col] = i
			// sudoku.clone() ???
			if Solve(matrix, row, col+1) {
				return true
			}
		}
		matrix[row][col] = 0
	}
	return false
}

func Print(matrix [9][9]int) {
	for i := 0; i < NUM; i++ {
		for j := 0; j < NUM; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println("")
	}
	fmt.Println("")
}
