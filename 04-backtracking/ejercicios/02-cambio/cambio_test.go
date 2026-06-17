package cambio

import "testing"

func TestFormasDeCambiar(t *testing.T) {
	denom := []int{1, 2, 5}
	monto := 5
	got := FormasDeCambiar(monto, denom)
	if len(got) == 0 {
		t.Fatal("debería encontrar al menos una forma")
	}
	for _, comb := range got {
		suma := 0
		for _, m := range comb {
			suma += m
		}
		if suma != monto {
			t.Errorf("combinación %v suma %d, esperado %d", comb, suma, monto)
		}
	}
	// 5 se puede expresar como: {5}, {2,2,1}, {2,1,1,1}, {1,1,1,1,1}
	if len(got) < 3 {
		t.Errorf("se esperaban al menos 3 formas, se obtuvieron %d", len(got))
	}
}

func TestFormasDeCambiarImposible(t *testing.T) {
	denom := []int{3, 5}
	monto := 7
	got := FormasDeCambiar(monto, denom)
	if len(got) != 0 {
		t.Errorf("debería ser imposible, se obtuvo %v", got)
	}
}

func TestFormasDeCambiarCero(t *testing.T) {
	denom := []int{1, 2}
	got := FormasDeCambiar(0, denom)
	if len(got) != 0 {
		t.Errorf("para monto 0 debería devolver vacío")
	}
}
