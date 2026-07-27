package raiz

import "testing"

func TestRaizDigital(t *testing.T) {
	tests := []struct {
		n, esperado int
	}{
		{987, 6},
		{0, 0},
		{9, 9},
		{12345, 6}, // 1+2+3+4+5=15 → 1+5=6
		{9999, 9},  // 9+9+9+9=36 → 3+6=9
	}
	for _, tt := range tests {
		got := RaizDigital(tt.n)
		if got != tt.esperado {
			t.Errorf("RaizDigital(%d) = %d; esperado %d", tt.n, got, tt.esperado)
		}
	}
}
