// Package heapsort implementa el algoritmo de ordenamiento por montículo
// (HeapSort) de forma genérica, usando una función de comparación cmp.
package heapsort

// HeapSort ordena el slice s según el criterio cmp.
//
// cmp(a, b) devuelve un entero negativo si a va antes que b, positivo si a va
// después que b, y cero si son equivalentes (es el mismo contrato que
// cmp.Compare de la biblioteca estándar).
//
// A diferencia de MergeSort y QuickSort, no usa división y conquista: se apoya
// en un max-heap (montículo de máximos) representado implícitamente sobre el
// propio slice. Ordena sin crear estructuras auxiliares (in-place), no es
// estable y cuesta O(n log n) en todos los casos.
func HeapSort[T any](s []T, cmp func(a, b T) int) {
	n := len(s)

	// Fase 1: construir el heap de máximos.
	// En un árbol binario completo, los nodos con índice >= n/2 son hojas (ya
	// son heaps de un solo elemento). Se hunden de abajo hacia arriba los nodos
	// internos, desde el último nodo no hoja (n/2-1) hasta la raíz (0). Al
	// terminar, s[0] es el máximo. Esta fase cuesta O(n).
	for i := n/2 - 1; i >= 0; i-- {
		downHeap(s, n, i, cmp)
	}

	// Fase 2: extraer el máximo repetidamente.
	// Se intercambia la raíz (máximo) con la última posición del heap, se
	// reduce el tamaño del heap en uno (el máximo queda en su posición final
	// ordenada) y se hunde la nueva raíz para restaurar la propiedad de heap.
	for i := n - 1; i > 0; i-- {
		s[0], s[i] = s[i], s[0] // mover el máximo al final del heap
		downHeap(s, i, 0, cmp)  // restaurar el heap sobre el prefijo s[:i]
	}
}

// downHeap hunde el nodo i dentro de un heap de tamaño n.
//
// Precondición: los subárboles de i cumplen la propiedad de heap.
// Postcondición: el subárbol con raíz en i cumple la propiedad de heap.
//
// Compara el nodo con sus hijos (en las posiciones 2*i+1 y 2*i+2) y, si alguno
// es "mayor" que él según cmp, lo intercambia con el mayor de ellos y repite
// hacia abajo hasta que el nodo quede en su posición correcta. Para ordenar de
// menor a mayor se construye un max-heap; el comparador genérico cmp permite,
// sin cambiar el código, construir un min-heap (y así ordenar de mayor a menor)
// simplemente invirtiendo la comparación.
func downHeap[T any](s []T, n, i int, cmp func(a, b T) int) {
	for {
		mayor := i
		izq := 2*i + 1
		der := 2*i + 2

		// Si el hijo izquierdo existe y es mayor que el candidato actual...
		if izq < n && cmp(s[mayor], s[izq]) < 0 {
			mayor = izq
		}
		// Si el hijo derecho existe y es mayor que el candidato actual...
		if der < n && cmp(s[mayor], s[der]) < 0 {
			mayor = der
		}

		// Si el nodo ya es el mayor de los tres, la propiedad de heap se cumple.
		if mayor == i {
			break
		}

		// Si no, intercambiar con el hijo mayor y seguir hundiendo.
		s[i], s[mayor] = s[mayor], s[i]
		i = mayor
	}
}
