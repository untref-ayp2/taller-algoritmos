package mochila

import "testing"

func TestMochila01(t *testing.T) {
	objetos := []Objeto{
		{Peso: 2, Valor: 3},
		{Peso: 3, Valor: 4},
		{Peso: 4, Valor: 5},
		{Peso: 5, Valor: 6},
	}
	capacidad := 5
	valor, seleccion := Mochila01(objetos, capacidad)
	// Óptimo: objeto {2,3} + {3,4} = valor 7
	if valor != 7 {
		t.Errorf("valor máximo = %d, esperado 7", valor)
	}
	totalPeso := 0
	for _, o := range seleccion {
		totalPeso += o.Peso
	}
	if totalPeso > capacidad {
		t.Errorf("peso total %d supera capacidad %d", totalPeso, capacidad)
	}
}

func TestMochila01Vacia(t *testing.T) {
	valor, seleccion := Mochila01(nil, 10)
	if valor != 0 || len(seleccion) != 0 {
		t.Errorf("sin objetos: valor=%d, selección=%v", valor, seleccion)
	}
}

func TestMochila01CeroCapacidad(t *testing.T) {
	objetos := []Objeto{{Peso: 1, Valor: 10}}
	valor, seleccion := Mochila01(objetos, 0)
	if valor != 0 || len(seleccion) != 0 {
		t.Errorf("capacidad 0: valor=%d, selección=%v", valor, seleccion)
	}
}
