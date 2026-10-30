# Planificación de Tareas con Plazos

**Función a implementar:** `PlanificarTareas(trabajos []Trabajo) []int`

## El problema

Tenemos un conjunto de trabajos. Cada trabajo ofrece una ganancia si se
completa, pero debe terminarse antes de su *deadline*. El tiempo se divide
en unidades discretas (por ejemplo, horas o días) y en cada unidad sólo
puede ejecutarse un trabajo. El objetivo es **maximizar la ganancia total**
seleccionando un subconjunto de trabajos que puedan agendarse sin
solaparse. La función devuelve los IDs de los trabajos seleccionados, en
orden cronológico (de slot 0 en adelante).

## Conceptos clave

- **Deadline:** es el último momento (índice de tiempo) en el que un trabajo
  puede terminar. Si un trabajo tiene deadline 5, sus unidades de ejecución
  deben ubicarse en slots con índice < 5 (es decir, en los slots 0..4).
  Un trabajo con deadline d ocupa exactamente una unidad de tiempo y debe
  ubicarse en algún slot j tal que 0 <= j < d.

- **Slot de tiempo:** es una unidad indivisible del horario, identificada
  por un índice entero (slot 0, slot 1, slot 2, ...). Cada slot admite a lo
  sumo un trabajo. El arreglo `slots` modela el horario completo: en
  `slots[j]` guardamos el ID del trabajo asignado al slot j, o 0 si el slot
  está libre. Asumimos que los IDs de los trabajos arrancan en 1, así el
  valor cero (el valor por defecto de Go) sirve de marca de slot libre y no
  hay que inicializar el arreglo.

- **Trabajo:** registro `Trabajo{ID, Deadline, Ganancia}`. Un trabajo con
  `Deadline <= 0` no tiene dónde ejecutarse y se descarta.

- **Ganancia:** recompensa por completar el trabajo. Como los trabajos no
  tienen "costo" distinto del tiempo (y todos ocupan 1 slot), maximizar la
  ganancia total se reduce a meter la mayor cantidad de ganancia posible en
  los slots disponibles.

## Estrategia greedy

La clave para maximizar la ganancia es una regla de prioridad simple:

1. Ordenar los trabajos de mayor a menor ganancia.
2. Para cada trabajo (en ese orden), intentar colocarlo en el slot libre
   más cercano a su deadline, es decir, el de mayor índice j < deadline.
   Si ese slot está ocupado, retrocedemos (j--) hasta encontrar uno libre;
   si no queda ninguno, el trabajo no se puede agendar y se descarta.

**¿Por qué elegir el slot más tardío posible?** Porque dejar los slots
tempranos libres maximiza las opciones para los trabajos restantes: un
trabajo con deadline chico sólo puede ir en slots tempranos, mientras que
uno con deadline grande puede ir en cualquiera. Consumir primero los slots
tardíos "reserva" los tempranos para quienes realmente los necesitan.

**¿Por qué ordenar por ganancia?** Es una elección greedy clásica:
preferimos intentar incluir los trabajos más rentables primero. No garantiza
óptimo absoluto en toda variante del problema, pero es la heurística
estándar y funciona muy bien (y es el enfoque que se espera en este
ejercicio).

## Ejemplo a mano

Supongamos un horario de 5 slots (índices 0..4) y los siguientes 10
trabajos, listados por ID tal como llegan por entrada. Ojo: el orden por ID
no tiene nada que ver con el orden por ganancia. El de mayor ganancia es el
ID 1 (ganancia 100) y su deadline es 3, así que sólo puede ejecutarse en
los slots 0, 1 o 2.

**Entrada (ordenada por ID):**

| ID | Deadline | Ganancia |
|----|----------|----------|
| 1  | 3        | 100      | ← mayor ganancia
| 2  | 3        | 30       |
| 3  | 1        | 70       |
| 4  | 5        | 20       |
| 5  | 5        | 90       |
| 6  | 4        | 40       |
| 7  | 5        | 60       |
| 8  | 3        | 10       |
| 9  | 5        | 80       |
| 10 | 2        | 50       |

**Paso 1:** ordenar por ganancia de mayor a menor (si hay empate, da igual
el orden). Fijate cómo queda reordenada la lista, con los IDs todo
revueltos:

| ID          | 1  | 5  | 9  | 3  | 7  | 10 | 6  | 2  | 4  | 8  |
|-------------|----|----|----|----|----|----|----|----|----|----|
| Ganancia    | 100| 90 | 80 | 70 | 60 | 50 | 40 | 30 | 20 | 10 |
| Deadline    | 3  | 5  | 5  | 1  | 5  | 2  | 4  | 3  | 5  | 3  |

**Paso 2:** recorrer esa lista ordenada y colocar cada trabajo en el slot
libre más cercano a su deadline. El arreglo `slots` se muestra con su
contenido real: 0 = libre, cualquier otro valor = ID del trabajo que lo
ocupó.

```
Estado inicial:  slots: [ 0  0  0  0  0 ]

Coloco ID 1 (g=100, d=3): pruebo el slot 2 (el más tardío permitido) y
está libre → lo ocupa.
                          slots: [ 0  0  1  0  0 ]

Coloco ID 5 (g=90, d=5): pruebo el slot 4, libre → lo ocupa.
                          slots: [ 0  0  1  0  5 ]

Coloco ID 9 (g=80, d=5): pruebo el slot 4, pero está ocupado por el 5 →
retrocedo al slot 3, libre → lo ocupa.
                          slots: [ 0  0  1  9  5 ]

Coloco ID 3 (g=70, d=1): con deadline 1 sólo puede ir en el slot 0,
libre → lo ocupa.
                          slots: [ 3  0  1  9  5 ]

Coloco ID 7 (g=60, d=5): pruebo el slot 4 (ocupado por el 5), retrocedo
al 3 (ocupado por el 9), al 2 (ocupado por el 1) → el slot 1 está libre
→ lo ocupa.
                          slots: [ 3  7  1  9  5 ]

A partir de acá el horario quedó completo y todos los trabajos restantes
se descartan:

Coloco ID 10 (g=50, d=2): slots 1 y 0 ocupados → descartado.
Coloco ID 6 (g=40, d=4): slots 3, 2, 1 y 0 ocupados → descartado.
Coloco ID 2 (g=30, d=3): slots 2, 1 y 0 ocupados → descartado.
Coloco ID 4 (g=20, d=5): slots 4, 3, 2, 1 y 0 ocupados → descartado.
Coloco ID 8 (g=10, d=3): slots 2, 1 y 0 ocupados → descartado.
```

**Resultado:** recorriendo el horario de slot 0 en adelante queda
`[3 7 1 9 5]`, con ganancia total 70 + 60 + 100 + 80 + 90 = 400.

Casos borde:

- Trabajos con `Deadline <= 0` (por ejemplo, los del test
  `TestDeadlineCero`) no caben en ningún slot: se descartan y el resultado
  es un slice vacío.
- Un solo trabajo con deadline 1 → `[su ID]`, porque sólo ocupa el slot 0.
- Si la entrada está vacía, devolver un slice vacío (no `nil`).
