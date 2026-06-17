package figura

import "testing"

func TestCirculoArea(t *testing.T) {
	c := Circulo{Radio: 5}
	got := c.Area()
	if got <= 0 {
		t.Errorf("Circulo.Area() = %v; esperado > 0", got)
	}
}

func TestRectanguloPerimetro(t *testing.T) {
	r := Rectangulo{Base: 3, Altura: 4}
	got := r.Perimetro()
	want := 14.0
	if got != want {
		t.Errorf("Rectangulo.Perimetro() = %v; esperado %v", got, want)
	}
}

func TestDibujoArea(t *testing.T) {
	d := &Dibujo{}
	d.Agregar(Circulo{Radio: 1})
	d.Agregar(Rectangulo{Base: 2, Altura: 3})
	got := d.Area()
	if got <= 0 {
		t.Errorf("Dibujo.Area() = %v; esperado > 0", got)
	}
}
