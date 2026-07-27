package editor

import "errors"

var ErrSinOperaciones = errors.New("no hay operaciones para deshacer/rehacer")

type Editor struct {
	texto []rune
}

func NuevoEditor() *Editor {
	return &Editor{}
}

func (e *Editor) Contenido() string {
	return string(e.texto)
}

func (e *Editor) insertar(texto string, pos int) {
	runes := []rune(texto)
	if pos < 0 {
		pos = 0
	}
	if pos > len(e.texto) {
		pos = len(e.texto)
	}
	resultado := make([]rune, 0, len(e.texto)+len(runes))
	resultado = append(resultado, e.texto[:pos]...)
	resultado = append(resultado, runes...)
	resultado = append(resultado, e.texto[pos:]...)
	e.texto = resultado
}

func (e *Editor) borrar(pos, cant int) []rune {
	if pos < 0 || pos >= len(e.texto) || cant <= 0 {
		return nil
	}
	fin := pos + cant
	if fin > len(e.texto) {
		fin = len(e.texto)
	}
	borrado := make([]rune, fin-pos)
	copy(borrado, e.texto[pos:fin])
	resultado := make([]rune, 0, len(e.texto)-(fin-pos))
	resultado = append(resultado, e.texto[:pos]...)
	resultado = append(resultado, e.texto[fin:]...)
	e.texto = resultado
	return borrado
}

type Comando interface {
	Ejecutar(e *Editor)
	Deshacer(e *Editor)
}

type ComandoInsertar struct {
	texto string
	pos   int
}

func (c ComandoInsertar) Ejecutar(e *Editor) {
	e.insertar(c.texto, c.pos)
}

func (c ComandoInsertar) Deshacer(e *Editor) {
	e.borrar(c.pos, len([]rune(c.texto)))
}

type ComandoBorrar struct {
	pos     int
	borrado []rune
}

func (c *ComandoBorrar) Ejecutar(e *Editor) {
	c.borrado = e.borrar(c.pos, len(c.borrado))
}

func (c *ComandoBorrar) Deshacer(e *Editor) {
	e.insertar(string(c.borrado), c.pos)
}

type Historial struct {
	// TODO: implementar pilas de undo y redo, y referencia al editor
}

func NuevoHistorial(editor *Editor) *Historial {
	// TODO: implementar
	return &Historial{}
}

func (h *Historial) Ejecutar(c Comando) {
	// TODO: implementar
}

func (h *Historial) Deshacer() error {
	// TODO: implementar
	return nil
}

func (h *Historial) Rehacer() error {
	// TODO: implementar
	return nil
}
