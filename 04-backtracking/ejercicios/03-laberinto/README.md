# Laberinto

**Función a implementar:** `Resolver(laberinto [][]int) []Posicion`

(El tipo ya está definido en el esqueleto:
`type Posicion struct { Fila, Col int }`)

## El problema

Dada una matriz `N × N` que representa un laberinto (0 = paso libre,
1 = obstáculo), encontrar un camino desde `(0,0)` hasta `(N-1, N-1)`
moviéndose **sólo hacia abajo y hacia la derecha**. La función devuelve
el camino como un slice de `Posicion`, o `nil` si no existe camino. No
modifica el laberinto.

El camino devuelto **incluye** tanto la posición inicial `(0,0)` como
la final `(N-1, N-1)`.

## Conceptos clave

- **Matriz como slice de slices:** la celda de la fila `f` y la
  columna `c` es `laberinto[f][c]`.

- **Sólo abajo y derecha:** como nunca se retrocede, un camino no
  puede tener ciclos: no hace falta llevar registro de celdas
  visitadas. Es una simplificación importante respecto de otros
  problemas de laberinto.

- **solucionParcial:** el camino recorrido hasta ahora (`[]Posicion`).

- **esSolucion:** la posición actual es `(N-1, N-1)`.

- **esFactible:** el movimiento llega a una celda dentro de la matriz,
  que además es 0 (no obstáculo).

## Estrategia de backtracking

| Elemento del esquema | Definición en este ejercicio |
|---|---|
| `solucionParcial` | un `[]Posicion` con las celdas del camino |
| `esSolucion` | `fila == N-1 && col == N-1` |
| `extender` | dos opciones: ir abajo (`fila+1`) o a la derecha (`col+1`) |
| `esFactible` | la celda está en el rango y `laberinto[fila][col] == 0` |
| `registrar` | agregar la posición actual al camino antes de recursar |
| `borrar` | la vuelta atrás es implícita: si la llamada recursiva no    encontró camino, la actual sigue sin su posición agregada |

Desde la celda actual se intenta primero **abajo** y después
**derecha** (el orden es una decisión de implementación: si una de las
dos direcciones llega al destino, se devuelve ese camino). Si ninguna
funciona, no hay camino desde esa celda y se devuelve `nil`.

Nota: el test no exige un camino en particular; sólo valida que el
camino empiece en `(0,0)`, termine en `(N-1, N-1)` y que cada paso sea
de distancia 1 hacia abajo o derecha sin pisar obstáculos.

## Ejemplo a mano

El laberinto del test:

```
0 0 0 1
1 1 0 1
1 1 0 0
1 1 1 0
```

Exploración (probando abajo primero):

| Paso | Posición | Intento | Resultado |
|------|----------|---------|-----------|
| 1 | (0,0) | abajo → (1,0) | obstáculo ✗ |
| 2 | (0,0) | derecha → (0,1) | libre ✓ |
| 3 | (0,1) | abajo → (1,1) | obstáculo ✗ |
| 4 | (0,1) | derecha → (0,2) | libre ✓ |
| 5 | (0,2) | abajo → (1,2) | libre ✓ |
| 6 | (1,2) | abajo → (2,2) | libre ✓ |
| 7 | (2,2) | abajo → (3,2) | obstáculo ✗ |
| 8 | (2,2) | derecha → (2,3) | libre ✓ |
| 9 | (2,3) | abajo → (3,3) | destino ✓ |

Camino devuelto: `(0,0), (0,1), (0,2), (1,2), (2,2), (2,3), (3,3)` —
7 posiciones.

Casos borde:

- Sin camino: el 2 × 2 con obstáculos en `(0,1)` y `(1,0)` devuelve
  `nil` (desde `(0,0)` ambas direcciones están bloqueadas).
- Tablero 1 × 1: `(0,0)` ya es el destino → camino de una posición.
