package mochila

import (
	"fmt"
	"testing"
)

func TestMochilaTab(t *testing.T) {
	items := []Item{
		{Peso: 3, Valor: 2},
		{Peso: 4, Valor: 3},
		{Peso: 1, Valor: 4},
		{Peso: 2, Valor: 2},
		{Peso: 5, Valor: 5},
	}
	cap := 5
	got := MochilaTab(items, cap)
	want := 7 // items 2 (4,3) + 3 (1,4) = peso 5, valor 7
	if got != want {
		t.Errorf("MochilaTab = %d, want %d", got, want)
	}
}

func TestMochilaMemo(t *testing.T) {
	items := []Item{
		{Peso: 3, Valor: 2},
		{Peso: 4, Valor: 3},
		{Peso: 1, Valor: 4},
		{Peso: 2, Valor: 2},
		{Peso: 5, Valor: 5},
	}
	cap := 5
	got := MochilaMemo(items, cap)
	want := 7
	if got != want {
		t.Errorf("MochilaMemo = %d, want %d", got, want)
	}
}

func TestMochilaTabConItems(t *testing.T) {
	items := []Item{
		{Peso: 3, Valor: 2},
		{Peso: 4, Valor: 3},
		{Peso: 1, Valor: 4},
		{Peso: 2, Valor: 2},
		{Peso: 5, Valor: 5},
	}
	cap := 5
	valor, seleccion := MochilaTabConItems(items, cap)
	if valor != 7 {
		t.Errorf("valor = %d, want 7", valor)
	}
	peso := 0
	for _, it := range seleccion {
		peso += it.Peso
	}
	if peso > cap {
		t.Errorf("peso total %d supera capacidad %d", peso, cap)
	}
}

func TestMochilaTabVacia(t *testing.T) {
	if got := MochilaTab(nil, 10); got != 0 {
		t.Errorf("sin items = %d, want 0", got)
	}
}

func TestMochilaMemoVacia(t *testing.T) {
	if got := MochilaMemo(nil, 10); got != 0 {
		t.Errorf("sin items = %d, want 0", got)
	}
}

func TestMochilaTabCapacidadCero(t *testing.T) {
	items := []Item{{Peso: 1, Valor: 10}}
	if got := MochilaTab(items, 0); got != 0 {
		t.Errorf("cap 0 = %d, want 0", got)
	}
}

func TestMochilaMemoCapacidadCero(t *testing.T) {
	items := []Item{{Peso: 1, Valor: 10}}
	if got := MochilaMemo(items, 0); got != 0 {
		t.Errorf("cap 0 = %d, want 0", got)
	}
}

// Ejemplo de uso con los datos del apunte.
func Example() {
	items := []Item{
		{Peso: 3, Valor: 2},
		{Peso: 4, Valor: 3},
		{Peso: 1, Valor: 4},
		{Peso: 2, Valor: 2},
		{Peso: 5, Valor: 5},
	}
	cap := 5
	fmt.Println("Tabulación:", MochilaTab(items, cap))
	fmt.Println("Memoización:", MochilaMemo(items, cap))
	valor, sel := MochilaTabConItems(items, cap)
	fmt.Printf("Valor=%d, Items=%v\n", valor, sel)
	// Output:
	// Tabulación: 7
	// Memoización: 7
	// Valor=7, Items=[{1 4} {4 3}]
}
