# Coeficiente Binomial

**Función a implementar:** `CoeficienteBinomial(n, k int) int`

## El problema

Calcular C(n, k), la cantidad de subconjuntos de `k` elementos que se
pueden elegir de un conjunto con `n` elementos. Por ejemplo,
C(5, 2) = 10.

La fórmula clásica es C(n,k) = n! / (k! · (n-k)!), pero calcular
factoriales directamente crece rapidísimo: C(30, 15) = 155117520 (uno
de los casos del test). Con PD se construye el resultado celda por
celda, sin factoriales.

## Conceptos clave

- **Triángulo de Pascal:** cada valor se arma sumando los dos de la
  fila anterior:

  ```
  C(n, k) = C(n-1, k-1) + C(n-1, k)
  ```

  Los bordes de cada fila valen 1: C(n, 0) = 1 (no elegir nada) y
  C(n, n) = 1 (elegir todo).

- **Subestructura óptima y subproblemas superpuestos:** cada C(i, j)
  aparece muchas veces al construir las filas superiores; con PD se
  calcula una sola vez.

- **Estado:** `dp[i][j]` = C(i, j) para `0 <= i <= n` y `0 <= j <= k`.

## Estrategia: recurrencia (tabulación)

- Caso base: `dp[i][0] = 1` para toda fila, y `dp[i][i] = 1` (o sea
  `dp[i][j] = 1` cuando `j == 0` o `j == i`).
- Recurrencia: `dp[i][j] = dp[i-1][j-1] + dp[i-1][j]`.
- Llenado fila por fila, de `i = 0` hasta `n`; dentro de cada fila,
  de `j = 0` hasta `min(i, k)` (no hace falta calcular columnas más
  allá de `k`).
- Respuesta: `dp[n][k]`. Complejidad O(n × k).

Convención: si `k < 0` o `k > n`, no hay subconjuntos posibles → 0.

## Ejemplo a mano

C(5, 2): primeras filas del triángulo (columnas `j = 0, 1, 2`):

```
i=0: 1
i=1: 1  1
i=2: 1  2  1
i=3: 1  3  3
i=4: 1  4  6
i=5: 1  5 10
```

Detalle de la última celda:

- `dp[5][2] = dp[4][1] + dp[4][2] = 4 + 6 = 10`
  (la fórmula C(n,k) = C(n-1,k-1) + C(n-1,k)).

`CoeficienteBinomial(5, 2) = 10`, como espera el test. Cada celda se
calcula una sola vez: 6 × 3 = 18 celdas en este caso.

Casos borde:

- `C(10, 0) = 1` y `C(10, 10) = 1` (bordes del triángulo).
- `C(6, 3) = 20`.
- `C(30, 15) = 155117520` — con recursión naïve se recalcularía una
  cantidad enorme de veces; con PD sólo 30 × 16 celdas.
