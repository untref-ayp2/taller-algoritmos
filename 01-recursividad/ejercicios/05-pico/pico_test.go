package pico

import "testing"

func TestEncontrarPico(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 2, 3, 1, 0, -2}, 2},
		{[]int{1, 2, 3, 4, 5, 3, 1}, 4},
		{[]int{1, 2, 1}, 1},
		{[]int{1, 3, 5, 7, 6, 4, 2}, 3},
		{[]int{0, 2, 4, 6, 8, 7, 5, 3, 1}, 4},
	}
	for _, tt := range tests {
		got := EncontrarPico(tt.arr)
		if got != tt.want {
			t.Errorf("EncontrarPico(%v) = %d; esperado %d", tt.arr, got, tt.want)
		}
	}
}
