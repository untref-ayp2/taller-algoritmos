package main

import "fmt"

// Componente define la operación común a archivos y carpetas.
type Componente interface {
	Tamanio() int64
}

// Archivo representa un elemento simple (hoja) del sistema.
type Archivo struct {
	nombre string
	bytes  int64
}

func (a *Archivo) Tamanio() int64 {
	return a.bytes
}

// Carpeta representa un elemento compuesto que contiene otros componentes.
type Carpeta struct {
	nombre      string
	componentes []Componente
}

func (c *Carpeta) Tamanio() int64 {
	var total int64
	for _, comp := range c.componentes {
		total += comp.Tamanio()
	}
	return total
}

func (c *Carpeta) Agregar(comp Componente) {
	c.componentes = append(c.componentes, comp)
}

func main() {
	readme := &Archivo{nombre: "README.md", bytes: 2048}
	licencia := &Archivo{nombre: "LICENSE", bytes: 1024}

	src := &Carpeta{nombre: "src"}
	src.Agregar(&Archivo{nombre: "main.go", bytes: 4096})
	src.Agregar(&Archivo{nombre: "utils.go", bytes: 1536})

	docs := &Carpeta{nombre: "docs"}
	docs.Agregar(&Archivo{nombre: "manual.pdf", bytes: 524288})

	proyecto := &Carpeta{nombre: "mi-proyecto"}
	proyecto.Agregar(readme)
	proyecto.Agregar(licencia)
	proyecto.Agregar(src)
	proyecto.Agregar(docs)

	fmt.Println(proyecto.Tamanio())
	// Salida: 532992
}
