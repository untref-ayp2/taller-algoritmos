package grilla

import (
	"reflect"
	"testing"
)

func TestCaminoMinimo(t *testing.T) {
	grilla := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	got := CaminoMinimo(grilla)
	esperado := 7 // 1→3→1→1→1
	if got != esperado {
		t.Errorf("CaminoMinimo = %d; esperado %d", got, esperado)
	}
}

func TestCaminoMinimoUnaCelda(t *testing.T) {
	got := CaminoMinimo([][]int{{5}})
	if got != 5 {
		t.Errorf("CaminoMinimo = %d; esperado 5", got)
	}
}

func TestCaminoMinimoUnaFila(t *testing.T) {
	grilla := [][]int{{1, 2, 3, 4, 5}}
	got := CaminoMinimo(grilla)
	esperado := 15
	if got != esperado {
		t.Errorf("CaminoMinimo = %d; esperado %d", got, esperado)
	}
}

func TestCaminoMinimoConRuta(t *testing.T) {
	grilla := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	costo, ruta := CaminoMinimoConRuta(grilla)
	if costo != 7 {
		t.Errorf("costo = %d; esperado 7", costo)
	}
	esperado := []Pos{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}}
	if !reflect.DeepEqual(ruta, esperado) {
		t.Errorf("ruta = %v; esperado %v", ruta, esperado)
	}
}
