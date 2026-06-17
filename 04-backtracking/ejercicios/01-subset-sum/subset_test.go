package subset

import "testing"

func TestSubsetSumEncuentra(t *testing.T) {
	nums := []int{3, 7, 1, 8, 4}
	k := 10
	got := SubsetSum(nums, k)
	if got == nil {
		t.Fatal("SubsetSum devolvió nil, debería encontrar solución")
	}
	suma := 0
	for _, n := range got {
		suma += n
	}
	if suma != k {
		t.Errorf("suma = %d, esperado %d", suma, k)
	}
}

func TestSubsetSumNoExiste(t *testing.T) {
	nums := []int{3, 7, 1, 8, 4}
	k := 99
	got := SubsetSum(nums, k)
	if got != nil {
		t.Errorf("debería devolver nil, se obtuvo %v", got)
	}
}

func TestTodasLasCombinaciones(t *testing.T) {
	nums := []int{2, 3, 6, 7}
	k := 7
	got := TodasLasCombinaciones(nums, k)
	if len(got) == 0 {
		t.Fatal("debería encontrar combinaciones")
	}
	for _, comb := range got {
		suma := 0
		for _, n := range comb {
			suma += n
		}
		if suma != k {
			t.Errorf("combinación %v suma %d, esperado %d", comb, suma, k)
		}
	}
}
