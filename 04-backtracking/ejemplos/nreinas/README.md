# Ejemplo: Problema de las N Reinas

Este ejemplo está desarrollado en detalle en el capítulo
**«Backtracking (Vuelta Atrás)»** del apunte, que incluye el enunciado,
el esquema general del algoritmo, el análisis de complejidad y la
implementación. Acá sólo se repasa qué hace el código.

## El problema

Colocar N reinas en un tablero N × N de manera que ninguna ataque a
otra: ni en la misma fila, ni en la misma columna, ni en ninguna
diagonal.

## Qué hace el código

- `NReinas(n int) []int`: devuelve un slice donde el **índice** es la
  fila y el **valor** es la columna de la reina de esa fila. Por
  ejemplo, `[2, 0, 3, 1]` para N = 4 (la misma representación que en
  el apunte).

- `esFactible(fila, columna, solucionParcial) bool`: es la
  `esFactible` del esquema. Chequea que no haya otra reina en la misma
  columna (valores repetidos en el slice) ni en las diagonales, usando
  que `fila - columna` es constante en una diagonal y `fila + columna`
  en la otra.

- `backtracking(n, fila, solucionParcial, solucion)`: la recursión
  general. Si `fila == n` ya hay una reina por fila: encontró la
  solución. Si no, prueba todas las columnas de la fila actual; para
  cada una que sea factible, agrega la reina (`registrar`) y continúa
  en la fila siguiente. Si ninguna columna funciona, la rama se corta
  (vuelta atrás).

## Cómo correrlo

```bash
go run ./04-backtracking/ejemplos/nreinas/
```

Imprime una solución válida para 8 reinas.
