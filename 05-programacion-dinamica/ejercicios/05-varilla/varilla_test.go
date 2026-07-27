package varilla

import "testing"

func TestCorteVarilla(t *testing.T) {
	precios := []int{1, 5, 8, 9}
	got := CorteVarilla(4, precios)
	if got != 10 {
		t.Errorf("CorteVarilla(4) = %d; esperado 10", got)
	}
}

func TestCorteVarillaCero(t *testing.T) {
	got := CorteVarilla(0, []int{1, 5, 8})
	if got != 0 {
		t.Errorf("CorteVarilla(0) = %d; esperado 0", got)
	}
}

func TestCorteVarillaLarga(t *testing.T) {
	precios := []int{1, 5, 8, 9, 10, 17, 17, 20}
	got := CorteVarilla(8, precios)
	if got != 22 {
		t.Errorf("CorteVarilla(8) = %d; esperado 22", got)
	}
}
