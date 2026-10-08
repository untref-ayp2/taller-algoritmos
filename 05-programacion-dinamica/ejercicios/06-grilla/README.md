# Camino de Costo Mínimo en una Grilla

**Funciones a implementar:**

```go
func CaminoMinimo(grilla [][]int) int
func CaminoMinimoConRuta(grilla [][]int) (int, []Pos)
```

(El tipo ya está definido en el esqueleto:
`type Pos struct { Fila, Col int }`)

## El problema

Dada una grilla rectangular con un costo por celda, ir desde la celda
`(0,0)` hasta la esquina inferior derecha `(N-1, M-1)` moviéndose
**sólo hacia abajo o hacia la derecha**, minimizando la **suma** de los
costos de las celdas recorridas (la celda inicial se incluye).

- `CaminoMinimo` devuelve sólo el costo mínimo.
- `CaminoMinimoConRuta` devuelve el costo **y** el camino: un slice de
  `Pos` desde `(0,0)` hasta el destino.

## Conceptos clave

- **Subestructura óptima:** para llegar a `(i,j)` por el camino
  óptimo, el penúltimo paso vino de arriba `(i-1,j)` o de la
  izquierda `(i,j-1)`; y ambos tramos del camino también tienen que
  ser óptimos.

- **Estado:** `dp[i][j]` = costo mínimo para llegar a esa celda.

- **Reconstrucción de la solución:** `dp` guarda costos, no el camino.
  Para devolver el camino hay que **recorrer la tabla hacia atrás**,
  desde `(N-1, M-1)`: en cada paso se mira de qué vecino (arriba o
  izquierda) vino el mínimo y se avanza hacia ahí. Al llegar a `(0,0)`
  se invierte la lista para que quede en orden.

## Estrategia: recurrencia (tabulación)

- Caso base: `dp[0][0] = grilla[0][0]`.
- Primera fila (sólo se llega desde la izquierda):
  `dp[0][j] = dp[0][j-1] + grilla[0][j]`.
- Primera columna (sólo desde arriba):
  `dp[i][0] = dp[i-1][0] + grilla[i][0]`.
- Celda interior:
  `dp[i][j] = grilla[i][j] + min(dp[i-1][j], dp[i][j-1])`.

Llenado fila por fila, de arriba a la izquierda. Complejidad
O(N × M) en tiempo y en espacio.

## Ejemplo a mano

Grilla del test:

```
1 3 1
1 5 1
4 2 1
```

Tabla `dp` (costo mínimo para llegar a cada celda):

|         | col 0 | col 1      | col 2      |
|---------|-------|------------|------------|
| fila 0  | 1     | 1+3 = 4    | 4+1 = 5    |
| fila 1  | 1+1=2 | 5+min(4,2)=7 | 1+min(5,7)=6 |
| fila 2  | 2+4=6 | 2+min(7,6)=8 | 1+min(6,8)=**7** |

Costo mínimo: **7** (el camino 1→3→1→1→1).

Reconstrucción desde `(2,2)`:

| Celda | Arriba | Izquierda | De dónde vino |
|-------|--------|-----------|---------------|
| (2,2) | `dp[1][2] = 6` | `dp[2][1] = 8` | arriba → (1,2) |
| (1,2) | `dp[0][2] = 5` | `dp[1][1] = 7` | arriba → (0,2) |
| (0,2) | —      | `dp[0][1] = 4` | izquierda → (0,1) |
| (0,1) | —      | `dp[0][0] = 1` | izquierda → (0,0) |

Camino invertido: `(0,0) → (0,1) → (0,2) → (1,2) → (2,2)` — exactamente
la ruta que espera el test.

Casos borde:

- Grilla de una sola celda → el costo es el de esa celda y la ruta
  tiene una sola posición.
- Una sola fila → el camino recorre toda la fila hacia la derecha.
- En caso de empate entre arriba e izquierda, la implementación debe
  elegir de forma consistente (el test del ejemplo no tiene empates
  sobre el camino).
