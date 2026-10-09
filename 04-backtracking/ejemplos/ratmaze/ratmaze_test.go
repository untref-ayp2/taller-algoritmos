package main

import (
	"strings"
	"testing"
)

func TestRatMazeEncuentraCamino(t *testing.T) {
	maze := [][]bool{
		{true, true, false, true, true, true, true},
		{true, true, false, true, false, false, true},
		{true, true, false, true, true, false, true},
		{true, true, false, false, true, false, true},
		{true, true, true, true, true, false, true},
	}

	ratmaze := NuevoRatMaze(maze)

	if !ratmaze.Resolver() {
		t.Fatal("Resolver() = false, se esperaba encontrar un camino")
	}
	if !esCaminoValido(maze, ratmaze.solucion) {
		t.Fatalf("la solución no es un camino válido:\n%s", ratmaze.Tablero(-1, -1))
	}
}

func TestRatMazeSinCamino(t *testing.T) {
	maze := [][]bool{
		{true, false},
		{false, false},
	}

	ratmaze := NuevoRatMaze(maze)

	if ratmaze.Resolver() {
		t.Fatal("Resolver() = true, no existe camino hasta el destino")
	}
}

func TestRatMazeRegistraVueltaAtras(t *testing.T) {
	maze := [][]bool{
		{true, true, false},
		{true, false, false},
		{true, true, true},
	}

	ratmaze := NuevoRatMaze(maze)

	huboVueltaAtras := false
	ratmaze.ResolverConPaso(func(x, y int, avanzando bool) {
		if !avanzando {
			huboVueltaAtras = true
		}
	})

	if !huboVueltaAtras {
		t.Error("se esperaba observar al menos una vuelta atrás")
	}
}

func TestRatMazeTablero(t *testing.T) {
	maze := [][]bool{
		{true, false},
		{true, true},
	}

	ratmaze := NuevoRatMaze(maze)
	ratmaze.Resolver()

	tablero := ratmaze.Tablero(-1, -1)
	if !strings.Contains(tablero, "#") {
		t.Errorf("el tablero debería dibujar las paredes con '#':\n%s", tablero)
	}
	if !strings.Contains(tablero, "o") {
		t.Errorf("el tablero debería dibujar el camino con 'o':\n%s", tablero)
	}
}

func TestAnimarRatMaze(t *testing.T) {
	maze := [][]bool{
		{true, true, false},
		{true, false, false},
		{true, true, true},
	}

	ratmaze := NuevoRatMaze(maze)

	var salida strings.Builder
	ratmaze.Animar(&salida, 0)

	if !strings.Contains(salida.String(), "Pasos") {
		t.Errorf("la animación debería mostrar el contador de pasos")
	}
}

// esCaminoValido comprueba que solucion sea un camino simple desde (0,0)
// hasta el extremo inferior derecho, usando sólo celdas libres del laberinto.
func esCaminoValido(maze, solucion [][]bool) bool {
	if len(maze) == 0 || len(solucion) != len(maze) {
		return false
	}
	filas := len(maze)
	columnas := len(maze[0])
	for y := 0; y < filas; y++ {
		if len(maze[y]) != columnas || len(solucion[y]) != columnas {
			return false
		}
	}

	// Sólo celdas libres pueden formar parte de la solución.
	for y := 0; y < filas; y++ {
		for x := 0; x < columnas; x++ {
			if solucion[y][x] && !maze[y][x] {
				return false
			}
		}
	}

	// El camino empieza en (0,0) y termina en el extremo inferior derecho.
	if !solucion[0][0] || !solucion[filas-1][columnas-1] {
		return false
	}

	// Cada celda del camino tiene grado 2, salvo los extremos (grado 1):
	// así se comprueba que es un camino simple y conexo.
	for y := 0; y < filas; y++ {
		for x := 0; x < columnas; x++ {
			if !solucion[y][x] {
				continue
			}
			grado := 0
			if y > 0 && solucion[y-1][x] {
				grado++
			}
			if y < filas-1 && solucion[y+1][x] {
				grado++
			}
			if x > 0 && solucion[y][x-1] {
				grado++
			}
			if x < columnas-1 && solucion[y][x+1] {
				grado++
			}
			esExtremo := (y == 0 && x == 0) || (y == filas-1 && x == columnas-1)
			if esExtremo && grado != 1 {
				return false
			}
			if !esExtremo && grado != 2 {
				return false
			}
		}
	}
	return true
}
