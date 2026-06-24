package cambio

import "testing"

func TestCambioTab(t *testing.T) {
	tests := []struct {
		name     string
		monto    int
		monedas  []int
		esperado int
	}{
		{
			name:     "monto 0",
			monto:    0,
			monedas:  []int{1, 3, 4},
			esperado: 0,
		},
		{
			name:     "caso simple con monedas canónicas",
			monto:    8,
			monedas:  []int{1, 5, 10},
			esperado: 4, // 5+1+1+1
		},
		{
			name:     "PD supera al ávido (1,3,4 para monto=6)",
			monto:    6,
			monedas:  []int{1, 3, 4},
			esperado: 2, // 3+3, el ávido daría 4+1+1=3
		},
		{
			name:     "monto sin solución",
			monto:    7,
			monedas:  []int{5, 10},
			esperado: -1,
		},
		{
			name:     "una sola denominación",
			monto:    10,
			monedas:  []int{2},
			esperado: 5, // 2+2+2+2+2
		},
		{
			name:     "monto con monedas desordenadas",
			monto:    9,
			monedas:  []int{5, 2, 1},
			esperado: 3, // 5+2+2
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CambioTab(tt.monto, tt.monedas)
			if got != tt.esperado {
				t.Errorf("CambioTab(%d, %v) = %d, esperado %d",
					tt.monto, tt.monedas, got, tt.esperado)
			}
		})
	}
}

func TestCambioMemo(t *testing.T) {
	tests := []struct {
		name     string
		monto    int
		monedas  []int
		esperado int
	}{
		{
			name:     "monto 0",
			monto:    0,
			monedas:  []int{1, 3, 4},
			esperado: 0,
		},
		{
			name:     "caso simple",
			monto:    8,
			monedas:  []int{1, 5, 10},
			esperado: 4,
		},
		{
			name:     "PD supera al ávido",
			monto:    6,
			monedas:  []int{1, 3, 4},
			esperado: 2,
		},
		{
			name:     "sin solución",
			monto:    7,
			monedas:  []int{5, 10},
			esperado: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CambioMemo(tt.monto, tt.monedas)
			if got != tt.esperado {
				t.Errorf("CambioMemo(%d, %v) = %d, esperado %d",
					tt.monto, tt.monedas, got, tt.esperado)
			}
		})
	}
}

func TestAmbosMetodosCoinciden(t *testing.T) {
	monedas := []int{1, 4, 5, 7, 9}
	for m := 0; m <= 30; m++ {
		tab := CambioTab(m, monedas)
		memo := CambioMemo(m, monedas)
		if tab != memo {
			t.Errorf("monto=%d: CambioTab=%d, CambioMemo=%d", m, tab, memo)
		}
	}
}
