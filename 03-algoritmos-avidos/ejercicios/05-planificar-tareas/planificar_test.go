package planificar

import (
	"reflect"
	"testing"
)

func TestPlanificarTareas(t *testing.T) {
	trabajos := []Trabajo{
		{ID: 1, Deadline: 2, Ganancia: 100},
		{ID: 2, Deadline: 1, Ganancia: 50},
		{ID: 3, Deadline: 2, Ganancia: 200},
	}
	got := PlanificarTareas(trabajos)
	esperado := []int{1, 3}
	if !reflect.DeepEqual(got, esperado) {
		t.Errorf("PlanificarTareas = %v; esperado %v", got, esperado)
	}
}

func TestDeadlineCero(t *testing.T) {
	trabajos := []Trabajo{
		{ID: 1, Deadline: 0, Ganancia: 100},
		{ID: 2, Deadline: 0, Ganancia: 50},
	}
	got := PlanificarTareas(trabajos)
	if len(got) != 0 {
		t.Errorf("PlanificarTareas = %v; esperado []", got)
	}
}

func TestUnSoloTrabajo(t *testing.T) {
	trabajos := []Trabajo{{ID: 1, Deadline: 1, Ganancia: 42}}
	got := PlanificarTareas(trabajos)
	esperado := []int{1}
	if !reflect.DeepEqual(got, esperado) {
		t.Errorf("PlanificarTareas = %v; esperado %v", got, esperado)
	}
}
