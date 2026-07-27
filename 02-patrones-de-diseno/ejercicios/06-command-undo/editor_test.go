package editor

import "testing"

func TestInsertar(t *testing.T) {
	e := NuevoEditor()
	h := NuevoHistorial(e)
	h.Ejecutar(ComandoInsertar{texto: "Hola", pos: 0})
	if e.Contenido() != "Hola" {
		t.Errorf("contenido = %q; esperado %q", e.Contenido(), "Hola")
	}
}

func TestInsertarYDeshacer(t *testing.T) {
	e := NuevoEditor()
	h := NuevoHistorial(e)
	h.Ejecutar(ComandoInsertar{texto: "Hola", pos: 0})
	h.Ejecutar(ComandoInsertar{texto: " Mundo", pos: 4})
	if e.Contenido() != "Hola Mundo" {
		t.Fatalf("contenido = %q; esperado %q", e.Contenido(), "Hola Mundo")
	}
	err := h.Deshacer()
	if err != nil {
		t.Fatalf("inesperado error: %v", err)
	}
	if e.Contenido() != "Hola" {
		t.Errorf("luego de undo = %q; esperado %q", e.Contenido(), "Hola")
	}
}

func TestDeshacerYRehacer(t *testing.T) {
	e := NuevoEditor()
	h := NuevoHistorial(e)
	h.Ejecutar(ComandoInsertar{texto: "Hola", pos: 0})
	h.Ejecutar(ComandoInsertar{texto: " Mundo", pos: 4})
	h.Deshacer()
	err := h.Rehacer()
	if err != nil {
		t.Fatalf("inesperado error: %v", err)
	}
	if e.Contenido() != "Hola Mundo" {
		t.Errorf("luego de redo = %q; esperado %q", e.Contenido(), "Hola Mundo")
	}
}

func TestDeshacerVacio(t *testing.T) {
	e := NuevoEditor()
	h := NuevoHistorial(e)
	err := h.Deshacer()
	if err == nil {
		t.Error("esperado error al deshacer con historial vacío")
	}
}
