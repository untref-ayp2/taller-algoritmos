package main

import (
	"fmt"
	"math"
)

// Figura define la interfaz común.
type Figura interface {
	Area() float64
}

// Rectangulo (Simple)
type Rectangulo struct {
	Base   float64
	Altura float64
}

func (r *Rectangulo) Area() float64 {
	return r.Base * r.Altura
}

// Circulo (Simple)
type Circulo struct {
	Radio float64
}

func (c *Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

// Triangulo (Simple)
type Triangulo struct {
	Base   float64
	Altura float64
}

func (t *Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

// Grupo (Compuesto)
type Grupo struct {
	figuras []Figura
}

func (g *Grupo) Agregar(f Figura) {
	g.figuras = append(g.figuras, f)
}

func (g *Grupo) Area() float64 {
	var areaTotal float64
	for _, f := range g.figuras {
		areaTotal += f.Area()
	}
	return areaTotal
}

func main() {
	locomotora := &Grupo{}
	locomotora.Agregar(&Rectangulo{Base: 7, Altura: 3})
	locomotora.Agregar(&Circulo{Radio: 1})
	locomotora.Agregar(&Circulo{Radio: 1})
	locomotora.Agregar(&Triangulo{Base: 2, Altura: 4})
	locomotora.Agregar(&Rectangulo{Base: 2, Altura: 3})

	vagon1 := &Grupo{}
	vagon1.Agregar(&Rectangulo{Base: 7, Altura: 3})
	vagon1.Agregar(&Circulo{Radio: 1})
	vagon1.Agregar(&Circulo{Radio: 1})

	vagon2 := &Grupo{}
	vagon2.Agregar(&Rectangulo{Base: 7, Altura: 3})
	vagon2.Agregar(&Circulo{Radio: 1})
	vagon2.Agregar(&Circulo{Radio: 1})

	tren := &Grupo{}
	tren.Agregar(locomotora)
	tren.Agregar(vagon1)
	tren.Agregar(vagon2)

	fmt.Println("El área del tren es:", tren.Area())
	// Salida: El área del tren es: 91.84955592153875
}
