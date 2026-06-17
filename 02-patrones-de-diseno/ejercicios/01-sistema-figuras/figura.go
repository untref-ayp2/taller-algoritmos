package figura

type Figura interface {
	Area() float64
	Perimetro() float64
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	// TODO: implementar
	return 0
}

func (c Circulo) Perimetro() float64 {
	// TODO: implementar
	return 0
}

type Rectangulo struct {
	Base, Altura float64
}

func (r Rectangulo) Area() float64 {
	// TODO: implementar
	return 0
}

func (r Rectangulo) Perimetro() float64 {
	// TODO: implementar
	return 0
}

type Dibujo struct {
	figuras []Figura
}

func (d *Dibujo) Agregar(f Figura) {
	d.figuras = append(d.figuras, f)
}

func (d Dibujo) Area() float64 {
	// TODO: implementar
	return 0
}

func (d Dibujo) Perimetro() float64 {
	// TODO: implementar
	return 0
}
