package mochila

import (
	"math"
	"testing"
)

func TestMochilaFraccionaria(t *testing.T) {
	items := []Item{
		{Peso: 10, Valor: 60},
		{Peso: 20, Valor: 100},
		{Peso: 30, Valor: 120},
	}
	got := MochilaFraccionaria(items, 50)
	esperado := 240.0
	if math.Abs(got-esperado) > 0.001 {
		t.Errorf("MochilaFraccionaria(cap=50) = %v; esperado %v", got, esperado)
	}
}

func TestMochilaCapacidadCero(t *testing.T) {
	items := []Item{{Peso: 10, Valor: 60}}
	got := MochilaFraccionaria(items, 0)
	if got != 0 {
		t.Errorf("MochilaFraccionaria(cap=0) = %v; esperado 0", got)
	}
}

func TestMochilaFraccionar(t *testing.T) {
	items := []Item{{Peso: 10, Valor: 60}}
	got := MochilaFraccionaria(items, 5)
	esperado := 30.0
	if math.Abs(got-esperado) > 0.001 {
		t.Errorf("MochilaFraccionaria(cap=5) = %v; esperado %v", got, esperado)
	}
}
