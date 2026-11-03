package ordenarlista

import "testing"

func listaDesdeSlice(vals []int) *Nodo {
	if len(vals) == 0 {
		return nil
	}
	cabeza := &Nodo{Valor: vals[0]}
	actual := cabeza
	for i := 1; i < len(vals); i++ {
		actual.Sig = &Nodo{Valor: vals[i]}
		actual = actual.Sig
	}
	return cabeza
}

func sliceDesdeLista(cabeza *Nodo) []int {
	var res []int
	for n := cabeza; n != nil; n = n.Sig {
		res = append(res, n.Valor)
	}
	return res
}

func intsIgual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMergeSortLista(t *testing.T) {
	cabeza := listaDesdeSlice([]int{4, 2, 1, 3})
	ordenada := MergeSortLista(cabeza)
	got := sliceDesdeLista(ordenada)
	esperado := []int{1, 2, 3, 4}
	if !intsIgual(got, esperado) {
		t.Errorf("MergeSortLista = %v; esperado %v", got, esperado)
	}
}

func TestMergeSortListaVacia(t *testing.T) {
	if MergeSortLista(nil) != nil {
		t.Error("esperado nil para lista vacía")
	}
}

func TestMergeSortListaUnitaria(t *testing.T) {
	cabeza := &Nodo{Valor: 42}
	ordenada := MergeSortLista(cabeza)
	if ordenada == nil || ordenada.Valor != 42 || ordenada.Sig != nil {
		t.Error("lista unitaria no preservada")
	}
}
