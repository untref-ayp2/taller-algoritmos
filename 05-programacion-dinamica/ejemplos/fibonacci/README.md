# Ejemplo: Fibonacci con Programación Dinámica

El problema de Fibonacci —incluido el árbol de llamadas del algoritmo
recursivo naïve y las versiones con tabulación y memoización— está
desarrollado en el capítulo **«Programación Dinámica»** del apunte.
Este directorio contiene las dos versiones eficientes con sus tests.

## Qué hace el código

- `FibonacciTab(n int) int`: **tabulación** (bottom-up). Llena un
  slice `dp` de 0 a n con `dp[i] = dp[i-1] + dp[i-2]`. O(n) en tiempo,
  O(n) en espacio.

- `FibonacciMemo(n int) int`: **memoización** (top-down). Versión
  recursiva con un mapa que guarda los valores ya calculados; si un
  subproblema ya está en el mapa se recupera sin recalcular. O(n) en
  tiempo (y O(n) de pila + caché).

Ambas funciones devuelven lo mismo; la versión naïve sin caché es
exponencial y está explicada en el apunte.

## Cómo correrlo

```bash
go test ./05-programacion-dinamica/ejemplos/fibonacci/
```
