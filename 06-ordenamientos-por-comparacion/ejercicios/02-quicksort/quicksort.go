package quicksort

func QuickSort[T any](s []T, less func(a, b T) bool) {
	// TODO: implementar QuickSort adaptando la versión de ../ejemplos/quicksort/
	// (que usa cmp) para que use less.
	// Debe ordenar el slice recibido sin crear uno nuevo.
}

func partition[T any](s []T, less func(a, b T) bool) int {
	// TODO: particionar eligiendo el elemento del medio como pivote (partición de 2 vías)
	return 0
}
