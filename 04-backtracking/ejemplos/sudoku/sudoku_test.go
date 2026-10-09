package main

import (
	"strings"
	"testing"
)

func TestSudokuUnicaSolucion(t *testing.T) {
	tablero := [9][9]int{
		{0, 0, 0, 1, 0, 5, 0, 6, 8},
		{0, 0, 0, 0, 0, 0, 7, 0, 1},
		{9, 0, 1, 0, 0, 0, 0, 3, 0},
		{0, 0, 7, 0, 2, 6, 0, 0, 0},
		{5, 0, 0, 0, 0, 0, 0, 0, 3},
		{0, 0, 0, 8, 7, 0, 4, 0, 0},
		{0, 3, 0, 0, 0, 0, 8, 0, 5},
		{1, 0, 5, 0, 0, 0, 0, 0, 0},
		{7, 9, 0, 4, 0, 1, 0, 0, 0},
	}

	sudoku := NuevoSudoku(tablero)
	sudoku.Resolver()

	if len(sudoku.soluciones) != 1 {
		t.Fatalf("soluciones = %d, se esperaba 1", len(sudoku.soluciones))
	}
	verificarSoluciones(t, tablero, sudoku.soluciones)
}

func TestSudokuMultiplesSoluciones(t *testing.T) {
	tablero := [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 0, 0},
	}

	sudoku := NuevoSudoku(tablero)
	sudoku.Resolver()

	if len(sudoku.soluciones) != 2 {
		t.Fatalf("soluciones = %d, se esperaban 2", len(sudoku.soluciones))
	}
	verificarSoluciones(t, tablero, sudoku.soluciones)
}

func TestSudokuMuchasSoluciones(t *testing.T) {
	tablero := [9][9]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 0, 0, 0, 4, 5, 6},
		{2, 3, 1, 0, 0, 0, 0, 0, 0},
		{5, 6, 4, 0, 0, 0, 0, 0, 0},
		{8, 9, 7, 0, 0, 0, 0, 0, 0},
		{3, 1, 2, 0, 0, 0, 0, 0, 0},
		{6, 4, 5, 0, 0, 0, 0, 0, 0},
		{9, 7, 8, 0, 0, 0, 0, 0, 0},
	}

	sudoku := NuevoSudoku(tablero)
	sudoku.Resolver()

	if len(sudoku.soluciones) != 32604 {
		t.Fatalf("soluciones = %d, se esperaban 32604", len(sudoku.soluciones))
	}
}

func TestSudokuTablero(t *testing.T) {
	tablero := [9][9]int{{5, 0, 0, 0, 0, 0, 0, 0, 1}}

	sudoku := NuevoSudoku(tablero)
	tableroStr := sudoku.Tablero(1, 0, 3)

	if !strings.Contains(tableroStr, "5") || !strings.Contains(tableroStr, "1") {
		t.Errorf("el tablero debería mostrar los valores fijos:\n%s", tableroStr)
	}
	if !strings.Contains(tableroStr, "┏") || !strings.Contains(tableroStr, "┛") {
		t.Errorf("el tablero debería tener bordes dibujados:\n%s", tableroStr)
	}
	if !strings.Contains(tableroStr, "Pasos") {
		t.Errorf("el tablero debería mostrar el contador de pasos:\n%s", tableroStr)
	}
}

func TestSudokuAnimar(t *testing.T) {
	tablero := [9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 0, 7, 9},
	}

	sudoku := NuevoSudoku(tablero)

	var salida strings.Builder
	sudoku.Animar(&salida, 0)

	if !strings.Contains(salida.String(), "Pasos") {
		t.Errorf("la animación debería mostrar el contador de pasos")
	}
}

func TestSudokuAnimarSinSolucion(t *testing.T) {
	// La primera fila ya tiene 1..8 y la columna 8 (y su cuadrante) tienen
	// un 9: la celda (8, 0) no admite ningún valor, así que no hay solución.
	tablero := [9][9]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 9},
	}

	sudoku := NuevoSudoku(tablero)

	var salida strings.Builder
	sudoku.Animar(&salida, 0)

	if !strings.Contains(salida.String(), "no tiene solución") {
		t.Errorf("la animación debería avisar que el tablero no tiene solución:\n%s", salida.String())
	}
}

func verificarSoluciones(t *testing.T, original [9][9]int, soluciones [][9][9]int) {
	t.Helper()
	for i, sol := range soluciones {
		if !esSudokuValido(sol) {
			t.Errorf("la solución %d no respeta las reglas del sudoku", i+1)
		}
		if !respetaOriginal(original, sol) {
			t.Errorf("la solución %d no respeta las pistas originales", i+1)
		}
	}
}

// esSudokuValido comprueba que cada fila, columna y cuadrante contenga
// exactamente los dígitos del 1 al 9 una sola vez.
func esSudokuValido(board [9][9]int) bool {
	for i := 0; i < 9; i++ {
		var fila, columna, caja [9]int
		for j := 0; j < 9; j++ {
			fila[j] = board[i][j]
			columna[j] = board[j][i]
			caja[j] = board[(i/3)*3+j/3][(i%3)*3+j%3]
		}
		if !sinRepetidos(fila) || !sinRepetidos(columna) || !sinRepetidos(caja) {
			return false
		}
	}
	return true
}

func sinRepetidos(valores [9]int) bool {
	var vistos [10]bool
	for _, v := range valores {
		if v < 1 || v > 9 || vistos[v] {
			return false
		}
		vistos[v] = true
	}
	return true
}

// respetaOriginal comprueba que la solución conserve todas las pistas
// (celdas distintas de 0) del tablero original.
func respetaOriginal(original, solucion [9][9]int) bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if original[y][x] != 0 && original[y][x] != solucion[y][x] {
				return false
			}
		}
	}
	return true
}
