package counting

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	tests := []struct {
		input  []int
		maxVal int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{4, 2, 2, 8, 3, 3, 1}, 8},
		{[]int{0, 0, 0, 0}, 0},
		{[]int{5, 4, 3, 2, 1, 0}, 5},
	}
	for _, tt := range tests {
		got := CountingSort(tt.input, tt.maxVal)
		expected := make([]int, len(tt.input))
		copy(expected, tt.input)
		// ordenar expected con sort nativo para comparar
		for i := 1; i < len(expected); i++ {
			for j := i; j > 0 && expected[j-1] > expected[j]; j-- {
				expected[j-1], expected[j] = expected[j], expected[j-1]
			}
		}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("CountingSort(%v, %d) = %v, esperado %v", tt.input, tt.maxVal, got, expected)
		}
	}
}

func TestCountingSortEstable(t *testing.T) {
	// Verificar que el algoritmo sea estable si corresponde
	// (no es obligatorio para esta implementación)
}
