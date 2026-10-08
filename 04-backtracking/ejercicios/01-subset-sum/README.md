# Subset Sum (subconjunto con suma exacta)

**Funciones a implementar:**

```go
func SubsetSum(nums []int, k int) []int
func TodasLasCombinaciones(nums []int, k int) [][]int
```

## El problema

Dado un conjunto de números enteros `nums` y un valor objetivo `k`,
encontrar un subconjunto cuya suma sea exactamente `k`.

- `SubsetSum` devuelve **un** subconjunto (un slice con los números
  elegidos) o `nil` si no existe.
- `TodasLasCombinaciones` devuelve **todos** los subconjuntos que suman
  `k`, uno por elemento del slice devuelto. Si no hay ninguno, un slice
  vacío.

Un subconjunto no necesita ser contiguo: de `[3, 7, 1, 8, 4]` con
`k = 10`, `[3, 7]` es una respuesta válida aunque el 3 y el 7 no sean
vecinos en el slice.

## Conceptos clave

- **Subconjunto:** cada número se incluye o se excluye; **no se
  repite**. De `[2, 3, 6, 7]`, la respuesta para `k = 7` es `[7]`:
  usar el 2 dos veces (`[2, 2, 3]`) no es un subconjunto, es un
  multiconjunto.

- **Combinación vs. permutación:** el orden no importa: `[3, 7]` y
  `[7, 3]` son el mismo subconjunto y no deben contarse dos veces.

- **Backtracking** (capítulo 4-4 del apunte): construir la solución
  paso a paso, descartando caminos inviables y *deshaciendo* (vuelta
  atrás) cuando una rama no funciona.

- **Poda:** si la suma parcial ya superó `k` (con números positivos),
  agregar más números sólo la puede empeorar: se corta la rama.

## Estrategia de backtracking

Elementos del esquema del apunte para este problema:

| Elemento del esquema | Definición en este ejercicio |
|---|---|
| `solucionParcial` | un `[]int` con los números elegidos hasta ahora |
| `esSolucion` | la suma del slice es exactamente `k` (o sea `target == 0`) |
| `extender` | probar los números de `nums` desde un índice `start` en adelante |
| `esFactible` | `target - nums[i] >= 0`, o sea que la suma no se pasa de `k` |
| `registrar` | `append(current, nums[i])` |
| `borrar` | `current = current[:len(current)-1]` |

La clave para no repetir permutaciones es el índice `start`: al
extender, sólo se miran los números desde `start` en adelante, y la
rama que **toma** el número `i` continúa desde `i + 1`. Así `[3, 7]`
se arma una sola vez (tomo el 3 del índice 0, después el 7 del índice
1) y nunca se arma `[7, 3]`.

Los dos niveles de la recursión según la función:

- `SubsetSum`: al encontrar la primera solución **corta** y la devuelve;
  si una rama no encuentra nada (`nil`), sigue probando con los demás
  números.
- `TodasLasCombinaciones`: al encontrar una solución la **acumula** en
  el resultado y sigue explorando las ramas que quedan, buscando todas.

Poda: `if target < 0 { return }` — si la suma parcial se pasó de `k`,
esa rama está muerta.

Dato: el problema también se puede resolver con programación dinámica
(capítulo 4-5: tiene subestructura óptima y subproblemas superpuestos),
pero acá se pide practicar backtracking.

## Ejemplo a mano

`nums = [2, 3, 5, 7]`, `k = 10`. Árbol de búsqueda: cada línea es una
opción agregada, y entre paréntesis va la suma parcial. Las hojas sin
más números disponibles terminan en vuelta atrás:

```
[] (0)
├── +2 → [2] (2)
│   ├── +3 → [2,3] (5)
│   │   ├── +5 → [2,3,5] (10) ✓ solución
│   │   └── +7 → [2,3,7] (12) ✗ poda (suma > 10)
│   ├── +5 → [2,5] (7)
│   │   └── +7 → [2,5,7] (14) ✗ poda
│   └── +7 → [2,7] (9) → no quedan números → vuelta atrás
├── +3 → [3] (3)
│   ├── +5 → [3,5] (8)
│   │   └── +7 → [3,5,7] (15) ✗ poda
│   └── +7 → [3,7] (10) ✓ solución
├── +5 → [5] (5)
│   └── +7 → [5,7] (12) ✗ poda
└── +7 → [7] (7) → no quedan números → vuelta atrás
```

- `SubsetSum([2,3,5,7], 10)` → `[2, 3, 5]` (la primera que encuentra).
- `TodasLasCombinaciones([2,3,5,7], 10)` → `[[2,3,5], [3,7]]` (en ese
  orden de exploración; los tests sólo validan las sumas, no el orden).

Casos borde:

- No existe solución (`k = 99` con `[3, 7, 1, 8, 4]`, cuya suma total
  es 23) → `SubsetSum` devuelve `nil` y `TodasLasCombinaciones` un
  slice vacío.
- Si `nums` viene vacío → no hay subconjuntos que probar: mismo
  resultado que el caso anterior.
