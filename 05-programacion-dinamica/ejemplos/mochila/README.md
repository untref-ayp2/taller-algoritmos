# Ejemplo: Mochila 0/1 con Programación Dinámica

El problema de la mochila —enunciado, ecuación de recurrencia, tabla,
applet interactivo y ambas versiones— está desarrollado en el capítulo
**«Programación Dinámica»** del apunte. Este directorio contiene las
implementaciones con sus tests.

## Qué hace el código

- `Item`: struct con `Peso` y `Valor`.

- `MochilaTab(items []Item, capacidad int) int`: **tabulación**
  (bottom-up). Tabla de `(n+1) × (capacidad+1)`; cada celda guarda el
  máximo valor alcanzable con los primeros `i` objetos y esa capacidad.
  O(n × W) en tiempo y espacio.

- `MochilaMemo(items []Item, capacidad int) int`: **memoización**
  (top-down). Misma recurrencia por recursión, con una matriz `memo`
  inicializada en `-1` («no calculado»). O(n × W).

- `MochilaTabConItems(...)`: además del valor máximo, devuelve **qué
  objetos** se eligieron, recorriendo la tabla desde la última celda
  hacia atrás (reconstrucción de la solución).

Recordatorio del enunciado: es la mochila **0/1** — cada objeto se toma
una sola vez o no se toma; no se puede fraccionar (la mochila
fraccionaria es el ejercicio `03-algoritmos-avidos/03-mochila-fraccionaria/`).

## Cómo correrlo

```bash
go test ./05-programacion-dinamica/ejemplos/mochila/
```
