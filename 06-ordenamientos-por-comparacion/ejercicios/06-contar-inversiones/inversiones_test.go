package inversiones

import "testing"

func TestContarInversiones(t *testing.T) {
	got := ContarInversiones([]int{2, 4, 1, 3, 5})
	if got != 3 {
		t.Errorf("ContarInversiones = %d; esperado 3", got)
	}
}

func TestSinInversiones(t *testing.T) {
	got := ContarInversiones([]int{1, 2, 3, 4})
	if got != 0 {
		t.Errorf("ContarInversiones = %d; esperado 0", got)
	}
}

func TestOrdenInverso(t *testing.T) {
	got := ContarInversiones([]int{5, 4, 3, 2, 1})
	esperado := 10
	if got != esperado {
		t.Errorf("ContarInversiones = %d; esperado %d", got, esperado)
	}
}
