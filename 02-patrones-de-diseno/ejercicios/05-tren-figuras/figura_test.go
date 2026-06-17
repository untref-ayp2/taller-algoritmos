package figura

import (
	"math"
	"testing"
)

func TestRectanguloArea(t *testing.T) {
	r := &Rectangulo{Base: 3, Altura: 4}
	got := r.Area()
	want := 12.0
	if got != want {
		t.Errorf("Rectangulo.Area() = %v; esperado %v", got, want)
	}
}

func TestCirculoArea(t *testing.T) {
	c := &Circulo{Radio: 1}
	got := c.Area()
	want := math.Pi
	if got != want {
		t.Errorf("Circulo.Area() = %v; esperado %v", got, want)
	}
}

func TestTrianguloArea(t *testing.T) {
	tri := &Triangulo{Base: 4, Altura: 3}
	got := tri.Area()
	want := 6.0
	if got != want {
		t.Errorf("Triangulo.Area() = %v; esperado %v", got, want)
	}
}

func TestGrupoArea(t *testing.T) {
	g := &Grupo{}
	g.Agregar(&Rectangulo{Base: 2, Altura: 3})
	g.Agregar(&Circulo{Radio: 1})
	got := g.Area()
	if got <= 0 {
		t.Errorf("Grupo.Area() = %v; esperado > 0", got)
	}
}

func TestGrupoVacio(t *testing.T) {
	g := &Grupo{}
	got := g.Area()
	if got != 0 {
		t.Errorf("Grupo vacío debería tener área 0, se obtuvo %v", got)
	}
}

func TestGrupoAnidado(t *testing.T) {
	sub := &Grupo{}
	sub.Agregar(&Circulo{Radio: 2})
	sub.Agregar(&Rectangulo{Base: 1, Altura: 1})

	padre := &Grupo{}
	padre.Agregar(sub)
	padre.Agregar(&Triangulo{Base: 3, Altura: 2})

	got := padre.Area()
	if got <= 0 {
		t.Errorf("Grupo anidado.Area() = %v; esperado > 0", got)
	}
}
