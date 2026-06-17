package quicksort

import "testing"

func TestQuickSortMedianaTres(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{3, 1, 4, 1, 5, 9}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		original := make([]int, len(tt.input))
		copy(original, tt.input)
		QuickSortMedianaTres(tt.input)
		for i := 1; i < len(tt.input); i++ {
			if tt.input[i-1] > tt.input[i] {
				t.Errorf("QuickSortFallido: entrada %v, resultado %v", original, tt.input)
				break
			}
		}
	}
}
