package test

import (
	"math/rand"
	"sudu/service"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {

	for {
		op := rand.Intn(9) + 1
		t.Log(op)
		if op == 9 {
			break
		}
	}
	i := make([][]int, 9, 9)
	t.Log(len(i[0]))
	p := [9][9]int{}
	t.Log("\n", p)
	p[0][1] = 1
	t.Log("\n", p)
	p = [9][9]int{}
	t.Log(p)
}

func TestParam(t *testing.T) {
	ui := [2][2]int{}
	fortest(&ui)
	t.Log(ui)
}

func fortest(ma *[2][2]int) {
	ma[0][1] = 1
}

func TestSudoku(t *testing.T) {

	go service.GetSudoku(3)

	select {
	case <-time.After(8 * time.Second):
	}

}
