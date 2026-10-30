# Cambio de Monedas con Programación Dinámica

**Funciones a implementar:**

```go
func CambioTab(monto int, denominaciones []int) int
func CambioMemo(monto int, denominaciones []int) int
```

## El problema

Dado un monto y un conjunto de denominaciones de monedas, devolver la
**cantidad mínima** de monedas necesaria para formar ese monto, o `-1`
si no es posible. Hay que implementarlo dos veces: con **tabulación**
(bottom-up) y con **memoización** (top-down).

## Conceptos clave

- El mismo problema se vio antes con otros objetivos: de forma
  **ávida** en el capítulo 4-3 (funciona sólo con denominaciones
  canónicas) y generando **todas** las combinaciones con backtracking
  en `04-backtracking/02-cambio/`. Acá la pregunta es otra: la
  **cantidad mínima**, para **cualquier** conjunto de denominaciones.

- Por ejemplo, con `{1, 3, 4}` y monto 6:
  - el ávido toma la 4 primero → 4+1+1 = **3** monedas;
  - el óptimo es 3+3 = **2** monedas.

  PD siempre encuentra el óptimo, sin pedir condiciones a las
  denominaciones.

- **Estado:** `dp[i]` = cantidad mínima de monedas para formar el
  monto `i`.

- **Subestructura óptima:** para formar el monto `i`, la última moneda
  usada fue alguna denominación `d <= i`, y el resto (`i - d`) se pagó
  también de forma óptima.

## Estrategia: recurrencia

**Tabulación** (`CambioTab`):

- Caso base: `dp[0] = 0` (monto 0 no necesita monedas).
- Inicializar el resto de `dp[i]` con un valor «infinito» (p. ej.
  `monto + 1`, que no alcanza ninguna suma real).
- Recurrencia: `dp[i] = min(dp[i-d] + 1)` recorriendo cada
  denominación `d` con `d <= i`.
- Llenado de `i = 1` hasta `monto`.
- Si al final `dp[monto]` sigue «infinito» → `-1`.

**Memoización** (`CambioMemo`): la misma recurrencia, pero como
función recursiva `f(m) = 1 + min(f(m-d))` que guarda cada `f(m)` en
un mapa antes de volver. Si un subproblema ya está en el mapa, se
recupera sin recalcular.

El test verifica además que ambas variantes coincidan monto por monto
(de 0 a 30).

## Ejemplo a mano

`monto = 6`, `denominaciones = {1, 3, 4}`:

| i | 0 | 1 | 2 | 3 | 4 | 5 | 6 |
|---|---|---|---|---|---|---|---|
| `dp[i]` | 0 | 1 | 2 | 1 | 1 | 2 | 2 |

Pasos:

- `i=1`: sólo entra la 1 → `dp[1] = dp[0]+1 = 1`.
- `i=2`: sólo la 1 → `dp[2] = dp[1]+1 = 2`.
- `i=3`: con 1 → `dp[2]+1 = 3`; con 3 → `dp[0]+1 = 1` → mínimo 1.
- `i=4`: con 1 → 2; con 3 → 2; con 4 → `dp[0]+1 = 1` → mínimo 1.
- `i=5`: con 1 → 2; con 3 → 3; con 4 → 2 → mínimo 2.
- `i=6`: con 1 → 3; con 3 → `dp[3]+1 = 2`; con 4 → 3 → mínimo 2.

`CambioTab(6, {1,3,4}) = 2` (3+3), contra los 3 monedas del ávido.

Casos borde:

- `monto = 0` → 0 (no hace falta ninguna moneda).
- Imposible: monto 7 con `{5, 10}` → ninguna celda se alcanza con esas
  monedas → `-1`.
- `{2}` con monto 10 → 5 (2+2+2+2+2); monto 9 con `{5, 2, 1}` → 3
  (5+2+2): las denominaciones **no** necesitan venir ordenadas.
