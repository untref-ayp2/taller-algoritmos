package counting

import (
	"reflect"
	"testing"
)

type par struct {
	valor int
	orden int // orden de aparición original
}

func TestCountingSortStable(t *testing.T) {
	tests := []struct {
		input  []par
		maxVal int
	}{
		{
			[]par{{2, 1}, {1, 1}, {2, 2}, {1, 2}},
			2,
		},
		{
			[]par{{4, 1}, {2, 1}, {2, 2}, {8, 1}, {3, 1}, {3, 2}, {1, 1}},
			8,
		},
		{
			[]par{{0, 1}, {0, 2}, {0, 3}},
			0,
		},
		{
			[]par{{5, 1}, {3, 1}, {5, 2}, {3, 2}, {1, 1}},
			5,
		},
	}
	for _, tt := range tests {
		// Extraer solo los valores para ordenar
		vals := make([]int, len(tt.input))
		for i, p := range tt.input {
			vals[i] = p.valor
		}
		gotVals := CountingSortStable(vals, tt.maxVal)

		// Reconstruir el slice ordenado original para comparar estabilidad
		expected := make([]par, len(tt.input))
		copy(expected, tt.input)
		// Ordenar por valor, y ante empate por orden (estable)
		for i := 1; i < len(expected); i++ {
			for j := i; j > 0; j-- {
				if expected[j-1].valor > expected[j].valor ||
					(expected[j-1].valor == expected[j].valor && expected[j-1].orden > expected[j].orden) {
					expected[j-1], expected[j] = expected[j], expected[j-1]
				}
			}
		}

		gotPars := make([]par, len(gotVals))
		// Reconstruir pares desde gotVals buscando en el slice original
		disponibles := make([]par, len(tt.input))
		copy(disponibles, tt.input)

		for i, v := range gotVals {
			encontrado := -1
			for j, p := range disponibles {
				if p.valor == v {
					encontrado = j
					break
				}
			}
			gotPars[i] = disponibles[encontrado]
			disponibles = append(disponibles[:encontrado], disponibles[encontrado+1:]...)
		}

		if !reflect.DeepEqual(gotPars, expected) {
			t.Errorf("CountingSortStable inestable.\n  got:      %v\n  esperado: %v", gotPars, expected)
		}
	}
}

func TestCountingSortStableOrden(t *testing.T) {
	tests := [][]int{
		{},
		{1},
		{4, 2, 2, 8, 3, 3, 1},
		{0, 0, 0, 0},
		{5, 4, 3, 2, 1, 0},
	}
	for _, input := range tests {
		maxVal := 0
		for _, v := range input {
			if v > maxVal {
				maxVal = v
			}
		}
		got := CountingSortStable(input, maxVal)
		expected := make([]int, len(input))
		copy(expected, input)
		for i := 1; i < len(expected); i++ {
			for j := i; j > 0 && expected[j-1] > expected[j]; j-- {
				expected[j-1], expected[j] = expected[j], expected[j-1]
			}
		}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("CountingSortStable(%v, %d) = %v, esperado %v", input, maxVal, got, expected)
		}
	}
}
