# Ejemplo: Problema de las N Reinas

Este ejemplo está desarrollado en detalle en el capítulo
**«Backtracking (Vuelta Atrás)»** del apunte, que incluye el enunciado,
el esquema general del algoritmo, el análisis de complejidad y la
implementación. Acá sólo se repasa qué hace el código.

## El problema

Colocar N reinas en un tablero N × N de manera que ninguna ataque a
otra: ni en la misma fila, ni en la misma columna, ni en ninguna
diagonal.

## Qué hace el código

- `NReinas(n int) []int`: devuelve un slice donde el **índice** es la
  fila y el **valor** es la columna de la reina de esa fila. Por
  ejemplo, `[2, 0, 3, 1]` para N = 4 (la misma representación que en
  el apunte).

- `esSolucion(n, solucionParcial) bool`: verifica si la solución parcial ya
  tiene una reina en cada fila (su longitud es `n`).

- `esFactible(fila, columna, solucionParcial) bool`: es la
  `esFactible` del esquema. Chequea que no haya otra reina en la misma
  columna (valores repetidos en el slice) ni en las diagonales, usando
  que `fila - columna` es constante en una diagonal y `fila + columna`
  en la otra.

- `registrar(solucionParcial, columna)` agrega la reina elegida, y
  `borrar(solucionParcial)` la quita al volver atrás.

- `NReinasConPaso(n, obs)`: igual que `NReinas`, pero además notifica
  cada evento a un `ObservadorReinas` con dos callbacks: `Intento(fila,
  columna, factible, parcial)` en **cada intento** de colocación, y
  `Solucion(solucion) bool` al completar una solución. Devolver `true`
  desde `Solucion` obliga a seguir el backtracking; `false` lo detiene.

- `NReinasTodas(n)`: devuelve todas las soluciones. Útil para saber
  cuántas hay (por ejemplo, 2 para N = 4, 4 para N = 6).

- `backtracking(n, fila, solucionParcial, obs)`: la recursión general del
  esquema. Si `esSolucion` ya hay una reina por fila, avisa al observador.
  Si no, prueba todas las columnas de la fila actual (`extender`); para
  cada una avisa del intento, y si `esFactible` registra la reina
  (`registrar`) y continúa en la fila siguiente. Al volver, la quita
  (`borrar`). Si ninguna columna funciona, la rama se corta (vuelta atrás).

## Visualización

`nreinas_visual.go` es la capa de dibujo, separada del backtracking:

- `TableroNReinas(n, reinas, actualFila, actualColumna, factible)` devuelve
  el tablero con las reinas ya ubicadas (♛), resaltando el intento actual
  en verde si es factible o en rojo (`×`) si está atacado.
- `LeerCantidadReinas(entrada, salida)` pregunta cuántas reinas se quieren
  (entre 4 y 10; con Enter usa 6).
- `AnimarReinas(n, entrada, salida, demora)` recorre los intentos, dibuja
  cada paso con una pausa y, al encontrar una solución, la deja apilada
  arriba mientras anima debajo la búsqueda de la próxima; tras cada
  solución pregunta si seguir.

## Cómo correrlo

```bash
go run ./04-backtracking/ejemplos/nreinas/
```

Primero pregunta cuántas reinas querés (entre 4 y 10; por defecto 6) y
después anima la búsqueda: se ve cómo cada reina prueba cada columna hasta
encontrar su lugar. Al hallar una solución el programa pregunta `¿Buscar
otra solución? [s/n]` y, si se responde que sí, fuerza al backtracking a
continuar hasta la próxima. Cuando se agotan, avisa cuántas encontró.

## Para experimentar

El programa ya te deja elegir al arrancar: pregunta cuántas reinas querés
(entre 4 y 10; con Enter usa 6). Además, desde el código:

- `NReinasTodas(n)` te dice **cuántas** soluciones hay: 2 para 4 reinas, 10
  para 5, 4 para 6 y 92 para 8.
- `NReinasConPaso(n, obs)` te deja engancharte en cada intento y en cada
  solución; el observador decide si seguir buscando o parar.
- La pausa entre pasos es el último argumento de `AnimarReinas` (en
  `main.go`): bajala para 8 o 10 reinas, que tienen muchísimas más ramas.

## Tests

```bash
go test ./04-backtracking/ejemplos/nreinas/
```

Verifican que las soluciones sean válidas, cuántas hay para cada N y que la
animación se detenga o continúe según la respuesta del usuario. Agregá tus
propios valores de N y probá.
