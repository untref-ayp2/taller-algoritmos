package bandera

import (
	"reflect"
	"testing"
)

func TestBanderas(t *testing.T) {
	got := OrdenarBanderas([]int{2, 0, 2, 1, 1, 0})
	esperado := []int{0, 0, 1, 1, 2, 2}
	if !reflect.DeepEqual(got, esperado) {
		t.Errorf("OrdenarBanderas = %v; esperado %v", got, esperado)
	}
}

func TestBanderasYaOrdenado(t *testing.T) {
	got := OrdenarBanderas([]int{0, 0, 0})
	esperado := []int{0, 0, 0}
	if !reflect.DeepEqual(got, esperado) {
		t.Errorf("OrdenarBanderas = %v; esperado %v", got, esperado)
	}
}

func TestBanderasVacio(t *testing.T) {
	got := OrdenarBanderas([]int{})
	if len(got) != 0 {
		t.Errorf("esperado slice vacío, obtenido %v", got)
	}
}
