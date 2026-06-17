package mergesort

import "testing"

func TestMergeSortLista(t *testing.T) {
	lista := &Lista{}
	for _, v := range []int{3, 1, 4, 1, 5, 9} {
		lista.Agregar(v)
	}
	MergeSort(lista)
	// Verificar orden
	act := lista.cabeza
	for act != nil && act.next != nil {
		if act.Valor > act.next.Valor {
			t.Errorf("lista desordenada: %d > %d", act.Valor, act.next.Valor)
		}
		act = act.next
	}
}

func TestMergeSortVacia(t *testing.T) {
	lista := &Lista{}
	MergeSort(lista)
	if lista.cabeza != nil {
		t.Error("lista vacía debería mantener cabeza nil")
	}
}

func TestMergeSortUnElemento(t *testing.T) {
	lista := &Lista{}
	lista.Agregar(42)
	MergeSort(lista)
	if lista.cabeza == nil || lista.cabeza.Valor != 42 {
		t.Error("lista de un elemento debería mantenerse igual")
	}
}
