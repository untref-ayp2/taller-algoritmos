# Minimizar Tiempo de Espera (SJF)

**Función a implementar:** `OrdenOptimo(tiempos []int) []int`

## El problema

Hay `n` trabajos que llegan todos juntos (al tiempo 0) y una sola máquina
que los procesa de a uno, en orden, sin interrupciones. El trabajo `i` tarda
`tiempos[i]` unidades de tiempo en procesarse. Hay que elegir **en qué
orden** ejecutarlos de manera de minimizar la suma de los tiempos de espera
de todos los trabajos. La función devuelve los **índices** de los trabajos
en el orden óptimo (no los tiempos): para `tiempos = [5, 1, 3]` devuelve
`[1, 2, 0]`, es decir, primero el trabajo 1 (tarda 1), luego el 2 (tarda 3)
y al final el 0 (tarda 5).

## Conceptos clave

- **Trabajo:** se identifica por su **índice** en el slice de entrada y por
  su tiempo de procesamiento. Puede haber trabajos con el mismo tiempo, por
  eso el resultado son índices y no tiempos.

- **Tiempo de espera de un trabajo:** desde que empieza la jornada hasta
  que *arranca* a procesarse. Como todos llegan juntos al tiempo 0, es
  simplemente la suma de los tiempos de los trabajos que van **antes** que
  él en el orden elegido. El primero espera 0.

- **Suma de esperas (objetivo):** el total a minimizar. Una forma útil de
  verlo: un trabajo que va en la posición `k` hace esperar a todos los que
  quedan después; si quedan `c` trabajos después, su tiempo se "copia" `c`
  veces en la suma total. El primero de todos (con `n - 1` detrás) pesa
  `n - 1` veces; el último no hace esperar a nadie y pesa 0 veces.

- **SJF (Shortest Job First):** disciplina de planificación que atiende
  primero el trabajo más corto. Es el nombre estándar del algoritmo ávido
  de este problema.

- **Permutación:** el resultado debe ser una permutación de `0 .. n-1`:
  todos los índices exactamente una vez. Si todos los tiempos son iguales,
  cualquier orden es óptimo.

## Estrategia greedy

Ordenar los trabajos de **menor a mayor** tiempo de procesamiento y devolver
sus índices en ese orden.

**¿Por qué funciona?** La intuición es que un trabajo corto no "ensucia"
mucho el calendario: dejarlo primero hace esperar poco tiempo a todos los
detrás. Formalmente, con un argumento de intercambio: si en el orden hay un
trabajo largo A (`t_A`) inmediatamente antes de uno corto B (`t_B` con
`t_B < t_A`), invertirlos **mejora** el resultado. Al invertir:

- B pasa de esperar `t_A` más para arrancar a esperar 0 extra (mejora `t_A`);
- A empieza `t_B` más tarde, o sea suma `t_B` a su espera;
- los trabajos anteriores y posteriores al par no cambian.

La suma total baja `t_A - t_B > 0`. Es decir, toda inversión (par fuera de
orden) se puede corregir y mejorar: el único orden sin inversiones es el
ordenado de menor a mayor, y por lo tanto es el óptimo.

Nota de honestidad del greedy: SJF minimiza la **suma** de esperas, pero
puede hacer esperar mucho a los trabajos largos (justo los que van al
final). Es el óptimo para este objetivo, no para todos los objetivos
posibles.

## Ejemplo a mano

`tiempos = [5, 1, 3]` (índices 0, 1, 2).

**Paso 1:** ordenar por tiempo.

| Orden | Índice | Tiempo |
|-------|--------|--------|
| 1     | 1      | 1      |
| 2     | 2      | 3      |
| 3     | 0      | 5      |

Resultado de la función: `[1, 2, 0]`.

**Paso 2:** calcular las esperas para comparar con otros órdenes.

| Orden (índices) | Orden de tiempos | Esperas (índice: espera)          | Suma |
|-----------------|------------------|-----------------------------------|------|
| `[0, 1, 2]` (original) | 5, 1, 3   | 0: 0, 1: 5, 2: 5+1 = 6           | 11   |
| `[2, 1, 0]`     | 3, 1, 5          | 2: 0, 1: 3, 0: 3+1 = 4            | 7    |
| `[1, 2, 0]` (SJF) | 1, 3, 5        | 1: 0, 2: 1, 0: 1+3 = 4            | **5** |

El orden SJF da 5 contra 11 del orden original: casi la mitad.

Casos borde:

- Un solo trabajo `[7]` → `[0]` (espera 0, único orden posible).
- Tiempos iguales `[3, 3, 3]` → cualquier permutación; el orden por
  tiempos da `[0, 1, 2]`.
- Slice vacío → devolver un slice vacío (no `nil`).
