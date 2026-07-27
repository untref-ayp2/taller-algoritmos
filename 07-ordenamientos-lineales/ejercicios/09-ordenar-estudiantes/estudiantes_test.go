package estudiantes

import (
	"reflect"
	"testing"
)

func TestOrdenarPorPromedio(t *testing.T) {
	estudiantes := []Estudiante{
		{"Ana", 7.50},
		{"Luis", 9.20},
		{"Pedro", 5.00},
	}
	got := OrdenarPorPromedio(estudiantes)
	esperado := []Estudiante{
		{"Pedro", 5.00},
		{"Ana", 7.50},
		{"Luis", 9.20},
	}
	if !reflect.DeepEqual(got, esperado) {
		t.Errorf("OrdenarPorPromedio = %v; esperado %v", got, esperado)
	}
}

func TestOrdenarEstabilidad(t *testing.T) {
	estudiantes := []Estudiante{
		{"Ana", 8.00},
		{"Luis", 8.00},
		{"Pedro", 8.00},
	}
	got := OrdenarPorPromedio(estudiantes)
	if got[0].Nombre != "Ana" || got[1].Nombre != "Luis" || got[2].Nombre != "Pedro" {
		t.Errorf("no se preservó el orden original: %v", got)
	}
}

func TestOrdenarVacio(t *testing.T) {
	got := OrdenarPorPromedio([]Estudiante{})
	if len(got) != 0 {
		t.Error("esperado slice vacío")
	}
}
