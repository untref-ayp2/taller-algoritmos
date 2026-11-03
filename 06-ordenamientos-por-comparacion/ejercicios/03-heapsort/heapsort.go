package heapsort

func HeapSort[T any](s []T, less func(a, b T) bool) {
	// TODO: implementar HeapSort adaptando la versión de ../ejemplos/heapsort/
	// (que usa cmp) para que use less.
	// Debe ordenar el slice recibido sin crear uno nuevo.
}

func downHeap[T any](s []T, n, i int, less func(a, b T) bool) {
	// TODO: hundir el nodo i dentro de un heap de tamaño n
}
