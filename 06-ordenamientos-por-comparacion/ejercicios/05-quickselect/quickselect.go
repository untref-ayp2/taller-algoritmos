package quickselect

func QuickSelect[T any](s []T, k int, less func(a, b T) bool) T {
	// TODO: devolver el k-ésimo elemento más pequeño
	var zero T
	return zero
}

func partition[T any](s []T, lo, hi int, less func(a, b T) bool) int {
	// TODO: particionar s[lo:hi+1] usando s[hi] como pivote
	return 0
}
