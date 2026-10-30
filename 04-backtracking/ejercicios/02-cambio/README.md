# Cambio de Monedas (todas las combinaciones)

**Función a implementar:** `FormasDeCambiar(monto int, denominaciones []int) [][]int`

## El problema

Dado un monto y un conjunto de denominaciones de monedas, encontrar
**todas** las combinaciones posibles de monedas que suman exactamente
el monto. Cada denominación se puede usar las veces que haga falta.

## Conceptos clave

- Este problema ya se vio en el capítulo 4-3 del apunte (y en el
  ejercicio `03-algoritmos-avidos/02-maquina-expendedora/`) con otro
  objetivo: ahí se pedía **una sola** solución, la que usa la **menor
  cantidad** de monedas, resuelta de forma ávida. Acá se piden
  **todas** las combinaciones posibles: para eso el greedy no alcanza
  y hace falta backtracking.

- **Combinación con repetición:** las monedas de una misma
  denominación se pueden repetir sin límite (`{2, 2, 1}` es válido
  para monto 5), pero el orden no importa: `{2, 1, 2}` es la misma
  combinación que `{2, 2, 1}` y no debe contarse dos veces.

## Estrategia de backtracking

| Elemento del esquema | Definición en este ejercicio |
|---|---|
| `solucionParcial` | un `[]int` con las monedas elegidas hasta ahora |
| `esSolucion` | la suma de las monedas es exactamente `monto` |
| `extender` | probar las denominaciones desde un índice `start` en adelante |
| `esFactible` | `monto - denominaciones[i] >= 0`, o sea que el resto no queda negativo |
| `registrar` | `append(current, denominaciones[i])` |
| `borrar` | `current = current[:len(current)-1]` |

El índice `start` evita permutar: la rama que usa la denominación `i`
sigue mirando **desde** `i` en adelante y nunca vuelve a una
denominación anterior. A diferencia de subset sum sigue desde `i` y no
desde `i + 1`, porque la misma denominación se puede repetir. Cada
multiconjunto se arma de una única forma.

Poda: si el resto se puso negativo, la rama se corta.

Con `denominaciones = [1, 2, 5]` y `monto = 5` se obtienen las cuatro
formas posibles: `{5}`, `{2,2,1}`, `{2,1,1,1}` y `{1,1,1,1,1}`.

Casos borde:

- Imposible: `denominaciones = [3, 5]`, `monto = 7` → ningún entero
  `3a + 5b` da 7, así que devuelve un slice vacío.
- `monto = 0` → no hay monedas que elegir: slice vacío.
