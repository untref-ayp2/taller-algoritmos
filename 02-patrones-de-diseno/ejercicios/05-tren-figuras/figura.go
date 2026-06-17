package figura

// Figura define la interfaz común para todas las figuras.
type Figura interface {
	Area() float64
}

// Rectangulo representa un rectángulo.
type Rectangulo struct {
	Base   float64
	Altura float64
}

func (r *Rectangulo) Area() float64 {
	// TODO: implementar
	return 0
}

// Circulo representa un círculo.
type Circulo struct {
	Radio float64
}

func (c *Circulo) Area() float64 {
	// TODO: implementar usando math.Pi
	return 0
}

// Triangulo representa un triángulo.
type Triangulo struct {
	Base   float64
	Altura float64
}

func (t *Triangulo) Area() float64 {
	// TODO: implementar
	return 0
}

// Grupo es un compuesto que contiene otras figuras (simples o grupos).
type Grupo struct {
	figuras []Figura
}

// Agregar añade una figura al grupo.
func (g *Grupo) Agregar(f Figura) {
	// TODO: implementar
}

// Area suma el área de todas las figuras del grupo.
func (g *Grupo) Area() float64 {
	// TODO: implementar
	return 0
}
