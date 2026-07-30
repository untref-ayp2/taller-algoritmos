package iterador

type Nodo struct {
	valor int
	ant   *Nodo
	sig   *Nodo
}

// ListaDoble es una lista doblemente enlazada.
type ListaDoble struct {
	cabeza *Nodo
	cola   *Nodo
}

// AgregarAlFinal agrega un elemento al final de la lista.
func (l *ListaDoble) AgregarAlFinal(valor int) {
	// Completar
}

// Iterador permite recorrer la lista en ambos sentidos.
type Iterador struct {
	// Completar
}

// Iterador devuelve un iterador posicionado en el primer elemento de la lista.
func (l *ListaDoble) Iterador() *Iterador {
	// Completar
	return nil
}

// Siguiente avanza al próximo nodo. La primera llamada posiciona el iterador
// en el primer elemento. Devuelve false cuando no hay más elementos.
func (it *Iterador) Siguiente() bool {
	// Completar
	return false
}

// Anterior retrocede al nodo anterior. Devuelve false cuando no hay más
// elementos hacia atrás.
func (it *Iterador) Anterior() bool {
	// Completar
	return false
}

// Valor devuelve el valor del nodo actual.
func (it *Iterador) Valor() int {
	// Completar
	return 0
}
