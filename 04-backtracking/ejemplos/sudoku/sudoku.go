package main

// Este archivo contiene el backtracking del sudoku.
//
// Arriba está el núcleo del algoritmo (el tablero y la recursión); al final,
// los puntos de entrada.

// ---------------------------------------------------------------------------
// Backtracking
// ---------------------------------------------------------------------------

// Sudoku resuelve un tablero 9×9 con backtracking. Las celdas vacías se
// representan con 0. Esta capa es puramente algorítmica: no conoce nada
// de la terminal ni de cómo se dibuja el tablero.
type Sudoku struct {
	tablero    [9][9]int
	soluciones [][9][9]int
}

// esSolucion verifica si no queda ninguna celda vacía en el tablero.
func (s *Sudoku) esSolucion() bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if s.tablero[y][x] == 0 {
				return false
			}
		}
	}
	return true
}

// proximaVacia devuelve la primera celda vacía, es decir, la próxima a
// extender.
func (s *Sudoku) proximaVacia() (x, y int) {
	for y = 0; y < 9; y++ {
		for x = 0; x < 9; x++ {
			if s.tablero[y][x] == 0 {
				return x, y
			}
		}
	}
	return -1, -1
}

// esFactible verifica si el valor n puede ir en (x, y) sin repetirse en su
// fila, su columna ni su cuadrante de 3×3.
func (s *Sudoku) esFactible(x, y, n int) bool {
	for i := 0; i < 9; i++ {
		if s.tablero[y][i] == n || s.tablero[i][x] == n {
			return false
		}
	}
	y0 := (y / 3) * 3
	x0 := (x / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s.tablero[y0+i][x0+j] == n {
				return false
			}
		}
	}
	return true
}

// registrar coloca el valor n en la celda (x, y).
func (s *Sudoku) registrar(x, y, n int) {
	s.tablero[y][x] = n
}

// borrar quita el valor de la celda (x, y) (vuelta atrás).
func (s *Sudoku) borrar(x, y int) {
	s.tablero[y][x] = 0
}

// guardarSolucion guarda una copia del tablero completo.
func (s *Sudoku) guardarSolucion() {
	var copia [9][9]int
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			copia[y][x] = s.tablero[y][x]
		}
	}
	s.soluciones = append(s.soluciones, copia)
}

// backtracking es la recursión general del esquema. Si paso no es nil, se
// invoca cada vez que se asigna un valor; la capa de visualización usa eso
// para animar la búsqueda sin que la lógica dependa de ella. Si primera es
// true, se detiene en la primera solución.
func (s *Sudoku) backtracking(paso func(x, y, n int), primera bool) bool {
	// esSolucion: no quedan celdas vacías.
	if s.esSolucion() {
		s.guardarSolucion()
		return true
	}

	// extender: probar los valores posibles en la próxima celda vacía.
	x, y := s.proximaVacia()
	for n := 1; n <= 9; n++ {
		if !s.esFactible(x, y, n) {
			continue
		}

		// registrar: pruebo el valor n en la celda vacía.
		s.registrar(x, y, n)
		if paso != nil {
			paso(x, y, n)
		}

		// llamada recursiva: resuelvo el resto del tablero.
		if s.backtracking(paso, primera) && primera {
			return true
		}

		// borrar (vuelta atrás): deshago la asignación.
		s.borrar(x, y)
	}

	// ningún valor sirve: esta rama no tiene solución.
	return false
}

// ---------------------------------------------------------------------------
// Puntos de entrada
// ---------------------------------------------------------------------------

func NuevoSudoku(tablero [9][9]int) *Sudoku {
	return &Sudoku{tablero: tablero}
}

// Resolver busca todas las soluciones del tablero y las deja en soluciones.
// Devuelve true si existe al menos una. El tablero original no se modifica.
func (s *Sudoku) Resolver() bool {
	s.soluciones = nil
	return s.backtracking(nil, false)
}
