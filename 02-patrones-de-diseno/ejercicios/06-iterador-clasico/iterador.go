package iterador

// Nodo representa un elemento de la lista enlazada.
type Nodo struct {
	Valor     int
	Siguiente *Nodo
}

// ListaEnlazada es una colección de nodos.
type ListaEnlazada struct {
	Primero *Nodo
}

// InsertarAlInicio agrega un valor al comienzo de la lista.
func (l *ListaEnlazada) InsertarAlInicio(valor int) {
	// TODO: implementar
}

// Iterador define los métodos para recorrer una colección.
type Iterador interface {
	Primero()
	Siguiente()
	HaySiguiente() bool
	Actual() int
}

// IteradorLista implementa Iterador para una ListaEnlazada.
type IteradorLista struct {
	lista  *ListaEnlazada
	actual *Nodo
}

// CrearIterador devuelve un nuevo iterador para la lista.
func (l *ListaEnlazada) CrearIterador() Iterador {
	// TODO: devolver un IteradorLista apuntando al primer nodo
	return &IteradorLista{lista: l}
}

// Primero se posiciona en el primer elemento de la colección.
func (it *IteradorLista) Primero() {
	// TODO: posicionar it.actual en l.Primero
}

// Siguiente avanza al siguiente elemento.
func (it *IteradorLista) Siguiente() {
	// TODO: avanzar it.actual al siguiente nodo
}

// HaySiguiente devuelve true si todavía quedan elementos.
func (it *IteradorLista) HaySiguiente() bool {
	// TODO: implementar
	return false
}

// Actual devuelve el valor del elemento actual.
func (it *IteradorLista) Actual() int {
	// TODO: implementar
	return 0
}
