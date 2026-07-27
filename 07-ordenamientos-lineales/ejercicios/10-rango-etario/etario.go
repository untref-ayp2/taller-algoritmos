package etario

type Persona struct {
	Nombre string
	Edad   int
}

type GrupoEtario struct {
	Rango     string
	Personas  []Persona
}

func AgruparPorEdad(personas []Persona) []GrupoEtario {
	// TODO: implementar con Counting Sort
	return nil
}
