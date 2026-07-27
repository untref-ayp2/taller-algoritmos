package figura

import "math"

type Figura interface {
	Area() float64
	Nombre() string
}

type Circulo struct {
	radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c Circulo) Nombre() string {
	return "Círculo"
}

type Rectangulo struct {
	base, altura float64
}

func (r Rectangulo) Area() float64 {
	return r.base * r.altura
}

func (r Rectangulo) Nombre() string {
	return "Rectángulo"
}

type Cuadrado struct {
	lado float64
}

func (c Cuadrado) Area() float64 {
	return c.lado * c.lado
}

func (c Cuadrado) Nombre() string {
	return "Cuadrado"
}

func CrearFigura(descripcion string) (Figura, error) {
	// TODO: implementar el parser de strings como "circulo 5", "rectangulo 3 4", "cuadrado 2"
	return nil, nil
}
