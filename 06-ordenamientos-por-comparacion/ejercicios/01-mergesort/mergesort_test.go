package mergesort

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

func TestMergeSortEmpty(t *testing.T) {
	s := []int{}
	MergeSort(s, func(a, b int) bool { return a < b })
}

func TestMergeSortSingle(t *testing.T) {
	s := []int{42}
	MergeSort(s, func(a, b int) bool { return a < b })
	if s[0] != 42 {
		t.Errorf("expected 42, got %d", s[0])
	}
}

func TestMergeSortAsc(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	MergeSort(s, func(a, b int) bool { return a < b })
	if !isSorted(s, func(a, b int) bool { return a < b }) {
		t.Errorf("not sorted ascending: %v", s)
	}
}

func TestMergeSortDesc(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	MergeSort(s, func(a, b int) bool { return a > b })
	if !isSorted(s, func(a, b int) bool { return a > b }) {
		t.Errorf("not sorted descending: %v", s)
	}
}

func TestMergeSortAlreadySorted(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	MergeSort(s, func(a, b int) bool { return a < b })
	if !isSorted(s, func(a, b int) bool { return a < b }) {
		t.Errorf("already sorted slice should remain sorted")
	}
}

func TestMergeSortReverseSorted(t *testing.T) {
	s := []int{5, 4, 3, 2, 1}
	MergeSort(s, func(a, b int) bool { return a < b })
	if !isSorted(s, func(a, b int) bool { return a < b }) {
		t.Errorf("reverse sorted should become sorted: %v", s)
	}
}

func TestMergeSortDuplicates(t *testing.T) {
	s := []int{7, 7, 7, 7, 7}
	MergeSort(s, func(a, b int) bool { return a < b })
	if !isSorted(s, func(a, b int) bool { return a < b }) {
		t.Errorf("duplicates: %v", s)
	}
}

func TestMergeSortStrings(t *testing.T) {
	s := []string{"zorro", "abeja", "casa", "bosque"}
	MergeSort(s, func(a, b string) bool { return a < b })
	if !isSorted(s, func(a, b string) bool { return a < b }) {
		t.Errorf("strings not sorted: %v", s)
	}
}

func TestMergeSortStable(t *testing.T) {
	type par struct {
		key int
		ord int
	}
	s := []par{
		{key: 2, ord: 1},
		{key: 1, ord: 2},
		{key: 2, ord: 3},
		{key: 1, ord: 4},
	}
	MergeSort(s, func(a, b par) bool { return a.key < b.key })
	if s[0].ord != 2 || s[1].ord != 4 || s[2].ord != 1 || s[3].ord != 3 {
		t.Errorf("MergeSort no es estable: %+v", s)
	}
}

func TestMergeSortLarge(t *testing.T) {
	s := make([]int, 1000)
	for i := range s {
		s[i] = 1000 - i
	}
	MergeSort(s, func(a, b int) bool { return a < b })
	if !slices.IsSorted(s) {
		t.Errorf("large slice not sorted")
	}
}

func TestMergeSortPersonas(t *testing.T) {
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
		MergeSort(s, c.less)
		if !isSorted(s, c.less) {
			t.Errorf("%s: no ordenado: %+v", c.nombre, s)
		}
	}
}
