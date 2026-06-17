package sesion

import (
	"reflect"
	"testing"
)

func TestMaxSesiones(t *testing.T) {
	sesiones := []Sesion{
		{1, 3},
		{2, 5},
		{3, 7},
		{5, 6},
		{6, 8},
		{8, 9},
	}
	got := MaxSesiones(sesiones)
	if len(got) == 0 {
		t.Fatal("MaxSesiones devolvió slice vacío")
	}
	// Verificar que no se superpongan
	for i := 1; i < len(got); i++ {
		if got[i].Inicio < got[i-1].Fin {
			t.Errorf("sesiones superpuestas: %v y %v", got[i-1], got[i])
		}
	}
}

func TestMaxSesionesOptimo(t *testing.T) {
	sesiones := []Sesion{
		{0, 6},
		{1, 2},
		{3, 4},
		{5, 7},
	}
	got := MaxSesiones(sesiones)
	// La solución óptima es [{1,2}, {3,4}, {5,7}] (3 sesiones)
	if len(got) != 3 {
		t.Errorf("se esperaban 3 sesiones, se obtuvieron %d", len(got))
	}
}

func TestMaxSesionesVacio(t *testing.T) {
	got := MaxSesiones(nil)
	if got == nil {
		t.Error("debe devolver slice vacío, no nil")
	}
	if len(got) != 0 {
		t.Errorf("se esperaba slice vacío, longitud %d", len(got))
	}
}

func TestMaxSesionesUna(t *testing.T) {
	sesiones := []Sesion{{1, 2}}
	got := MaxSesiones(sesiones)
	if len(got) != 1 {
		t.Errorf("se esperaba 1 sesión, se obtuvieron %d", len(got))
	}
	if !reflect.DeepEqual(got, sesiones) {
		t.Errorf("sesión modificada: %v", got)
	}
}
