package persona

// Elegir uno de los algoritmos implementados en los ejercicios 1-3.
// Descomentar la importación del elegido:
//
//	import "github.com/untref-ayp2/taller-algoritmos/06-ordenamientos-por-comparacion/ejercicios/01-mergesort"
//	import "github.com/untref-ayp2/taller-algoritmos/06-ordenamientos-por-comparacion/ejercicios/02-quicksort"
//	import "github.com/untref-ayp2/taller-algoritmos/06-ordenamientos-por-comparacion/ejercicios/03-heapsort"

type Persona struct {
	Nombre string
	Edad   int
}

// SortPersonas ordena el slice de personas usando el algoritmo elegido
// y la función de comparación less.
func SortPersonas(personas []Persona, less func(a, b Persona) bool) {
	// TODO: llamar al algoritmo elegido (descomentar import + llamar a MergeSort, QuickSort o HeapSort)
}

// OrdenarPorEdadAsc ordena personas de menor a mayor edad.
func OrdenarPorEdadAsc(personas []Persona) {
	// TODO: usar SortPersonas con el comparador adecuado
}

// OrdenarPorNombreDesc ordena personas alfabéticamente en orden descendente.
func OrdenarPorNombreDesc(personas []Persona) {
	// TODO: usar SortPersonas con el comparador adecuado
}

// OrdenarPorEdadAscNombreDesc ordena por edad ascendente y, ante empate, por nombre descendente.
func OrdenarPorEdadAscNombreDesc(personas []Persona) {
	// TODO: usar SortPersonas con el comparador adecuado
}
