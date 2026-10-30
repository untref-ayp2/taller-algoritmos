# Partición en Dos Subconjuntos de Igual Suma

**Función a implementar:**
`ParticionIgual(nums []int) ([]int, []int, bool)`

## El problema

Dividir los números de `nums` en **dos subconjuntos disjuntos** (cada
número aparece exactamente una vez, en uno de los dos) cuya suma sea
la misma. La función devuelve `(subA, subB, true)` si es posible, o
`(nil, nil, false)` si no lo es.

## Conceptos clave

- **Suma total:** si es **impar** no se puede partir en dos mitades
  iguales: es una condición necesaria que permite cortar antes de
  explorar. No es suficiente: `[1, 2, 7]` suma 10 pero tampoco se
  puede (el 7 solo ya pasa la mitad).

- **Reducción a subset sum:** si un subconjunto suma `total / 2`, el
  resto de los números suma lo mismo automáticamente. Así que alcanza
  con buscar **un** subconjunto cuya suma sea `objetivo = total / 2`;
  el otro subconjunto es el complemento (los que no quedaron adentro).

- **solucionParcial:** un `[]int` con los números elegidos para el
  primer subconjunto.

- **esSolucion:** la suma parcial es exactamente `objetivo`.

- **esFactible:** `sumaParcial + nums[i] <= objetivo`.

- **Poda:** si la suma parcial superó el objetivo, la rama está muerta
  (con números positivos ya no se vuelve atrás).

- También se puede resolver con programación dinámica (capítulo 4-5:
  subestructura óptima y subproblemas superpuestos); acá se practica
  backtracking.

## Estrategia de backtracking

1. Sumar todo. Si la suma es impar → `(nil, nil, false)` sin explorar.
2. Recorrer los números desde un índice `start` en adelante (igual que
   en subset sum: cada número se usa a lo sumo una vez y no se permuta).
3. Para cada número: si la suma parcial no se pasa del objetivo,
   agregarlo (`registrar`) y recursar; si la rama falla, quitarlo
   (`borrar`) y seguir con el siguiente.
4. Si la suma parcial llega a `objetivo`, el primer subconjunto está:
   el segundo es el complemento.

Ordenar los números de **mayor a menor** ayuda a podar antes: los
números grandes acercan rápido al objetivo y hacen fallar las ramas
inviables cuanto antes. (Los tests sólo comparan las sumas de los dos
subconjuntos, no el orden de los elementos.)

## Ejemplo a mano

Caso del test: `nums = [1, 5, 11, 5]`.

- Suma total: 22 → par, sigue en pie. `objetivo = 11`.
- Ordenados de mayor a menor: `[11, 5, 5, 1]`.

```
[] (suma 0, objetivo 11)
└── +11 → 11 ✓ solución
```

Primer subconjunto: `[11]`. Segundo (complemento): `[1, 5, 5]`.
Sumas: 11 y 11 ✓.

Para ver una poda, tomemos `[4, 5, 2, 1]`: suma 10, `objetivo = 6`,
ordenados `[5, 4, 2, 1]`:

```
[] (suma 0, objetivo 6)
├── +5 → 5
│   ├── +4 → 9 ✗ poda (9 > 6)
│   ├── +2 → 7 ✗ poda (7 > 6)
│   └── +1 → 6 ✓ solución → [5, 1]
```

Primer subconjunto `[5, 1]`, complemento `[4, 2]`; sumas 6 y 6 ✓.

Casos borde:

- Suma impar: `[1, 2, 3, 5]` suma 11 → `(nil, nil, false)` sin
  explorar.
- Todos iguales: `[2, 2, 2, 2]` → objetivo 4 → `[2, 2]` y `[2, 2]`.
- Si ningún subconjunto llega a la mitad → `(nil, nil, false)`.
