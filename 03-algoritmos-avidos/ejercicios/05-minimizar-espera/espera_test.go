package espera

import (
	"reflect"
	"testing"
)

func TestOrdenOptimo(t *testing.T) {
	tiempos := []int{5, 1, 3}
	got := OrdenOptimo(tiempos)
	esperado := []int{1, 2, 0}
	if !reflect.DeepEqual(got, esperado) {
		t.Errorf("OrdenOptimo(%v) = %v; esperado %v", tiempos, got, esperado)
	}
}

func TestUnSoloTrabajo(t *testing.T) {
	got := OrdenOptimo([]int{7})
	esperado := []int{0}
	if !reflect.DeepEqual(got, esperado) {
		t.Errorf("OrdenOptimo([7]) = %v; esperado %v", got, esperado)
	}
}

func TestMismaSumaEspera(t *testing.T) {
	got := OrdenOptimo([]int{3, 3, 3})
	if len(got) != 3 {
		t.Errorf("longitud = %d; esperado 3", len(got))
	}
	for i, idx := range got {
		if idx < 0 || idx >= 3 {
			t.Errorf("índice fuera de rango en posición %d: %d", i, idx)
		}
	}
}
