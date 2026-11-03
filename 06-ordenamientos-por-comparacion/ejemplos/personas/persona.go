// Package personas define el tipo Persona y funciones de comparación (cmp)
// para ordenar personas por distintos criterios.
//
// Los comparadores siguen el contrato de cmp.Compare: devuelven un entero
// negativo si a va antes que b, positivo si va después, y cero si son
// equivalentes. Así se pueden pasar directamente a MergeSort, QuickSort y
// HeapSort de la carpeta ejemplos/.
package personas

import "cmp"

// Persona representa a una persona con apellido, nombre, edad y número de DNI.
type Persona struct {
	Apellido string
	Nombre   string
	Edad     int
	DNI      int
}

// Ejemplo devuelve una lista de personas de muestra, usada tanto en la demo
// como en los tests de los ejercicios.
func Ejemplo() []Persona {
	return []Persona{
		{Apellido: "Garcia", Nombre: "Ana", Edad: 34, DNI: 29876543},
		{Apellido: "Alvarez", Nombre: "Bruno", Edad: 22, DNI: 41234567},
		{Apellido: "Fernandez", Nombre: "Carla", Edad: 41, DNI: 27123456},
		{Apellido: "Benitez", Nombre: "Diego", Edad: 22, DNI: 39876543},
		{Apellido: "Alvarez", Nombre: "Elena", Edad: 57, DNI: 18345678},
	}
}

// PorEdad compara a y b por edad ascendente (de menor a mayor).
func PorEdad(a, b Persona) int { return cmp.Compare(a.Edad, b.Edad) }

// PorEdadDesc compara a y b por edad descendente (de mayor a menor).
func PorEdadDesc(a, b Persona) int { return cmp.Compare(b.Edad, a.Edad) }

// PorDNI compara a y b por número de DNI ascendente.
func PorDNI(a, b Persona) int { return cmp.Compare(a.DNI, b.DNI) }

// PorDNIDesc compara a y b por número de DNI descendente.
func PorDNIDesc(a, b Persona) int { return cmp.Compare(b.DNI, a.DNI) }

// PorApellido compara a y b alfabéticamente por apellido (ascendente).
func PorApellido(a, b Persona) int { return cmp.Compare(a.Apellido, b.Apellido) }

// PorApellidoDesc compara a y b alfabéticamente por apellido (descendente).
func PorApellidoDesc(a, b Persona) int { return cmp.Compare(b.Apellido, a.Apellido) }
