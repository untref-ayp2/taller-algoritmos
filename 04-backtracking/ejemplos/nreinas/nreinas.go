package main

// Este archivo contiene el backtracking del problema de las N reinas.
// La solución parcial es un []int: la posición i corresponde a la fila i y
// su valor es la columna donde va la reina de esa fila.
//
// Arriba está el núcleo del algoritmo; al final, el observador y los puntos
// de entrada.

// ---------------------------------------------------------------------------
// Backtracking
// ---------------------------------------------------------------------------

// esSolucion verifica si la solución parcial ya tiene una reina en cada fila.
func esSolucion(n int, solucionParcial []int) bool {
	return len(solucionParcial) == n
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

// registrar agrega la columna elegida a la solución parcial.
func registrar(solucionParcial []int, columna int) []int {
	return append(solucionParcial, columna)
}

// borrar quita la última reina registrada (vuelta atrás).
func borrar(solucionParcial []int) []int {
	return solucionParcial[:len(solucionParcial)-1]
}

// backtracking es la recursión general del esquema. Devuelve true si la
// búsqueda debe continuar y false si el observador pidió detenerla.
func backtracking(n int, fila int, solucionParcial []int, obs ObservadorReinas) bool {
	// esSolucion: ya hay una reina colocada en cada fila
	if esSolucion(n, solucionParcial) {
		if obs.Solucion != nil {
			return obs.Solucion(solucionParcial)
		}
		return false
	}

	// extender: probar cada columna de la fila actual
	for columna := 0; columna < n; columna++ {
		// esFactible: la reina no entra en conflicto con las ya colocadas
		factible := esFactible(fila, columna, solucionParcial)
		if obs.Intento != nil {
			obs.Intento(fila, columna, factible, solucionParcial)
		}

		if factible {
			// registrar: coloco la reina en (fila, columna)
			solucionParcial = registrar(solucionParcial, columna)

			// llamada recursiva: paso a la siguiente fila
			if !backtracking(n, fila+1, solucionParcial, obs) {
				return false
			}

			// borrar (vuelta atrás): quito la reina que puse en esta fila
			solucionParcial = borrar(solucionParcial)
		}
	}

	// ninguna columna funcionó en esta fila → vuelta atrás
	return true
}

// ---------------------------------------------------------------------------
// Observador y puntos de entrada
// ---------------------------------------------------------------------------

// ObservadorReinas recibe los eventos de la búsqueda de las N reinas.
// Cualquiera de los dos callbacks puede ser nil.
type ObservadorReinas struct {
	// Intento se invoca antes de probar cada (fila, columna). factible
	// indica si la posición no entra en conflicto con las reinas ya
	// colocadas en parcial.
	Intento func(fila, columna int, factible bool, parcial []int)
	// Solucion se invoca al completar una solución. Devolver true obliga a
	// seguir buscando (vuelta atrás); devolver false detiene la búsqueda.
	Solucion func(solucion []int) bool
}

// NReinas resuelve el problema de las N reinas usando backtracking.
// Recibe la cantidad de reinas y devuelve un slice con las columnas
// donde se ubica cada reina (el índice del slice es la fila).
func NReinas(n int) []int {
	return NReinasConPaso(n, ObservadorReinas{})
}

// NReinasConPaso resuelve el problema notificando cada evento a obs.
// Devuelve la primera solución. Si obs.Solucion es nil, la búsqueda se
// detiene en esa primera solución.
func NReinasConPaso(n int, obs ObservadorReinas) []int {
	// Capturo la primera solución sin interferir con la decisión del
	// observador de seguir o no.
	var primera []int
	original := obs.Solucion
	obs.Solucion = func(solucion []int) bool {
		if primera == nil {
			primera = append([]int(nil), solucion...)
		}
		if original != nil {
			return original(solucion)
		}
		return false
	}

	backtracking(n, 0, nil, obs)
	return primera
}

// NReinasTodas devuelve todas las soluciones del problema de las N reinas.
func NReinasTodas(n int) [][]int {
	var todas [][]int
	NReinasConPaso(n, ObservadorReinas{
		Solucion: func(solucion []int) bool {
			todas = append(todas, append([]int(nil), solucion...))
			return true // seguir buscando hasta agotarlas
		},
	})
	return todas
}
