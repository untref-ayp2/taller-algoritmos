package mergesort

func MergeSort[T any](s []T, less func(a, b T) bool) {
	// TODO: implementar MergeSort
	// Debe ordenar el slice recibido sin crear uno nuevo.
}

func merge[T any](s []T, mid int, less func(a, b T) bool) {
	// TODO: fusionar s[0:mid] y s[mid:n] ordenadamente
}
