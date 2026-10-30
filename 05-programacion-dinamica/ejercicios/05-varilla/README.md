# Corte de Varilla (Rod Cutting)

**Función a implementar:** `CorteVarilla(n int, precios []int) int`

## El problema

Se tiene una varilla de longitud `n` y una lista de precios donde
`precios[i]` es el precio de una varilla de longitud `i + 1` (o sea,
`precios[0]` es el precio de la longitud 1, `precios[1]` el de la
longitud 2, y así sucesivamente). La varilla se puede cortar en
trozos y vender cada trozo a su precio. Calcular la **ganancia
óptima** posible (la función devuelve sólo la ganancia, no los cortes).

## Conceptos clave

- **Subestructura óptima:** el primer corte parte el problema en dos:
  si el primer trozo tiene longitud `j`, la ganancia es
  `precios[j-1]` (ese trozo) más la ganancia óptima de la varilla que
  queda (`n - j`). La parte restante también debe resolverse de forma
  óptima.

- **Subproblemas superpuestos:** las ganancias óptimas de longitudes
  menores se necesitan una y otra vez para cada primer corte posible.

- **Estado:** `dp[i]` = ganancia óptima para una varilla de longitud
  `i`.

## Estrategia: recurrencia (tabulación)

- Caso base: `dp[0] = 0` (varilla de longitud 0 no rinde nada).
- Recurrencia: para cada longitud `i`, probar todos los primeros
  cortes `j` de 1 a `i`:

  ```
  dp[i] = max(precios[j-1] + dp[i-j])   para j = 1 .. i
  ```

  (si la tabla `precios` no cubre esa longitud, sólo se consideran
  los tamaños con precio disponible).

- Llenado de `i = 1` hasta `n`. Complejidad O(n²) en tiempo, O(n) en
  espacio.

## Ejemplo a mano

`n = 4`, `precios = [1, 5, 8, 9]` (caso del test):

| i | Primer corte j y ganancia                    | `dp[i]` |
|---|----------------------------------------------|---------|
| 1 | j=1: 1 + dp[0] = 1                           | 1       |
| 2 | j=1: 1+dp[1]=2 · j=2: 5+dp[0]=5              | 5       |
| 3 | j=1: 1+5=6 · j=2: 5+1=6 · j=3: 8+0=8         | 8       |
| 4 | j=1: 1+8=9 · j=2: 5+5=**10** · j=3: 8+1=9 · j=4: 9+0=9 | **10** |

Óptimo: cortar en dos trozos de longitud 2 (5 + 5 = 10). Vender la
varilla entera da 9 y el corte 3+1 da 9: no siempre «parecer» bueno
es lo mejor — de eso se trata PD.

Para `n = 8` con `precios = [1, 5, 8, 9, 10, 17, 17, 20]` el test
espera 22 (por ejemplo, 6+2 → 17 + 5).

Casos borde:

- `n = 0` → 0 (exigido por el test).
- Una sola longitud con precio → cortar o vender entera, lo que rinda
  más.
