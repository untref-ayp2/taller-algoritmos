package main

import "fmt"

// Medible es la interfaz que espera el cliente.
type Medible interface {
	Medir() float64
}

// Robot es el Adaptado: mide en metros y centímetros.
type Robot struct {
	metros      int
	centimetros int
}

func (r *Robot) Medir() (int, int) {
	return r.metros, r.centimetros
}

// RobotAdapter envuelve un Robot y adapta su interfaz a Medible.
type RobotAdapter struct {
	adaptado *Robot
}

func (a *RobotAdapter) Medir() float64 {
	metros, centimetros := a.adaptado.Medir()
	totalCentimetros := (metros * 100) + centimetros
	pulgadas := float64(totalCentimetros) / 2.54
	return pulgadas
}

func main() {
	robot := &Robot{metros: 10, centimetros: 50}
	adaptado := &RobotAdapter{adaptado: robot}
	distancia := adaptado.Medir()
	fmt.Printf("Distancia: %.2f pulgadas\n", distancia)
	// Salida: Distancia: 413.39 pulgadas
}
