// Package quicksort implementa el algoritmo de ordenamiento rápido (QuickSort)
// de forma genérica, usando una función de comparación cmp.
package quicksort

// QuickSort ordena el slice s según el criterio cmp.
//
// cmp(a, b) devuelve un entero negativo si a va antes que b, positivo si a va
// después que b, y cero si son equivalentes (es el mismo contrato que
// cmp.Compare de la biblioteca estándar).
//
// Es un algoritmo divide y vencerás: elige un pivote, particiona s en dos zonas
// (los menores a la izquierda y los mayores a la derecha) y ordena cada zona
// recursivamente. El pivote queda en su posición final. Ordena el slice recibido
// sin crear uno nuevo (in-place). No es estable.
func QuickSort[T any](s []T, cmp func(a, b T) int) {
	// Caso base: un slice de 0 o 1 elementos ya está ordenado.
	if len(s) < 2 {
		return
	}

	// Particionar: el pivote queda en su posición final y se devuelve su índice.
	p := partition(s, cmp)

	// Ordenar recursivamente las zonas a izquierda y derecha del pivote.
	QuickSort(s[:p], cmp)
	QuickSort(s[p+1:], cmp)
}

// partition elige un pivote, reorganiza s para que todos los elementos menores
// queden a su izquierda y los mayores a su derecha, y devuelve la posición final
// del pivote.
//
// Estrategia (partición de 2 vías):
//  1. Se elige como pivote el elemento del medio de s.
//  2. Se lo mueve al final para simplificar el recorrido.
//  3. Se usan dos índices: i avanza desde la izquierda buscando un elemento
//     mayor que el pivote, y j retrocede desde la derecha buscando uno menor.
//     Cuando ambos se detienen, se intercambian. Se repite hasta cruzarse.
//  4. Se coloca el pivote en su posición final y se devuelve su índice.
func partition[T any](s []T, cmp func(a, b T) int) int {
	lo, hi := 0, len(s)-1

	// 1. Elegir el elemento del medio como pivote.
	mid := (lo + hi) / 2
	pivote := s[mid]

	// 2. Mover el pivote al final para no pisarlo durante el recorrido.
	s[mid], s[hi] = s[hi], s[mid]

	i, j := lo, hi-1

	// 3. Recorrer con los dos índices hasta que se crucen.
	for i <= j {
		// Avanzar i mientras el elemento sea menor o igual que el pivote
		// (cmp(s[i], pivote) <= 0).
		for i <= j && cmp(s[i], pivote) <= 0 {
			i++
		}

		// Retroceder j mientras el elemento sea mayor o igual que el pivote
		// (cmp(s[j], pivote) >= 0).
		for i <= j && cmp(s[j], pivote) >= 0 {
			j--
		}

		// Si no se cruzaron, hay un elemento "grande" a la izquierda y uno
		// "chico" a la derecha: intercambiarlos.
		if i < j {
			s[i], s[j] = s[j], s[i]
		}
	}

	// 4. Colocar el pivote en su posición final (i) y devolverla.
	s[i], s[hi] = s[hi], s[i]
	return i
}
