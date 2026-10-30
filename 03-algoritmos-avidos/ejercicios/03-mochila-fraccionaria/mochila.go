package mochila

type Item struct {
	Peso  float64
	Valor float64
}

func (i Item) ValorPorPeso() float64 {
	if i.Peso == 0 {
		return 0
	}
	return i.Valor / i.Peso
}

func MochilaFraccionaria(items []Item, capacidad float64) float64 {
	// TODO: implementar el algoritmo ávido
	return 0
}
