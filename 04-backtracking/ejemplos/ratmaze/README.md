# Ejemplo: Laberinto (Rat in a Maze)

Resuelve el clásico problema del laberinto (*rat in a maze*) con
backtracking. El laberinto es una matriz de booleanos: `true` es una
celda libre y `false` una pared.

## Qué hace el código

- `NuevoRatMaze(maze [][]bool) *RatMaze`: crea el laberinto.

- `Resolver() bool`: busca un camino desde la celda superior izquierda
  `(0, 0)` hasta la inferior derecha y devuelve si encontró alguno. Las
  celdas que forman el camino quedan marcadas en `solucion`.

- `ResolverConPaso(paso) bool`: igual que `Resolver`, pero invoca
  `paso(x, y, avanzando)` en **cada** avance (`true`) y cada vuelta atrás
  (`false`). Así la visualización puede animar la búsqueda sin que el
  backtracking sepa nada de la terminal.

- `esSolucion(x, y) bool`: indica si la posición es el destino del
  laberinto.

- `esFactible(x, y) bool`: es la poda. Controla que la posición esté dentro
  del laberinto, que sea una celda libre y que no forme parte del camino
  actual.

- `registrar(x, y, paso)` marca la celda como parte del camino y avisa del
  avance; `borrar(x, y, paso)` la desmarca al volver atrás y avisa del
  retroceso.

- `backtracking(x, y, paso)`: la recursión general. Si llegó al destino
  (`esSolucion`) termina. Si no, registra la celda, prueba las cuatro
  direcciones (`extender`) y, si ninguna lleva a la salida, la borra
  (vuelta atrás) y devuelve `false`.

## Visualización

`ratmaze_visual.go` es la capa de dibujo, separada del backtracking:
`Tablero(actualX, actualY)` devuelve el laberinto en ASCII (`#` pared,
`.` libre, `o` camino, `x` celda explorada y abandonada, `@` cursor), y
`Animar(w, demora)` dibuja cada paso del recorrido con una pausa.

## Cómo correrlo

```bash
go run ./04-backtracking/ejemplos/ratmaze/
```

Anima un laberinto grande: se ve cómo la rata avanza, se mete en callejones
sin salida y vuelve atrás hasta encontrar el camino.

## Para experimentar

El laberinto del ejemplo está en `main.go`, en la variable `maze`: cada
`true` es una celda libre y cada `false` una pared. El origen es siempre la
celda de arriba a la izquierda `(0, 0)` y el destino la de abajo a la
derecha.

```go
maze := [][]bool{
	{true, true, true, /* ... */},
	// ...
}
```

Armá tu propio laberinto y correlo con
`go run ./04-backtracking/ejemplos/ratmaze/`. Algunas ideas:

- **Agregá callejones sin salida**: pasillos que terminan en una pared. Se
  ve cómo la rata se mete, se queda sin opciones y vuelve atrás.
- **Tapá la salida** (poné `false` en la esquina inferior derecha) para
  comprobar que el programa avisa que no hay camino.
- **Agrandalo**: cuantas más celdas libres, más pasos y más larga la
  animación (y también más lo que tarda).

Mantené todas las filas del mismo largo. El último argumento de `Animar`
(en `main.go`) es la pausa entre pasos.

## Tests

```bash
go test ./04-backtracking/ejemplos/ratmaze/
```

Verifican que la solución sea un camino simple y conexo desde el origen
hasta el destino, y el caso en que no existe camino. Agregá tus propios
laberintos y comprobá qué pasa.
