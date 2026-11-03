# Ejemplos: Ordenamientos por Comparación

Ejemplos funcionales de los tres ordenamientos por comparación vistos en el
capítulo, implementados de forma **genérica** con una función de comparación
`cmp`.

## Cómo correr la demo

```bash
go run ./06-ordenamientos-por-comparacion/ejemplos/
```

## Qué hay acá

| Carpeta | Contenido |
|---|---|
| `mergesort/` | `MergeSort[T any](s []T, cmp func(a, b T) int)` y su fusión. |
| `quicksort/` | `QuickSort[T any](s []T, cmp func(a, b T) int)` y su partición. |
| `heapsort/` | `HeapSort[T any](s []T, cmp func(a, b T) int)` y su hundimiento. |
| `personas/` | El tipo `Persona` y comparadores para ordenar por edad, DNI y apellido. |
| `main.go` | Demo que ordena personas por distintos criterios. |

## La función `cmp`

Los tres algoritmos reciben una función `cmp func(a, b T) int` que sigue el
contrato de `cmp.Compare` de la biblioteca estándar:

- devuelve un entero **negativo** si `a` va antes que `b`;
- **cero** si son equivalentes;
- **positivo** si `a` va después que `b`.

Con ese contrato, el algoritmo no necesita saber nada del tipo `T`: solo le
pregunta a `cmp` cómo se comparan dos elementos.

En los **ejercicios** se pide adaptar estos algoritmos para que usen, en lugar
de `cmp`, una función `less func(a, b T) bool`. Las ventajas y desventajas de
`less` se explican en `ejercicios/README.md`.

## Tests

```bash
go test ./06-ordenamientos-por-comparacion/ejemplos/...
```
