package sudoku

import "testing"

func TestSudokuConSolucion(t *testing.T) {
	tablero := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	if !ResolverSudoku(tablero) {
		t.Fatal("esperado true para tablero con solución")
	}
	if !esValido(tablero) {
		t.Error("el tablero resuelto no es válido")
	}
}

func TestSudokuImposible(t *testing.T) {
	tablero := [][]int{
		{5, 5, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	if ResolverSudoku(tablero) {
		t.Error("esperado false para tablero con dos 5 en primera fila")
	}
}

func TestSudokuCompleto(t *testing.T) {
	tablero := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	if !ResolverSudoku(tablero) {
		t.Fatal("esperado true para tablero ya completo y válido")
	}
}

func esValido(tablero [][]int) bool {
	for i := 0; i < 9; i++ {
		fila := make(map[int]bool)
		col := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if tablero[i][j] != 0 {
				if fila[tablero[i][j]] {
					return false
				}
				fila[tablero[i][j]] = true
			}
			if tablero[j][i] != 0 {
				if col[tablero[j][i]] {
					return false
				}
				col[tablero[j][i]] = true
			}
		}
	}
	for bi := 0; bi < 9; bi += 3 {
		for bj := 0; bj < 9; bj += 3 {
			box := make(map[int]bool)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					v := tablero[bi+i][bj+j]
					if v != 0 {
						if box[v] {
							return false
						}
						box[v] = true
					}
				}
			}
		}
	}
	return true
}
