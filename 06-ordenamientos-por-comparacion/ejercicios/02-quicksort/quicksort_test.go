package quicksort

import (
	"slices"
	"testing"

	"github.com/untref-ayp2/taller-algoritmos/06-ordenamientos-por-comparacion/ejemplos/personas"
)

func isSorted[T any](s []T, less func(a, b T) bool) bool {
	for i := 1; i < len(s); i++ {
		if less(s[i], s[i-1]) {
			return false
		}
	}
	return true
}

func TestQuickSortEmpty(t *testing.T) {
	s := []int{}
	QuickSort(s, func(a, b int) bool { return a < b })
}

func TestQuickSortSingle(t *testing.T) {
	s := []int{42}
	QuickSort(s, func(a, b int) bool { return a < b })
	if s[0] != 42 {
		t.Errorf("expected 42, got %d", s[0])
	}
}

func TestQuickSortAsc(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	QuickSort(s, func(a, b int) bool { return a < b })
	if !isSorted(s, func(a, b int) bool { return a < b }) {
		t.Errorf("not sorted ascending: %v", s)
	}
}

func TestQuickSortDesc(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	QuickSort(s, func(a, b int) bool { return a > b })
	if !isSorted(s, func(a, b int) bool { return a > b }) {
		t.Errorf("not sorted descending: %v", s)
	}
}

func TestQuickSortAlreadySorted(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	QuickSort(s, func(a, b int) bool { return a < b })
	if !isSorted(s, func(a, b int) bool { return a < b }) {
		t.Errorf("already sorted should remain sorted")
	}
}

func TestQuickSortReverseSorted(t *testing.T) {
	s := []int{5, 4, 3, 2, 1}
	QuickSort(s, func(a, b int) bool { return a < b })
	if !isSorted(s, func(a, b int) bool { return a < b }) {
		t.Errorf("reverse sorted should become sorted: %v", s)
	}
}

func TestQuickSortDuplicates(t *testing.T) {
	s := []int{7, 7, 7, 7, 7}
	QuickSort(s, func(a, b int) bool { return a < b })
	if !isSorted(s, func(a, b int) bool { return a < b }) {
		t.Errorf("duplicates: %v", s)
	}
}

func TestQuickSortStrings(t *testing.T) {
	s := []string{"zorro", "abeja", "casa", "bosque"}
	QuickSort(s, func(a, b string) bool { return a < b })
	if !isSorted(s, func(a, b string) bool { return a < b }) {
		t.Errorf("strings not sorted: %v", s)
	}
}

func TestQuickSortLarge(t *testing.T) {
	s := make([]int, 1000)
	for i := range s {
		s[i] = 1000 - i
	}
	QuickSort(s, func(a, b int) bool { return a < b })
	if !slices.IsSorted(s) {
		t.Errorf("large slice not sorted")
	}
}

func TestQuickSortPersonas(t *testing.T) {
	casos := []struct {
		nombre string
		less   func(a, b personas.Persona) bool
	}{
		{"edad asc", func(a, b personas.Persona) bool { return a.Edad < b.Edad }},
		{"edad desc", func(a, b personas.Persona) bool { return a.Edad > b.Edad }},
		{"dni asc", func(a, b personas.Persona) bool { return a.DNI < b.DNI }},
		{"dni desc", func(a, b personas.Persona) bool { return a.DNI > b.DNI }},
		{"apellido asc", func(a, b personas.Persona) bool { return a.Apellido < b.Apellido }},
		{"apellido desc", func(a, b personas.Persona) bool { return a.Apellido > b.Apellido }},
	}
	for _, c := range casos {
		s := personas.Ejemplo()
		QuickSort(s, c.less)
		if !isSorted(s, c.less) {
			t.Errorf("%s: no ordenado: %+v", c.nombre, s)
		}
	}
}
