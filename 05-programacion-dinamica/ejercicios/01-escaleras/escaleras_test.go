package escaleras

import "testing"

func TestFormasDeSubir(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1}, // 1
		{2, 2}, // 1+1, 2
		{3, 3}, // 1+1+1, 1+2, 2+1
		{4, 5},
		{5, 8},
		{6, 13},
	}
	for _, tt := range tests {
		got := FormasDeSubir(tt.n)
		if got != tt.want {
			t.Errorf("FormasDeSubir(%d) = %d, esperado %d", tt.n, got, tt.want)
		}
	}
}

func TestFormasDeSubirCero(t *testing.T) {
	// 0 escalones → 1 forma (no saltar)
	got := FormasDeSubir(0)
	if got != 1 {
		t.Errorf("FormasDeSubir(0) = %d, esperado 1", got)
	}
}
