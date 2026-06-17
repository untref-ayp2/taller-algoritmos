package iterador

type Nodo struct {
	Valor int
	sig   *Nodo
}

type Lista struct {
	cabeza *Nodo
}

func (l *Lista) Agregar(valor int) {
	// TODO: implementar
}

type Iterador struct {
	actual *Nodo
}

func (l *Lista) Iterador() *Iterador {
	// TODO: devolver un iterador nuevo apuntando al primer nodo
	return &Iterador{}
}

func (it *Iterador) Siguiente() bool {
	// TODO: avanzar al siguiente nodo; devolver true si hay elemento
	return false
}

func (it *Iterador) Valor() int {
	// TODO: devolver el valor del nodo actual
	return 0
}
