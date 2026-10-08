# Sudoku

**Función a implementar:** `ResolverSudoku(tablero [][]int) bool`

## El problema

Completar un tablero de sudoku 9 × 9. Cada celda contiene un dígito
de 1 a 9, o 0 si está vacía. La función modifica el tablero **in
place** y devuelve `true` si encontró una solución, `false` si es
imposible.

Reglas del sudoku: cada **fila**, cada **columna** y cada **caja 3 × 3**
debe tener los dígitos 1 a 9 sin repetirse.

## Conceptos clave

- **solucionParcial:** el tablero mismo, que se va escribiendo in
  place. La "opción" de cada paso es poner un número en la primera
  celda vacía.

- **esSolucion:** no quedan celdas con 0 (el tablero está completo).

- **extender:** para la celda vacía actual, probar los números de 1 a
  9 en orden.

- **esFactible:** el número no aparece en la fila, ni en la columna,
  ni en la caja 3 × 3 de esa celda. Es la `esFactible` del esquema del
  apunte (en la solución se llama `esValidoUbicacion`).

- **registrar:** escribir `tablero[fila][col] = num`.

- **borrar:** la vuelta atrás es `tablero[fila][col] = 0`, para dejar
  la celda como estaba antes de seguir probando.

- **Poda:** la factibilidad se chequea **antes** de escribir y recursar:
  un número que viola alguna regla se descarta sin explorar su rama.
  Sin esta poda, el árbol de búsqueda sería enorme.

## Estrategia de backtracking

1. Recorrer el tablero fila por fila y buscar la primera celda con 0.
2. Si no hay ninguna → tablero completo → devolver `true`.
3. Para cada número de 1 a 9: si es factible en esa celda, escribirlo
   y llamar a la función recursivamente.
   - Si la llamada devolvió `true` → propagar `true` (ya está resuelto).
   - Si devolvió `false` → **deshacer** (volver a poner 0) y probar el
     siguiente número.
4. Si ningún número funcionó → devolver `false` (ninguna opción
   extiende esta solución parcial).

## Ejemplo a mano

Primeras filas del tablero del test:

```
5 3 0 0 7 0 0 0 0
6 0 0 1 9 5 0 0 0
0 9 8 0 0 0 0 6 0
...
```

La primera celda vacía (recorrido fila por fila) es `(0, 2)`.
Calculemos sus candidatos:

| Fuente | Valores presentes |
|---|---|
| Fila 0 | 5, 3, 7 |
| Columna 2 | 8 |
| Caja superior-izquierda (filas 0-2, columnas 0-2) | 5, 3, 6, 9, 8 |

Prohibidos: 3, 5, 6, 7, 8, 9 → candidatos para `(0, 2)`: **1, 2, 4**.

Se prueba 1 (es factible) y se escribe; luego se busca la siguiente
celda vacía, `(0, 3)`, y se repite. Si en algún punto la rama se atasca
(la celda actual no admite ningún número), se vuelve atrás: se deshace
lo último escrito y se prueba el siguiente candidato. Por eso en este
punto puede terminar cambiando el 1 por un 2 o por un 4 más adelante:
la función prueba y deshace hasta que el tablero se complete o se
agoten todas las posibilidades.

El test imposible tiene dos 5 en la primera fila: es un conflicto que
ninguna completación puede arreglar, así que la búsqueda se agota y
devuelve `false`. El test de tablero completo y válido no tiene celdas
con 0: la función devuelve `true` sin escribir nada.

Casos borde:

- Tablero completo y válido → `true` (no hay nada que completar).
- Tablero con un conflicto ya presente → `false`.
- Las celdas que ya traían números no deben modificarse.
