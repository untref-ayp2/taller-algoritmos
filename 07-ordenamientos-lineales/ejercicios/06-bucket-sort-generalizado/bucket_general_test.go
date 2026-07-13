package bucket

import (
	"math"
	"reflect"
	"sort"
	"testing"
)

func TestBucketSortGeneral(t *testing.T) {
	tests := []struct {
		input []float64
		k     int
	}{
		{[]float64{}, 5},
		{[]float64{0.5}, 5},
		{[]float64{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68}, 5},
		{[]float64{0.1, 0.9, 0.2, 0.8, 0.3, 0.7, 0.4, 0.6, 0.5}, 10},
		{[]float64{100.5, 200.3, 150.7, 175.2, 125.9}, 5},
		{[]float64{-5.0, -2.5, 0.0, 2.5, 5.0}, 4},
		{[]float64{3.14, 2.71, 1.41, 1.73, 2.23}, 3},
	}
	for _, tt := range tests {
		got := BucketSort(tt.input, tt.k)
		expected := make([]float64, len(tt.input))
		copy(expected, tt.input)
		sort.Float64s(expected)

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("BucketSort(%v, %d) = %v, esperado %v", tt.input, tt.k, got, expected)
		}
	}
}

func TestBucketSortGeneralDefaultK(t *testing.T) {
	input := []float64{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68}
	got := BucketSort(input, 0)
	expected := make([]float64, len(input))
	copy(expected, input)
	sort.Float64s(expected)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("BucketSort con k=0 falló: %v, esperado %v", got, expected)
	}
}

// Para floats en [0, 1), BucketSort(arr, 10) debe dar el mismo resultado
// que la versión básica.
func TestBucketSortEquivalenteABasico(t *testing.T) {
	input := []float64{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68}
	got := BucketSort(input, 10)
	sort.Float64s(input)
	if !reflect.DeepEqual(got, input) {
		t.Errorf("BucketSort con 10 buckets no ordenó correctamente: %v", got)
	}
}

func TestBucketSortDistribucionUniforme(t *testing.T) {
	// Generar datos con distribución uniforme en [10, 20)
	n := 100
	input := make([]float64, n)
	for i := range input {
		input[i] = 10.0 + float64(i)*0.1 // 10.0, 10.1, 10.2, ..., 19.9
	}
	// Desordenar
	for i := n - 1; i > 0; i-- {
		j := int(math.Abs(float64(i*i*31+7))) % (i + 1)
		input[i], input[j] = input[j], input[i]
	}

	got := BucketSort(input, 10)
	expected := make([]float64, n)
	copy(expected, input)
	sort.Float64s(expected)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("BucketSort falló con datos uniformes en [10, 20)")
	}
}
