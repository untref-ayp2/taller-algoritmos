package main

import "fmt"

// Nodo de la lista enlazada.
type Nodo struct {
	valor int
	sig   *Nodo
}

// Lista es la colección.
type Lista struct {
	cabeza *Nodo
}

func (l *Lista) AgregarAlFinal(valor int) {
	nuevo := &Nodo{valor: valor}
	if l.cabeza == nil {
		l.cabeza = nuevo
		return
	}
	actual := l.cabeza
	for actual.sig != nil {
		actual = actual.sig
	}
	actual.sig = nuevo
}

// Iterador permite recorrer la lista sin exponer su estructura.
type Iterador struct {
	actual  *Nodo
	primera bool
}

func (l *Lista) Iterador() *Iterador {
	return &Iterador{actual: l.cabeza, primera: true}
}

func (it *Iterador) Siguiente() bool {
	if it.actual == nil {
		return false
	}
	if !it.primera {
		it.actual = it.actual.sig
		if it.actual == nil {
			return false
		}
	}
	it.primera = false
	return true
}

func (it *Iterador) Valor() int {
	return it.actual.valor
}

func main() {
	lista := &Lista{}
	lista.AgregarAlFinal(1)
	lista.AgregarAlFinal(2)
	lista.AgregarAlFinal(3)

	it := lista.Iterador()
	for it.Siguiente() {
		fmt.Println(it.Valor())
	}
	// Salida:
	// 1
	// 2
	// 3
}
