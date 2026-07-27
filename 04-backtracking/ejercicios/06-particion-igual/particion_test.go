package particion

import "testing"

func TestParticionValida(t *testing.T) {
	a, b, ok := ParticionIgual([]int{1, 5, 11, 5})
	if !ok {
		t.Fatal("esperado true")
	}
	if suma(a) != suma(b) {
		t.Errorf("sumas distintas: %d vs %d", suma(a), suma(b))
	}
}

func TestSumaImpar(t *testing.T) {
	_, _, ok := ParticionIgual([]int{1, 2, 3, 5})
	if ok {
		t.Error("esperado false para suma total impar (11)")
	}
}

func TestParticionPares(t *testing.T) {
	a, b, ok := ParticionIgual([]int{2, 2, 2, 2})
	if !ok {
		t.Fatal("esperado true")
	}
	if suma(a) != suma(b) {
		t.Errorf("sumas distintas: %d vs %d", suma(a), suma(b))
	}
}

func suma(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}
