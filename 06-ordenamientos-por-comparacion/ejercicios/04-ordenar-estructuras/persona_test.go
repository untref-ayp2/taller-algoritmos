package persona

import (
	"slices"
	"testing"
)

func personasIguales(a, b []Persona) bool {
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

func TestOrdenarPorEdadAsc(t *testing.T) {
	p := []Persona{
		{Nombre: "ana", Edad: 30},
		{Nombre: "beto", Edad: 20},
		{Nombre: "cata", Edad: 25},
	}
	OrdenarPorEdadAsc(p)
	if !slices.IsSortedFunc(p, func(a, b Persona) int {
		if a.Edad < b.Edad {
			return -1
		}
		if a.Edad > b.Edad {
			return 1
		}
		return 0
	}) {
		t.Errorf("no ordenado por edad asc: %+v", p)
	}
}

func TestOrdenarPorNombreDesc(t *testing.T) {
	p := []Persona{
		{Nombre: "ana", Edad: 30},
		{Nombre: "cata", Edad: 25},
		{Nombre: "beto", Edad: 20},
	}
	OrdenarPorNombreDesc(p)
	for i := 1; i < len(p); i++ {
		if p[i-1].Nombre < p[i].Nombre {
			t.Errorf("no ordenado por nombre desc en pos %d: %+v", i, p)
		}
	}
}

func TestOrdenarPorEdadAscNombreDesc(t *testing.T) {
	p := []Persona{
		{Nombre: "ana", Edad: 25},
		{Nombre: "cata", Edad: 30},
		{Nombre: "beto", Edad: 25},
		{Nombre: "dani", Edad: 20},
	}
	OrdenarPorEdadAscNombreDesc(p)
	for i := 1; i < len(p); i++ {
		if p[i-1].Edad > p[i].Edad {
			t.Errorf("no ordenado por edad asc en pos %d: %+v", i, p)
		}
		if p[i-1].Edad == p[i].Edad && p[i-1].Nombre < p[i].Nombre {
			t.Errorf("misma edad no ordenado por nombre desc en pos %d: %+v", i, p)
		}
	}
}

func TestOrdenarEmpty(t *testing.T) {
	p := []Persona{}
	OrdenarPorEdadAsc(p)
	OrdenarPorNombreDesc(p)
	OrdenarPorEdadAscNombreDesc(p)
}

func TestOrdenarSingle(t *testing.T) {
	p := []Persona{{Nombre: "unico", Edad: 10}}
	OrdenarPorEdadAsc(p)
	if p[0].Nombre != "unico" {
		t.Errorf("se modificó un elemento")
	}
}

func TestOrdenarEdadAscEstable(t *testing.T) {
	p := []Persona{
		{Nombre: "a", Edad: 30},
		{Nombre: "b", Edad: 20},
		{Nombre: "c", Edad: 30},
	}
	OrdenarPorEdadAsc(p)
	if !(p[0].Nombre == "b" && p[1].Nombre == "a" && p[2].Nombre == "c") {
		t.Errorf("no se preservó orden relativo para misma edad: %+v", p)
	}
}
