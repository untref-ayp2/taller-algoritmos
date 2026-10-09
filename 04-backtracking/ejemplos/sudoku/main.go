package main

import (
	"os"
	"time"
)

func main() {
	tablero := [9][9]int{
		{4, 0, 0, 1, 9, 5, 0, 6, 8},
		{8, 0, 0, 0, 0, 0, 7, 0, 1},
		{9, 0, 1, 6, 0, 0, 0, 3, 0},
		{0, 0, 7, 0, 2, 6, 0, 0, 0},
		{5, 0, 0, 0, 0, 0, 0, 0, 3},
		{0, 1, 0, 8, 7, 0, 4, 0, 0},
		{0, 3, 0, 0, 0, 0, 8, 0, 5},
		{1, 0, 5, 0, 0, 0, 0, 0, 0},
		{7, 9, 0, 4, 0, 1, 0, 0, 0},
	}

	sudoku := NuevoSudoku(tablero)
	sudoku.Animar(os.Stdout, 30*time.Millisecond)
}
