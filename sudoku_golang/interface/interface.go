package Sudointerface

type Sudoku interface {
	CreateSudoku() [][]int
	DigHole() [][]int
}
