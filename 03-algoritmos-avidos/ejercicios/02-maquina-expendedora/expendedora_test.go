package expendedora

import "testing"

func TestCambioAvido(t *testing.T) {
	denom := []int{1, 2, 5, 10, 20, 50, 100}
	cambio := CambioAvido(93, denom)
	// 93 = 50 + 20 + 20 + 2 + 1 → 5 monedas
	total := 0
	cant := 0
	for d, c := range cambio {
		total += d * c
		cant += c
	}
	if total != 93 {
		t.Errorf("suma del cambio = %d, esperado 93", total)
	}
	if cant == 0 {
		t.Error("no se devolvieron monedas")
	}
}

func TestCambioAvidoCero(t *testing.T) {
	denom := []int{1, 2, 5}
	cambio := CambioAvido(0, denom)
	if len(cambio) != 0 {
		t.Errorf("para monto 0 se esperaba mapa vacío, se obtuvo %v", cambio)
	}
}
