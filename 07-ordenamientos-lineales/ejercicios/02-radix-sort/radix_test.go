package radix

import (
	"reflect"
	"sort"
	"testing"
)

func TestRadixSort(t *testing.T) {
	tests := [][]int{
		{},
		{1},
		{170, 45, 75, 90, 2, 24, 802, 66},
		{3, 1, 4, 1, 5, 9, 2, 6, 5},
		{1000, 999, 888, 777, 666},
	}
	for _, input := range tests {
		got := RadixSort(input)
		expected := make([]int, len(input))
		copy(expected, input)
		sort.Ints(expected)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("RadixSort(%v) = %v, esperado %v", input, got, expected)
		}
	}
}
