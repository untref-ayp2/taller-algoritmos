package persona

// Persona representa una persona con nombre y edad.
type Persona struct {
	Nombre string
	Edad   int
}

// OrdenarPorEdadAsc ordena un slice de Personas por edad ascendente
// usando Counting Sort (rango acotado de edades). Ante edades iguales,
// preserva el orden original (estable).
func OrdenarPorEdadAsc(personas []Persona) []Persona {
	// TODO: implementar usando Counting Sort estable
	return nil
}

// OrdenarPorEdadDesc ordena un slice de Personas por edad descendente
// usando Counting Sort (rango acotado de edades). Ante edades iguales,
// preserva el orden original (estable).
func OrdenarPorEdadDesc(personas []Persona) []Persona {
	// TODO: implementar
	return nil
}
