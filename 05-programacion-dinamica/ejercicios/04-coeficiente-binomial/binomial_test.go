package binomial

import "testing"

func TestBinomial(t *testing.T) {
	tests := []struct {
		n, k, esperado int
	}{
		{5, 2, 10},
		{10, 0, 1},
		{10, 10, 1},
		{6, 3, 20},
		{30, 15, 155117520},
	}
	for _, tt := range tests {
		got := CoeficienteBinomial(tt.n, tt.k)
		if got != tt.esperado {
			t.Errorf("C(%d, %d) = %d; esperado %d", tt.n, tt.k, got, tt.esperado)
		}
	}
}
