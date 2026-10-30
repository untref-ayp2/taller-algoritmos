# Escaleras (formas de subir)

**Función a implementar:** `FormasDeSubir(n int) int`

## El problema

Una persona sube una escalera de `n` escalones saltando de a 1 o de a
2 escalones cada vez. Calcular cuántas **formas distintas** tiene de
llegar al escalón `n` partiendo del 0. Para `n = 3` hay 3 formas:
1+1+1, 1+2 y 2+1.

## Conceptos clave

- **Subestructura óptima** (capítulo 4-5 del apunte): para llegar al
  escalón `n`, el último salto fue de 1 o de 2:
  - si fue de 1, la persona venía del escalón `n-1`;
  - si fue de 2, venía del escalón `n-2`.

  Por lo tanto `formas(n) = formas(n-1) + formas(n-2)` — la misma
  recursión de la serie de Fibonacci que se vio en el apunte.

- **Subproblemas superpuestos:** la recursión naïve recalcularía
  `formas(i)` muchas veces (el árbol de llamadas de Fibonacci). PD
  calcula cada subproblema **una sola vez** y lo reutiliza.

- **Estado:** `dp[i]` = cantidad de formas de llegar al escalón `i`.

## Estrategia de programación dinámica (tabulación)

- Caso base: `dp[0] = 1` (una sola forma de no subir nada) y
  `dp[1] = 1` (un único salto de 1).
- Recurrencia: `dp[i] = dp[i-1] + dp[i-2]` para todo `i >= 2`.
- Llenado de abajo hacia arriba, de `i = 2` hasta `n`.

Complejidad: O(n) en tiempo y O(n) en espacio (con sólo dos variables
en vez del slice, O(1)).

También se puede hacer **top-down con memoización**: una función
recursiva `f(i)` que consulta `f(i-1)` y `f(i-2)`, guardando cada
resultado en un mapa antes de volver. Ambas variantes dan lo mismo.

## Ejemplo a mano

Tabla para `n = 6` (los pasos del llenado):

| i           | 0 | 1 | 2 | 3 | 4 | 5 | 6 |
|-------------|---|---|---|---|---|---|---|
| `dp[i]`     | 1 | 1 | 2 | 3 | 5 | 8 | 13 |

- `dp[2] = 1 + 1 = 2`
- `dp[3] = 1 + 2 = 3`
- `dp[4] = 2 + 3 = 5`
- `dp[5] = 3 + 5 = 8`
- `dp[6] = 5 + 8 = 13`

`FormasDeSubir(6) = 13`, como espera el test. Son los números de
Fibonacci: 1, 1, 2, 3, 5, 8, 13, …

Casos borde:

- `n = 0` → **1**: hay una forma de quedarse en el piso, no saltar
  nada. El test lo exige explícitamente.
- `n = 1` → 1; `n = 2` → 2 (1+1 o 2).
