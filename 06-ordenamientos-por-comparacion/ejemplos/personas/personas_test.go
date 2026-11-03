package personas

import (
	"slices"
	"testing"

	"github.com/untref-ayp2/taller-algoritmos/06-ordenamientos-por-comparacion/ejemplos/heapsort"
	"github.com/untref-ayp2/taller-algoritmos/06-ordenamientos-por-comparacion/ejemplos/mergesort"
	"github.com/untref-ayp2/taller-algoritmos/06-ordenamientos-por-comparacion/ejemplos/quicksort"
)

// probarCriterio verifica que los tres algoritmos dejen ordenado el slice de
// personas según el comparador cmp recibido.
func probarCriterio(t *testing.T, cmp func(a, b Persona) int) {
	t.Helper()

	algoritmos := []struct {
		nombre  string
		ordenar func([]Persona, func(a, b Persona) int)
	}{
		{"mergesort", mergesort.MergeSort[Persona]},
		{"quicksort", quicksort.QuickSort[Persona]},
		{"heapsort", heapsort.HeapSort[Persona]},
	}

	for _, alg := range algoritmos {
		s := Ejemplo()
		alg.ordenar(s, cmp)
		if !slices.IsSortedFunc(s, cmp) {
			t.Errorf("%s no ordenó correctamente: %+v", alg.nombre, s)
		}
	}
}

func TestOrdenarPorEdad(t *testing.T)         { probarCriterio(t, PorEdad) }
func TestOrdenarPorEdadDesc(t *testing.T)     { probarCriterio(t, PorEdadDesc) }
func TestOrdenarPorDNI(t *testing.T)          { probarCriterio(t, PorDNI) }
func TestOrdenarPorDNIDesc(t *testing.T)      { probarCriterio(t, PorDNIDesc) }
func TestOrdenarPorApellido(t *testing.T)     { probarCriterio(t, PorApellido) }
func TestOrdenarPorApellidoDesc(t *testing.T) { probarCriterio(t, PorApellidoDesc) }
