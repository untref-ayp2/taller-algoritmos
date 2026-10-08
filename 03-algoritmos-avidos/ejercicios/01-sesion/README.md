# Selección de Sesiones

**Función a implementar:** `MaxSesiones(sesiones []Sesion) []Sesion`

Dado un conjunto de sesiones con horario de inicio y fin, seleccionar la
máxima cantidad de sesiones que no se superpongan entre sí. Es el problema
clásico de *selección de actividades* (activity selection), el ejemplo
más directo de un algoritmo ávido.

## El problema

Cada sesión ocupa una franja horaria `[Inicio, Fin)`. Hay muchas formas de
elegir un subconjunto sin superposiciones, pero nos interesa el subconjunto
**mayor posible**: la solución óptima. No alcanza con devolver *cualquier*
conjunto válido (por ejemplo, el vacío es válido pero no es óptimo).

La función debe devolver las sesiones elegidas. Si la entrada está vacía,
hay que devolver un slice vacío y **no** `nil` (los tests lo verifican).

## Conceptos clave

- **Sesión:** registro `Sesion{Inicio, Fin}`. La sesión empieza en `Inicio`
  (inclusivo) y termina en `Fin` (exclusivo): el instante `Fin` ya no
  pertenece a la sesión.

- **Superposición:** dos sesiones se superponen si una empieza antes de que
  termine la otra, es decir si `a.Inicio < b.Fin && b.Inicio < a.Fin`. Si una
  termina exactamente cuando empieza la otra (`a.Fin == b.Inicio`) **no** se
  superponen: pueden encadenarse sin problemas.

- **Solución válida:** subconjunto de sesiones sin superposiciones.

- **Solución óptima:** solución válida con la mayor cantidad de sesiones
  posibles.

- **Ávido (greedy):** estrategia que, en cada paso, toma la decisión local
  que parece mejor en ese momento, sin deshacer decisiones anteriores ni
  explorar alternativas (sin backtracking). La pregunta clave de este
  ejercicio es *cuál* decisión local conviene tomar.

## Estrategia greedy

1. Ordenar las sesiones por hora de **fin** ascendente (las que terminan
   más temprano primero).
2. Elegir la primera de esa lista: la que termina antes.
3. Recorrer el resto de la lista en ese orden: si la sesión empieza en o
   después de la última sesión elegida, elegirla; si empieza antes,
   descartarla.
4. Repetir hasta recorrer toda la lista.

La complejidad está dominada por la ordenación: `O(n log n)`.

**¿Por qué elegir la que termina más temprano?** Es la sesión que menos
tiempo ocupa en el calendario: elegirla deja el máximo de tiempo disponible
para las sesiones que faltan. Elegir una que termina más tarde sólo puede
empeorar el resultado: si una solución óptima empezara con una sesión que
termina más tarde que la más temprana, siempre se puede intercambiar por la
más temprana (no se superpone con nada de esa solución porque empieza antes
de la anterior) y el resto de las sesiones sigue cabiendo. Ese es el
*argumento de intercambio* que justifica la corrección de un greedy: la
decisión local nunca empeora la solución global.

Ojo con las trampas habituales: no sirve ordenar por **inicio** (una sesión
que empieza temprano y termina tarde tapa a todas), ni elegir la **más
corta** (no siempre es la que termina antes). La clave es la hora de **fin**.

## Ejemplo a mano

Entrada (en cualquier orden):

| Sesión | Inicio | Fin |
|--------|--------|-----|
| A      | 0      | 6   |
| B      | 1      | 2   |
| C      | 3      | 4   |
| D      | 5      | 7   |

**Paso 1:** ordenar por hora de fin.

| Orden | Sesión | Inicio | Fin |
|-------|--------|--------|-----|
| 1     | B      | 1      | 2   |
| 2     | C      | 3      | 4   |
| 3     | A      | 0      | 6   |
| 4     | D      | 5      | 7   |

**Paso 2:** recorrer y decidir.

| Sesión | ¿Empieza después de la última elegida? | Decisión | Última fin |
|--------|-----------------------------------------|----------|------------|
| B      | es la primera                           | elegir   | 2          |
| C      | `3 >= 2` sí                             | elegir   | 4          |
| A      | `0 >= 4` no                             | descartar| 4          |
| D      | `5 >= 4` sí                             | elegir   | 7          |

**Resultado:** `[B, C, D]` → 3 sesiones. La sesión A, aunque empieza la
primera, se descarta porque tapa el horario que necesitan B, C y D.

Detalle importante: la función debe devolver las sesiones **en orden
cronológico** (de la que termina más temprano a la que termina más tarde);
los tests recorren el resultado suponiendo ese orden para verificar que no
haya superposiciones. El greedy ya las produce en ese orden, porque se
eligieron siguiendo la hora de fin.
