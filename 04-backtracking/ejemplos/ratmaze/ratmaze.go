package main

// Este archivo contiene el backtracking del laberinto de la rata.
//
// Arriba está el núcleo del algoritmo (el estado de la búsqueda y la
// recursión); al final, los puntos de entrada.

// ---------------------------------------------------------------------------
// Backtracking
// ---------------------------------------------------------------------------

// RatMaze resuelve el laberinto de la rata: encontrar un camino desde la
// celda superior izquierda hasta la inferior derecha, moviéndose en las
// cuatro direcciones por las celdas libres (true).
type RatMaze struct {
	maze     [][]bool
	solucion [][]bool
	visitado [][]bool
	destX    int
	destY    int
}

// PasoLaberinto se invoca en cada cambio del camino: avanzando es true al
// entrar a una celda y false al retroceder (vuelta atrás). Es el gancho que
// usa la capa de visualización para animar la búsqueda; puede ser nil.
type PasoLaberinto func(x, y int, avanzando bool)

// esSolucion indica si (x, y) es el destino del laberinto (y es transitable).
func (r *RatMaze) esSolucion(x, y int) bool {
	return x == r.destX && y == r.destY && r.maze[y][x]
}

// esFactible indica si se puede avanzar a (x, y): está dentro del laberinto,
// es una celda libre y no forma parte del camino actual.
func (r *RatMaze) esFactible(x, y int) bool {
	return y >= 0 && y < len(r.maze) &&
		x >= 0 && x < len(r.maze[y]) &&
		r.maze[y][x] &&
		!r.solucion[y][x]
}

// registrar marca (x, y) como parte del camino y avisa del avance.
func (r *RatMaze) registrar(x, y int, paso PasoLaberinto) {
	r.solucion[y][x] = true
	r.visitado[y][x] = true
	if paso != nil {
		paso(x, y, true)
	}
}

// borrar desmarca (x, y) al volver atrás y avisa del retroceso.
func (r *RatMaze) borrar(x, y int, paso PasoLaberinto) {
	r.solucion[y][x] = false
	if paso != nil {
		paso(x, y, false)
	}
}

// backtracking es la recursión general del esquema: marca la celda actual y
// prueba las cuatro direcciones hasta llegar al destino.
func (r *RatMaze) backtracking(x, y int, paso PasoLaberinto) bool {
	// esSolucion: llegamos al destino.
	if r.esSolucion(x, y) {
		r.registrar(x, y, paso)
		return true
	}

	// Poda: celda fuera del laberinto, con pared, o ya incluida en el camino.
	if !r.esFactible(x, y) {
		return false
	}

	// registrar: marco la celda como parte del camino.
	r.registrar(x, y, paso)

	// extender: pruebo las cuatro direcciones.
	if r.backtracking(x+1, y, paso) { // derecha
		return true
	}
	if r.backtracking(x, y+1, paso) { // abajo
		return true
	}
	if r.backtracking(x-1, y, paso) { // izquierda
		return true
	}
	if r.backtracking(x, y-1, paso) { // arriba
		return true
	}

	// borrar (vuelta atrás): la celda no lleva a ninguna solución.
	r.borrar(x, y, paso)
	return false
}

// ---------------------------------------------------------------------------
// Puntos de entrada
// ---------------------------------------------------------------------------

func NuevoRatMaze(maze [][]bool) *RatMaze {
	return &RatMaze{maze: maze}
}

// Resolver busca un camino y deja marcadas en solucion las celdas que lo
// forman. Devuelve true si encontró alguno.
func (r *RatMaze) Resolver() bool {
	return r.ResolverConPaso(nil)
}

// ResolverConPaso busca un camino y, si paso no es nil, lo invoca en cada
// avance y cada vuelta atrás. Así la capa de visualización puede animar la
// búsqueda sin que el backtracking conozca la terminal.
func (r *RatMaze) ResolverConPaso(paso PasoLaberinto) bool {
	r.solucion = make([][]bool, len(r.maze))
	r.visitado = make([][]bool, len(r.maze))
	for i := range r.maze {
		r.solucion[i] = make([]bool, len(r.maze[i]))
		r.visitado[i] = make([]bool, len(r.maze[i]))
	}
	r.destY = len(r.maze) - 1
	r.destX = len(r.maze[r.destY]) - 1
	return r.backtracking(0, 0, paso)
}
