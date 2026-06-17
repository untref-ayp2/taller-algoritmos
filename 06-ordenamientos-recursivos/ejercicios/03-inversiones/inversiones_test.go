package inversiones

import "testing"

func TestContarInversiones(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{4, 3, 2, 1}, 6}, // n*(n-1)/2 = 6
		{[]int{2, 3, 8, 6, 1}, 5},
		{[]int{5, 4, 3, 2, 1}, 10},
	}
	for _, tt := range tests {
		got := ContarInversiones(tt.arr)
		if got != tt.want {
			t.Errorf("ContarInversiones(%v) = %d, esperado %d", tt.arr, got, tt.want)
		}
	}
}
