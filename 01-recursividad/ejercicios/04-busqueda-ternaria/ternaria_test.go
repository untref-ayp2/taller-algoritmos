package ternaria

import "testing"

func TestBusquedaTernaria(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	tests := []struct {
		x    int
		want int
	}{
		{1, 0},
		{7, 3},
		{19, 9},
		{10, -1},
		{0, -1},
		{20, -1},
	}
	for _, tt := range tests {
		got := BusquedaTernaria(arr, tt.x)
		if got != tt.want {
			t.Errorf("BusquedaTernaria(%d) = %d; esperado %d", tt.x, got, tt.want)
		}
	}
}

func TestBusquedaTernariaArregloVacio(t *testing.T) {
	got := BusquedaTernaria([]int{}, 5)
	if got != -1 {
		t.Errorf("BusquedaTernaria en arreglo vacío = %d; esperado -1", got)
	}
}
