package bucket

import (
	"reflect"
	"sort"
	"testing"
)

func TestBucketSort(t *testing.T) {
	tests := [][]float64{
		{},
		{0.5},
		{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68},
		{0.1, 0.9, 0.2, 0.8, 0.3, 0.7, 0.4, 0.6, 0.5},
	}
	for _, input := range tests {
		got := BucketSort(input)
		expected := make([]float64, len(input))
		copy(expected, input)
		sort.Float64s(expected)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("BucketSort(%v) = %v, esperado %v", input, got, expected)
		}
	}
}
