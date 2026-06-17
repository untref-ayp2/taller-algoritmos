package iterador

import "testing"

func TestIteradorRecorreElementos(t *testing.T) {
	lista := &ListaEnlazada{}
	lista.InsertarAlInicio(3)
	lista.InsertarAlInicio(2)
	lista.InsertarAlInicio(1)

	it := lista.CrearIterador()
	var valores []int
	for it.Primero(); it.HaySiguiente(); it.Siguiente() {
		valores = append(valores, it.Actual())
	}

	if len(valores) != 3 {
		t.Fatalf("se esperaban 3 elementos, se obtuvieron %d", len(valores))
	}

	// Como se insertaron al inicio, el orden debería ser 1, 2, 3
	esperado := []int{1, 2, 3}
	for i, v := range valores {
		if v != esperado[i] {
			t.Errorf("posición %d: esperado %d, obtenido %d", i, esperado[i], v)
		}
	}
}

func TestIteradorListaVacia(t *testing.T) {
	lista := &ListaEnlazada{}
	it := lista.CrearIterador()

	it.Primero()
	if it.HaySiguiente() {
		t.Error("la lista vacía no debería tener elementos")
	}
}

func TestIteradorUnElemento(t *testing.T) {
	lista := &ListaEnlazada{}
	lista.InsertarAlInicio(42)

	it := lista.CrearIterador()
	it.Primero()
	if !it.HaySiguiente() {
		t.Fatal("debería haber un elemento")
	}
	if it.Actual() != 42 {
		t.Errorf("valor esperado 42, obtenido %d", it.Actual())
	}
	it.Siguiente()
	if it.HaySiguiente() {
		t.Error("no deberían quedar elementos después del único")
	}
}
