package potencia

import (
	"math"
	"testing"
)

func TestPotencia(t *testing.T) {
	got := Potencia(2, 10)
	if got != 1024 {
		t.Errorf("Potencia(2, 10) = %v; esperado 1024", got)
	}
}

func TestPotenciaNegativa(t *testing.T) {
	got := Potencia(2, -3)
	if math.Abs(got-0.125) > 0.0001 {
		t.Errorf("Potencia(2, -3) = %v; esperado 0.125", got)
	}
}

func TestPotenciaCero(t *testing.T) {
	got := Potencia(5, 0)
	if got != 1 {
		t.Errorf("Potencia(5, 0) = %v; esperado 1", got)
	}
}
