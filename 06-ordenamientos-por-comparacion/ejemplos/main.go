// Ejemplos de uso de los ordenamientos por comparación sobre un slice de
// Persona. Se ordena por edad, DNI y apellido, tanto de menor a mayor como de
// mayor a menor. Para correrlo:
//
//	go run ./06-ordenamientos-por-comparacion/ejemplos/
package main

import (
	"fmt"
	"slices"

	"github.com/untref-ayp2/taller-algoritmos/06-ordenamientos-por-comparacion/ejemplos/heapsort"
	"github.com/untref-ayp2/taller-algoritmos/06-ordenamientos-por-comparacion/ejemplos/mergesort"
	"github.com/untref-ayp2/taller-algoritmos/06-ordenamientos-por-comparacion/ejemplos/personas"
	"github.com/untref-ayp2/taller-algoritmos/06-ordenamientos-por-comparacion/ejemplos/quicksort"
)

func main() {
	originales := personas.Ejemplo()

	mostrar("Antes", originales)

	// El criterio de orden se separa del algoritmo: MergeSort ordena por edad,
	// DNI o apellido según el comparador cmp que reciba.
	ordenarYMostrar("Por edad (ascendente)", originales, personas.PorEdad)
	ordenarYMostrar("Por edad (descendente)", originales, personas.PorEdadDesc)
	ordenarYMostrar("Por DNI (ascendente)", originales, personas.PorDNI)
	ordenarYMostrar("Por DNI (descendente)", originales, personas.PorDNIDesc)
	ordenarYMostrar("Por apellido (ascendente)", originales, personas.PorApellido)
	ordenarYMostrar("Por apellido (descendente)", originales, personas.PorApellidoDesc)

	// Los tres algoritmos dejan el slice ordenado por edad. El orden entre
	// personas de igual edad puede variar: heapsort no es estable (se nota en
	// Bruno y Diego, que tienen la misma edad).
	fmt.Println("Los tres algoritmos ordenan por edad (asc):")
	mostrar("mergesort", ordenarConMergeSort(originales, personas.PorEdad))
	mostrar("quicksort", ordenarConQuickSort(originales, personas.PorEdad))
	mostrar("heapsort", ordenarConHeapSort(originales, personas.PorEdad))
}

// ordenarYMostrar ordena una copia del slice con MergeSort y la imprime.
func ordenarYMostrar(rotulo string, originales []personas.Persona, cmp func(a, b personas.Persona) int) {
	mostrar(rotulo, ordenarConMergeSort(originales, cmp))
}

// ordenarConMergeSort devuelve una copia de ps ordenada por cmp con MergeSort.
func ordenarConMergeSort(ps []personas.Persona, cmp func(a, b personas.Persona) int) []personas.Persona {
	copia := slices.Clone(ps)
	mergesort.MergeSort(copia, cmp)
	return copia
}

// ordenarConQuickSort devuelve una copia de ps ordenada por cmp con QuickSort.
func ordenarConQuickSort(ps []personas.Persona, cmp func(a, b personas.Persona) int) []personas.Persona {
	copia := slices.Clone(ps)
	quicksort.QuickSort(copia, cmp)
	return copia
}

// ordenarConHeapSort devuelve una copia de ps ordenada por cmp con HeapSort.
func ordenarConHeapSort(ps []personas.Persona, cmp func(a, b personas.Persona) int) []personas.Persona {
	copia := slices.Clone(ps)
	heapsort.HeapSort(copia, cmp)
	return copia
}

// mostrar imprime el rótulo y la lista de personas.
func mostrar(rotulo string, ps []personas.Persona) {
	fmt.Printf("%s:\n", rotulo)
	for _, p := range ps {
		fmt.Printf("  %s, %s — edad %d, DNI %d\n", p.Apellido, p.Nombre, p.Edad, p.DNI)
	}
	fmt.Println()
}
