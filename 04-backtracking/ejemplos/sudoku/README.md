# Ejemplo: Sudoku

Resuelve un tablero de Sudoku 9 × 9 con backtracking. Las celdas vacías
se representan con `0`.

## Qué hace el código

- `NuevoSudoku(tablero [9][9]int) *Sudoku`: crea el tablero.

- `esSolucion() bool`: indica si no queda ninguna celda vacía.

- `esFactible(x, y, n) bool`: controla que el valor `n` no se repita en la
  fila, la columna ni el cuadrante de 3 × 3.

- `Resolver() bool`: busca **todas** las soluciones y las deja en
  `soluciones`. Devuelve si existe al menos una.

- `backtracking(paso, primera)`: la recursión general. Si `esSolucion`
  guarda una copia del tablero; si no, busca la próxima celda vacía y
  prueba los valores del 1 al 9 (`extender`). Si el valor es factible lo
  registra (`registrar`), resuelve el resto del tablero y al volver deshace
  la asignación (`borrar`). El callback `paso` permite observar cada
  asignación desde afuera sin que el algoritmo conozca la terminal.

## Visualización

`sudoku_visual.go` es la capa de dibujo, completamente separada del
backtracking:

- `Tablero(x, y, pasos) string`: dibuja el tablero con bordes, muestra
  en gris las celdas vacías y resalta la celda actual.

- `Animar(w, demora)`: resuelve mostrando cada paso con una pausa y
  termina mostrando la solución encontrada (o avisa si el tablero no tiene
  solución).

## Cómo correrlo

```bash
go run ./04-backtracking/ejemplos/sudoku/
```

Corre en la terminal una animación de la resolución paso a paso.

## Para experimentar

El tablero del ejemplo está en `main.go`, en la variable `tablero`. Las
celdas vacías se escriben con `0` y el resto son las pistas:

```go
tablero := [9][9]int{
	{4, 0, 0, 1, 9, 5, 0, 6, 8},
	// ...
}
```

Cambiá las pistas y volvé a correr
`go run ./04-backtracking/ejemplos/sudoku/`. Algunas ideas:

- **Poné más ceros**: quedan más celdas libres, más soluciones y muchos más
  pasos; se ve mejor cómo el algoritmo prueba, se arrepiente y vuelve atrás.
- **Armá un tablero sin salida**: la animación busca, no encuentra y avisa
  que el tablero no tiene solución.
- **Dejá el tablero casi resuelto**: la solución aparece en pocos pasos,
  útil para ver el final sin esperar.

El último argumento de `Animar` (en `main.go`) es la pausa entre pasos:
bajala para que sea más rápido o subila para mirarlo con calma.

## Tests

```bash
go test ./04-backtracking/ejemplos/sudoku/
```

Verifican que las soluciones respeten las reglas del sudoku y las pistas
originales, con tableros de una, dos y muchísimas soluciones (¡32 604!).
Agregá tus propios tableros y averiguá cuántas soluciones tienen.
