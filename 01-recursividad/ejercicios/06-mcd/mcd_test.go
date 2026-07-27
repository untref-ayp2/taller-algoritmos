package mcd

import "testing"

func TestMCD(t *testing.T) {
	tests := []struct {
		a, b, esperado int
	}{
		{48, 18, 6},
		{0, 5, 5},
		{7, 0, 7},
		{17, 13, 1},
		{100, 10, 10},
	}
	for _, tt := range tests {
		got := MCD(tt.a, tt.b)
		if got != tt.esperado {
			t.Errorf("MCD(%d, %d) = %d; esperado %d", tt.a, tt.b, got, tt.esperado)
		}
	}
}
