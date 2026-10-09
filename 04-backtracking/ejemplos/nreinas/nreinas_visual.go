package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// Este archivo es la capa de visualización del ejemplo. Dibuja el tablero,
// pregunta por la cantidad de reinas y anima la búsqueda. No contiene lógica
// de backtracking: sólo consume los eventos que expone NReinasConPaso.

// limpiarPantalla borra la terminal y lleva el cursor al inicio (ANSI).
const limpiarPantalla = "\033[2J\033[H"

// TableroNReinas dibuja un tablero n × n. Las reinas ya ubicadas se ven como
// ♛; la celda (actualFila, actualColumna) que se está intentando se resalta
// en verde si es factible y en rojo con × si está atacada. Con
// actualFila < 0 no se dibuja ningún cursor.
func TableroNReinas(n int, reinas []int, actualFila, actualColumna int, factible bool) string {
	var b strings.Builder
	for fila := 0; fila < n; fila++ {
		for columna := 0; columna < n; columna++ {
			switch {
			case fila == actualFila && columna == actualColumna && factible:
				b.WriteString("\033[32m♛\033[0m ") // verde: la reina puede ir acá
			case fila == actualFila && columna == actualColumna:
				b.WriteString("\033[31m×\033[0m ") // rojo: posición atacada
			case fila < len(reinas) && reinas[fila] == columna:
				b.WriteString("♛ ")
			default:
				b.WriteString("· ")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

// leerLinea lee una línea completa de entrada y devuelve su texto sin
// espacios al borde. El segundo valor es false cuando no hay más entrada.
func leerLinea(entrada *bufio.Reader) (string, bool) {
	linea, err := entrada.ReadString('\n')
	if linea == "" && err != nil {
		return "", false
	}
	return strings.TrimSpace(linea), true
}

// LeerCantidadReinas pregunta cuántas reinas se quieren. Acepta un valor
// entre 4 y 10; con Enter (o si se acaba la entrada) usa 6. Ante un valor
// inválido vuelve a preguntar.
func LeerCantidadReinas(entrada *bufio.Reader, salida io.Writer) int {
	const (
		minimo     = 4
		maximo     = 10
		porDefecto = 6
	)

	for {
		fmt.Fprintf(salida, "¿Cuántas reinas? (entre %d y %d) [%d]: ", minimo, maximo, porDefecto)
		linea, ok := leerLinea(entrada)
		if !ok {
			fmt.Fprintln(salida)
			return porDefecto
		}
		if linea == "" {
			return porDefecto
		}

		n, err := strconv.Atoi(linea)
		if err != nil || n < minimo || n > maximo {
			fmt.Fprintf(salida, "Valor inválido: ingresá un número entre %d y %d.\n", minimo, maximo)
			continue
		}
		return n
	}
}

// dibujarSoluciones escribe, una arriba de la otra, las soluciones ya
// encontradas.
func dibujarSoluciones(w io.Writer, n int, soluciones [][]int) {
	for i, solucion := range soluciones {
		fmt.Fprintf(w, "Solución %d. Tablero de %dx%d\n", i+1, n, n)
		fmt.Fprint(w, TableroNReinas(n, solucion, -1, -1, true))
		fmt.Fprintln(w)
	}
}

// AnimarReinas resuelve el problema de las N reinas mostrando la búsqueda
// paso a paso. Al encontrar una solución la deja apilada arriba y, debajo,
// anima la búsqueda de la próxima; tras cada solución pregunta si seguir.
// Las respuestas [s/n] se leen de entrada y el dibujo se escribe en salida.
func AnimarReinas(n int, entrada *bufio.Reader, salida io.Writer, demora time.Duration) {
	pasos := 0
	var soluciones [][]int
	detenida := false

	dibujarEnCurso := func(parcial []int, fila, columna int, factible bool) {
		fmt.Fprint(salida, limpiarPantalla)
		dibujarSoluciones(salida, n, soluciones)
		if len(soluciones) > 0 {
			fmt.Fprintf(salida, "=== Buscando la solución %d ===\n\n", len(soluciones)+1)
		}
		fmt.Fprint(salida, TableroNReinas(n, parcial, fila, columna, factible))
		fmt.Fprintf(salida, "\nIntentos: %d\n", pasos)
		time.Sleep(demora)
	}

	NReinasConPaso(n, ObservadorReinas{
		Intento: func(fila, columna int, factible bool, parcial []int) {
			pasos++
			dibujarEnCurso(parcial, fila, columna, factible)
		},
		Solucion: func(solucion []int) bool {
			soluciones = append(soluciones, append([]int(nil), solucion...))

			fmt.Fprint(salida, limpiarPantalla)
			dibujarSoluciones(salida, n, soluciones)
			fmt.Fprintf(salida, "\nIntentos: %d\n", pasos)
			fmt.Fprint(salida, "¿Buscar otra solución? [s/n]: ")

			linea, ok := leerLinea(entrada)
			if !ok {
				detenida = true
				return false
			}
			respuesta := strings.ToLower(linea)
			if respuesta != "s" && respuesta != "si" && respuesta != "sí" {
				detenida = true
				return false
			}
			return true
		},
	})

	fmt.Fprint(salida, limpiarPantalla)
	dibujarSoluciones(salida, n, soluciones)
	if detenida {
		fmt.Fprintf(salida, "Búsqueda detenida. Soluciones encontradas: %d.\n", len(soluciones))
	} else {
		fmt.Fprintf(salida, "No hay más soluciones. Total: %d.\n", len(soluciones))
	}
}
