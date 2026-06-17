package main

import "fmt"

// Nodo de la lista enlazada.
type Nodo struct {
	Valor     int
	Siguiente *Nodo
}

// ListaEnlazada es la colección.
type ListaEnlazada struct {
	Primero *Nodo
}

func (l *ListaEnlazada) InsertarAlInicio(valor int) {
	if l.Primero == nil {
		l.Primero = &Nodo{Valor: valor}
	} else {
		nuevoNodo := &Nodo{Valor: valor, Siguiente: l.Primero}
		l.Primero = nuevoNodo
	}
}

// Iterador define el comportamiento para recorrer la colección.
type Iterador interface {
	Primero()
	Siguiente()
	HaySiguiente() bool
	Actual() int
}

// IteradorLista implementa Iterador para ListaEnlazada.
type IteradorLista struct {
	lista  *ListaEnlazada
	actual *Nodo
}

func (l *ListaEnlazada) CrearIterador() Iterador {
	return &IteradorLista{lista: l, actual: l.Primero}
}

func (it *IteradorLista) Primero() {
	it.actual = it.lista.Primero
}

func (it *IteradorLista) Siguiente() {
	it.actual = it.actual.Siguiente
}

func (it *IteradorLista) HaySiguiente() bool {
	return it.actual != nil
}

func (it *IteradorLista) Actual() int {
	return it.actual.Valor
}

func main() {
	lista := &ListaEnlazada{}
	lista.InsertarAlInicio(3)
	lista.InsertarAlInicio(2)
	lista.InsertarAlInicio(1)

	iterador := lista.CrearIterador()
	for iterador.Primero(); iterador.HaySiguiente(); iterador.Siguiente() {
		fmt.Println(iterador.Actual())
	}
	// Salida:
	// 1
	// 2
	// 3
}
