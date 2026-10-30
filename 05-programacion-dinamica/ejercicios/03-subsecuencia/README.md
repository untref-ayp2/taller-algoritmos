# Subsecuencia Común Más Larga (LCS)

**Función a implementar:** `LCS(a, b string) int`

## El problema

Dadas dos cadenas `a` y `b`, devolver la **longitud** de la subsecuencia
común más larga: la cadena más larga que aparece en ambas, respetando el
orden de los caracteres pero sin exigir que sean contiguos. Por ejemplo,
en `"AGGTAB"` y `"GXTXAYB"` la respuesta es 4 (`"GTAB"`).

## Conceptos clave

- **Subsecuencia vs. subcadena:** una subsecuencia respeta el orden
  pero puede saltarse caracteres; una subcadena es un trozo continuo.
  `"ace"` es subsecuencia de `"abcdef"` (saltea `b` y `d`), pero
  `"adf"` no lo es (rompe el orden). En cambio `"cde"` es subcadena de
  `"abcdef"`.

- **Subestructura óptima:** al comparar los últimos caracteres de las
  dos cadenas pasan dos cosas: o son iguales y se suman a la solución
  de los restos, o no lo son y conviene descartar un carácter de una
  de las dos cadenas (quedándose con lo mejor).

- **Subproblemas superpuestos:** las parejas `(prefijo de a, prefijo
  de b)` se repiten a lo largo del cálculo.

- **Estado:** `dp[i][j]` = longitud de la LCS de los primeros `i`
  caracteres de `a` y los primeros `j` de `b`.

## Estrategia: recurrencia (tabulación)

Tabla de `(m+1) × (n+1)` donde `m = len(a)` y `n = len(b)`, con una
fila y una columna de ceros: con una cadena vacía no hay subsecuencia
común (ese es el caso base).

Para `i >= 1` y `j >= 1`:

- si `a[i-1] == b[j-1]` → `dp[i][j] = dp[i-1][j-1] + 1`
  (el carácter emparejado se suma a la LCS de los restos);
- si no → `dp[i][j] = max(dp[i-1][j], dp[i][j-1])`
  (se descarta un carácter de `a` o de `b`, eligiendo el mejor
  resultado).

Llenado fila por fila, de arriba a la izquierda. Complejidad
O(m × n) en tiempo y en espacio. La respuesta es `dp[m][n]`.

## Ejemplo a mano

`a = "abcdef"`, `b = "acf"` (un caso del test). Filas = prefijos de
`a`, columnas = prefijos de `b`:

|       | ""  | a   | c   | f   |
|-------|-----|-----|-----|-----|
| ""    | 0   | 0   | 0   | 0   |
| a     | 0   | 1   | 1   | 1   |
| b     | 0   | 1   | 1   | 1   |
| c     | 0   | 1   | 2   | 2   |
| d     | 0   | 1   | 2   | 2   |
| e     | 0   | 1   | 2   | 2   |
| f     | 0   | 1   | 2   | **3** |

Algunas celdas:

- `dp[1][1]`: `a == a` → `dp[0][0] + 1 = 1`.
- `dp[3][2]`: `c == c` → `dp[2][1] + 1 = 1 + 1 = 2`.
- `dp[4][2]`: `d ≠ c` → `max(dp[3][2], dp[4][1]) = max(2, 1) = 2`.
- `dp[6][3]`: `f == f` → `dp[5][2] + 1 = 2 + 1 = 3`.

`LCS("abcdef", "acf") = 3` — la subsecuencia `"acf"` (que además es
subcadena, pero no hace falta que lo sea).

Casos borde:

- Cadenas vacías (`""` y `""`, o `"abc"` y `""`) → 0.
- Sin caracteres en común (`"abc"` y `"def"`) → 0.
- Cadenas idénticas → su longitud (`"abc"` y `"abc"` → 3).
