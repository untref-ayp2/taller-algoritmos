package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// Este archivo es la capa de visualización del ejemplo. Se encarga de
// dibujar el tablero en la terminal y de animar la búsqueda. No contiene
// lógica de backtracking: sólo consume el callback que expone backtracking.

// limpiarPantalla borra la terminal y lleva el cursor al inicio (ANSI).
const limpiarPantalla = "\033[2J\033[H"

// Tablero devuelve una representación del tablero con bordes. Las celdas
// vacías se ven en gris y la celda actual (actualX, actualY) resaltada.
func (s *Sudoku) Tablero(actualX, actualY, pasos int) string {
	return tableroString(s.tablero, actualX, actualY, pasos)
}

// Dibujar escribe el tablero en w.
func (s *Sudoku) Dibujar(w io.Writer, actualX, actualY, pasos int) {
	fmt.Fprint(w, s.Tablero(actualX, actualY, pasos))
}

// Animar resuelve el tablero mostrando cada asignación en w, con una pausa
// de demora entre paso y paso, y termina mostrando la solución encontrada.
func (s *Sudoku) Animar(w io.Writer, demora time.Duration) {
	s.soluciones = nil
	pasos := 0
	encontrada := s.backtracking(func(x, y, n int) {
		pasos++
		fmt.Fprint(w, limpiarPantalla)
		s.Dibujar(w, x, y, pasos)
		time.Sleep(demora)
	}, true)

	if encontrada {
		fmt.Fprint(w, limpiarPantalla)
		fmt.Fprintln(w, "Sudoku resuelto:")
		s.Dibujar(w, -1, -1, pasos)
	} else {
		fmt.Fprint(w, limpiarPantalla)
		fmt.Fprintln(w, "El tablero no tiene solución.")
	}
}

func tableroString(board [9][9]int, actualX, actualY, pasos int) string {
	var b strings.Builder
	b.WriteString("┏━━━━━━━┳━━━━━━━┳━━━━━━━┓\n")
	for y := 0; y < 9; y++ {
		b.WriteString("┃ ")
		for x := 0; x < 9; x++ {
			if x == 3 || x == 6 {
				b.WriteString("┃ ")
			}
			switch {
			case board[y][x] == 0:
				b.WriteString("\033[90m0\033[0m ")
			case x == actualX && y == actualY:
				b.WriteString(fmt.Sprintf("\033[36m%d\033[0m ", board[y][x]))
			default:
				b.WriteString(fmt.Sprintf("%d ", board[y][x]))
			}
			if x == 8 {
				b.WriteString("┃")
			}
		}
		if y == 2 || y == 5 {
			b.WriteString("\n┣━━━━━━━╋━━━━━━━╋━━━━━━━┫\n")
		} else {
			b.WriteString("\n")
		}
	}
	b.WriteString("┗━━━━━━━┻━━━━━━━┻━━━━━━━┛\n")
	b.WriteString(fmt.Sprintf("\nPasos: %d\n", pasos))
	return b.String()
}
