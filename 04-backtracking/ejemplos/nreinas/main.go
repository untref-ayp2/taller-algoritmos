package main

import "fmt"

func main() {
	sol := NReinas(8)
	fmt.Println("Solución para 8 reinas:", sol)
}

// NReinas resuelve el problema de las N reinas usando backtracking.
// Recibe la cantidad de reinas y devuelve un slice con las columnas
// donde se ubica cada reina (el índice del slice es la fila).
func NReinas(n int) []int {
	var solucionParcial []int
	var solucion []int
	backtracking(n, 0, solucionParcial, &solucion)
	return solucion
}

// esFactible verifica si se puede colocar una reina en (fila, columna)
// dada una configuración parcial de reinas ya colocadas.
func esFactible(fila int, columna int, solucionParcial []int) bool {
	for i := range solucionParcial {
		if solucionParcial[i] == columna ||
			fila+columna == i+solucionParcial[i] ||
			fila-columna == i-solucionParcial[i] {
			return false
		}
	}
	return true
}

// backtracking es la función recursiva que explora las posiciones
// válidas para las reinas.
//
// Recibe la cantidad total de reinas, la fila actual, la configuración
// parcial y un puntero donde almacenar la solución completa.
func backtracking(n int, fila int, solucionParcial []int, solucion *[]int) {
	if fila == n {
		*solucion = make([]int, len(solucionParcial))
		copy(*solucion, solucionParcial)
		return
	}

	for columna := range n {
		if esFactible(fila, columna, solucionParcial) {
			nuevaSolucion := append(solucionParcial, columna)
			backtracking(n, fila+1, nuevaSolucion, solucion)
		}
	}
}
