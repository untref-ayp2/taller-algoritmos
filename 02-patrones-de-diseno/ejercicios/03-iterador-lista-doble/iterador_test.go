package iterador

import "testing"

func TestSiguiente(t *testing.T) {
	lista := &ListaDoble{}
	lista.AgregarAlFinal(10)
	lista.AgregarAlFinal(20)
	lista.AgregarAlFinal(30)

	it := lista.Iterador()
	var valores []int
	for it.Siguiente() {
		valores = append(valores, it.Valor())
	}

	if len(valores) != 3 {
		t.Fatalf("esperaba 3 elementos, obtuve %d: %v", len(valores), valores)
	}
	esperados := []int{10, 20, 30}
	for i, esp := range esperados {
		if valores[i] != esp {
			t.Errorf("posición %d: esperaba %d, obtuve %d", i, esp, valores[i])
		}
	}
}

func TestAnterior(t *testing.T) {
	lista := &ListaDoble{}
	lista.AgregarAlFinal(10)
	lista.AgregarAlFinal(20)
	lista.AgregarAlFinal(30)

	// Avanzar hasta el final
	it := lista.Iterador()
	for it.Siguiente() {
	}
	// it quedó después del último; usar Anterior para retroceder

	var valores []int
	for it.Anterior() {
		valores = append(valores, it.Valor())
	}

	if len(valores) != 3 {
		t.Fatalf("esperaba 3 elementos, obtuve %d: %v", len(valores), valores)
	}
	esperados := []int{30, 20, 10}
	for i, esp := range esperados {
		if valores[i] != esp {
			t.Errorf("posición %d: esperaba %d, obtuve %d", i, esp, valores[i])
		}
	}
}

func TestAvanzarYRetroceder(t *testing.T) {
	lista := &ListaDoble{}
	lista.AgregarAlFinal(1)
	lista.AgregarAlFinal(2)
	lista.AgregarAlFinal(3)

	it := lista.Iterador()

	// Avanzar dos posiciones
	if !it.Siguiente() || it.Valor() != 1 {
		t.Fatal("primer Siguiente debería ser 1")
	}
	if !it.Siguiente() || it.Valor() != 2 {
		t.Fatal("segundo Siguiente debería ser 2")
	}

	// Retroceder una posición
	if !it.Anterior() || it.Valor() != 1 {
		t.Fatal("Anterior debería volver a 1")
	}

	// Avanzar de nuevo
	if !it.Siguiente() || it.Valor() != 2 {
		t.Fatal("Siguiente después de Anterior debería ser 2")
	}
}

func TestVacia(t *testing.T) {
	lista := &ListaDoble{}
	it := lista.Iterador()

	if it.Siguiente() {
		t.Error("lista vacía: Siguiente debería devolver false")
	}
	if it.Anterior() {
		t.Error("lista vacía: Anterior debería devolver false")
	}
}
