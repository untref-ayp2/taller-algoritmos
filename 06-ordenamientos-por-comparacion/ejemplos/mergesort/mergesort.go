// Package mergesort implementa el algoritmo de ordenamiento por mezcla
// (MergeSort) de forma genérica, usando una función de comparación cmp.
package mergesort

// MergeSort ordena el slice s según el criterio cmp.
//
// cmp(a, b) devuelve un entero negativo si a va antes que b, positivo si a va
// después que b, y cero si son equivalentes (es el mismo contrato que
// cmp.Compare de la biblioteca estándar).
//
// El algoritmo divide s en dos mitades, ordena cada una recursivamente y las
// fusiona. Ordena el slice recibido sin crear uno nuevo: el resultado queda en
// el propio s. Es estable (mantiene el orden relativo de elementos iguales) y
// cuesta O(n log n) en todos los casos.
func MergeSort[T any](s []T, cmp func(a, b T) int) {
	// Caso base: un slice de 0 o 1 elementos ya está ordenado.
	if len(s) < 2 {
		return
	}

	// Dividir: partir el slice en dos mitades de tamaño aproximado.
	mid := len(s) / 2
	MergeSort(s[:mid], cmp) // ordenar la mitad izquierda
	MergeSort(s[mid:], cmp) // ordenar la mitad derecha

	// Combinar: fusionar las dos mitades ya ordenadas.
	merge(s, mid, cmp)
}

// merge fusiona las dos mitades ordenadas de s en un único slice ordenado.
//
// Precondición: s[:mid] y s[mid:] están ordenadas según cmp.
// Postcondición: s queda ordenado según cmp.
//
// Se copia cada mitad a un slice auxiliar para poder recorrerlas y volcarlas
// en orden sobre s sin sobrescribir valores todavía no procesados. Ante un
// empate (cmp devuelve 0) se elige siempre el elemento de la mitad izquierda,
// lo que hace al algoritmo estable.
func merge[T any](s []T, mid int, cmp func(a, b T) int) {
	n := len(s)

	// Copias auxiliares de cada mitad (esto es el O(n) de memoria extra).
	izq := make([]T, mid)
	der := make([]T, n-mid)
	copy(izq, s[:mid])
	copy(der, s[mid:])

	i, j, k := 0, 0, 0

	// Mezclar mientras ambas mitades tengan elementos pendientes.
	for i < len(izq) && j < len(der) {
		if cmp(der[j], izq[i]) < 0 {
			// El elemento de la derecha es menor: va primero.
			s[k] = der[j]
			j++
		} else {
			// El de la izquierda es menor o hay empate: va primero (estable).
			s[k] = izq[i]
			i++
		}
		k++
	}

	// Volcar los elementos restantes de la mitad izquierda, si quedaron.
	for i < len(izq) {
		s[k] = izq[i]
		i++
		k++
	}

	// Volcar los elementos restantes de la mitad derecha, si quedaron.
	for j < len(der) {
		s[k] = der[j]
		j++
		k++
	}
}
