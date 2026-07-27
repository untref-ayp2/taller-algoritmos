package figura

import (
	"math"
	"testing"
)

func TestCrearCirculo(t *testing.T) {
	f, err := CrearFigura("circulo 5")
	if err != nil {
		t.Fatalf("inesperado error: %v", err)
	}
	if f.Nombre() != "Círculo" {
		t.Errorf("nombre = %q; esperado %q", f.Nombre(), "Círculo")
	}
	if math.Abs(f.Area()-78.5398) > 0.01 {
		t.Errorf("area = %v; esperado ~78.54", f.Area())
	}
}

func TestCrearRectangulo(t *testing.T) {
	f, err := CrearFigura("rectangulo 3 4")
	if err != nil {
		t.Fatalf("inesperado error: %v", err)
	}
	if f.Nombre() != "Rectángulo" {
		t.Errorf("nombre = %q; esperado %q", f.Nombre(), "Rectángulo")
	}
	if f.Area() != 12 {
		t.Errorf("area = %v; esperado 12", f.Area())
	}
}

func TestCrearFiguraInvalida(t *testing.T) {
	_, err := CrearFigura("triangulo 3 4 5")
	if err == nil {
		t.Error("esperado error para figura no reconocida")
	}
}
