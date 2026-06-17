package iterador

import "testing"

func TestIteradorRecorre(t *testing.T) {
	lista := &Lista{}
	lista.Agregar(1)
	lista.Agregar(2)
	lista.Agregar(3)

	it := lista.Iterador()
	var valores []int
	for it.Siguiente() {
		valores = append(valores, it.Valor())
	}
	if len(valores) != 3 {
		t.Errorf("se esperaban 3 elementos, se obtuvieron %d", len(valores))
	}
}

func TestIteradorVacio(t *testing.T) {
	lista := &Lista{}
	it := lista.Iterador()
	if it.Siguiente() {
		t.Error("la lista vacía no debería tener elementos")
	}
}
