package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// Este archivo es la capa de visualización del ejemplo. Dibuja el laberinto
// y anima la búsqueda. No contiene lógica de backtracking: sólo consume el
// callback que expone ResolverConPaso.

const (
	limpiarPantalla = "\033[2J\033[H"
	simboloPared    = "#"
	simboloLibre    = "."
	simboloCamino   = "o"
	simboloAgotado  = "x"
	simboloActual   = "@"
)

// Tablero dibuja el laberinto. Las celdas del camino actual se ven como 'o',
// las que ya se exploraron y se abandonaron como 'x', las paredes como '#'
// y el cursor (actualX, actualY) resaltado con '@'. Con actualX < 0 no se
// dibuja cursor.
func (r *RatMaze) Tablero(actualX, actualY int) string {
	var b strings.Builder
	for y := range r.maze {
		for x := range r.maze[y] {
			switch {
			case x == actualX && y == actualY:
				b.WriteString("\033[36m" + simboloActual + "\033[0m ")
			case r.solucion != nil && r.solucion[y][x]:
				b.WriteString("\033[32m" + simboloCamino + "\033[0m ")
			case r.visitado != nil && r.visitado[y][x]:
				b.WriteString("\033[90m" + simboloAgotado + "\033[0m ")
			case r.maze[y][x]:
				b.WriteString(simboloLibre + " ")
			default:
				b.WriteString(simboloPared + " ")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

// Animar resuelve el laberinto mostrando, paso a paso, el avance y las
// vueltas atrás hasta encontrar el camino.
func (r *RatMaze) Animar(w io.Writer, demora time.Duration) {
	pasos := 0
	encontrado := r.ResolverConPaso(func(x, y int, avanzando bool) {
		pasos++
		fmt.Fprint(w, limpiarPantalla)
		fmt.Fprint(w, r.Tablero(x, y))
		fmt.Fprintf(w, "\nPaso %d\n", pasos)
		time.Sleep(demora)
	})

	fmt.Fprint(w, limpiarPantalla)
	if encontrado {
		fmt.Fprintln(w, "Laberinto resuelto:")
	} else {
		fmt.Fprintln(w, "No existe un camino hasta el destino:")
	}
	fmt.Fprint(w, r.Tablero(-1, -1))
	fmt.Fprintf(w, "\nPasos: %d\n", pasos)
}

// Dibujar escribe el tablero sin cursor en w.
func (r *RatMaze) Dibujar(w io.Writer) {
	fmt.Fprint(w, r.Tablero(-1, -1))
}
