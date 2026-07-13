package persona

import (
	"reflect"
	"testing"
)

func TestOrdenarPorEdadAsc(t *testing.T) {
	input := []Persona{
		{"Ana", 25},
		{"Luis", 30},
		{"Pedro", 20},
		{"Maria", 30},
		{"Juan", 25},
	}
	got := OrdenarPorEdadAsc(input)
	expected := []Persona{
		{"Pedro", 20},
		{"Ana", 25},
		{"Juan", 25},
		{"Luis", 30},
		{"Maria", 30},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("OrdenarPorEdadAsc incorrecto.\n  got:      %v\n  esperado: %v", got, expected)
	}
}

func TestOrdenarPorEdadDesc(t *testing.T) {
	input := []Persona{
		{"Ana", 25},
		{"Luis", 30},
		{"Pedro", 20},
		{"Maria", 30},
		{"Juan", 25},
	}
	got := OrdenarPorEdadDesc(input)
	expected := []Persona{
		{"Luis", 30},
		{"Maria", 30},
		{"Ana", 25},
		{"Juan", 25},
		{"Pedro", 20},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("OrdenarPorEdadDesc incorrecto.\n  got:      %v\n  esperado: %v", got, expected)
	}
}

func TestOrdenarPorEdadAscEstable(t *testing.T) {
	// Dos personas con misma edad: debe preservar orden original
	input := []Persona{
		{"Primero", 25},
		{"Segundo", 25},
	}
	got := OrdenarPorEdadAsc(input)
	if got[0].Nombre != "Primero" || got[1].Nombre != "Segundo" {
		t.Errorf("Counting Sort no es estable: esperado [Primero, Segundo], got %v", got)
	}
}

func TestOrdenarPorEdadDescEstable(t *testing.T) {
	input := []Persona{
		{"Primero", 30},
		{"Segundo", 30},
	}
	got := OrdenarPorEdadDesc(input)
	if got[0].Nombre != "Primero" || got[1].Nombre != "Segundo" {
		t.Errorf("Counting Sort descendente no es estable: esperado [Primero, Segundo], got %v", got)
	}
}

func TestOrdenarVacio(t *testing.T) {
	if got := OrdenarPorEdadAsc([]Persona{}); got == nil {
		t.Error("Se esperaba slice vacío, no nil")
	}
	if got := OrdenarPorEdadDesc([]Persona{}); got == nil {
		t.Error("Se esperaba slice vacío, no nil")
	}
}
