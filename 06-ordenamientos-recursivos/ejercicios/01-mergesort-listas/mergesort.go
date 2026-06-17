package mergesort

type Nodo struct {
	Valor int
	next  *Nodo
}

type Lista struct {
	cabeza *Nodo
}

func (l *Lista) Agregar(valor int) {
	// TODO: agregar al final
}

func MergeSort(lista *Lista) {
	// TODO: ordenar la lista usando Merge Sort (sin arreglos auxiliares)
}
