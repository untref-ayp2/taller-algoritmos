# Máquina Expendedora (Cambio Ávido)

**Función a implementar:** `CambioAvido(monto int, denominaciones []int) map[int]int`

## El problema

Se debe devolver el cambio de un monto exacto usando las monedas
disponibles, con la **menor cantidad de monedas posible**. La función
devuelve un mapa `denominación → cantidad` (por ejemplo, para 93 con
denominaciones estándar: `map[50:1, 20:2, 2:1, 1:1]`, 5 monedas).

## Conceptos clave

- **Denominación:** el valor de cada moneda (1, 2, 5, 10, …). Las
  denominaciones llegan como un slice, que puede estar en cualquier orden
  e incluir valores mayores al monto (hay que ignorarlos).

- **Cambio:** un conjunto de monedas cuya suma sea **exactamente** el monto.
  No alcanza con acercarse: sobrar o faltar dinero no es una solución
  válida.

- **Menor cantidad de monedas:** el objetivo es minimizar la suma de las
  cantidades, no el valor de las monedas usadas.

- **Sistema canónico:** un conjunto de denominaciones para el cual la
  estrategia de "tomar siempre la moneda más grande que quepa" produce
  siempre la solución óptima. El sistema estándar `1, 2, 5, 10, 20, 50,
  100` (el de los tests) es canónico.

- **Ávido (greedy):** en cada paso se toma la moneda más grande posible, sin
  mirar el resto del camino. Es la decisión local más obvia: conviene
  cubrir el monto con la menor cantidad de monedas, y una moneda grande
  rinde más que varias chicas.

## Estrategia greedy

1. Ordenar las denominaciones de mayor a menor.
2. Recorrerlas; para cada denominación `d`:
   - si `d > monto restante`, la salteo (no entra);
   - si `d <= restante`, tomo `restante / d` monedas de ese valor (el
     máximo posible), resto esa cantidad del monto y sigo.
3. Termina cuando el restante llega a 0. Para `monto == 0`, el mapa
   resultado es vacío.

**¿Por qué funciona?** Si la moneda más grande entra sin pasarse, alguna
solución óptima la usa: reemplazar `k` monedas chicas por una grande que
vale lo mismo (o más, completando el resto con monedas chicas) nunca da
peor resultado. Aplicado en cada paso, queda un problema más chico con la
misma estructura y se repite la misma decisión. Ese es el argumento de
intercambio típico del greedy.

**Cuándo NO funciona:** el greedy no es óptimo para cualquier conjunto de
denominaciones. Con `denominaciones = {1, 3, 4}` y `monto = 6`, el greedy
elige `4 + 1 + 1` (3 monedas) cuando lo óptimo es `3 + 3` (2 monedas). El
ejercicio asume denominaciones canónicas justamente por esto; para conjuntos
arbitrarios haría falta otra técnica (que se verá más adelante en el curso).

## Ejemplo a mano

`monto = 93`, `denominaciones = [1, 2, 5, 10, 20, 50, 100]`.

| Paso | Denominación | Restante antes | Monedas tomadas | Restante después |
|------|--------------|----------------|-----------------|------------------|
| 1    | 100          | 93             | 0 (100 > 93)    | 93               |
| 2    | 50           | 93             | 1               | 43               |
| 3    | 20           | 43             | 2               | 3                |
| 4    | 10           | 3              | 0 (10 > 3)      | 3                |
| 5    | 5            | 3              | 0 (5 > 3)       | 3                |
| 6    | 2            | 3              | 1               | 1                |
| 7    | 1            | 1              | 1               | 0                |

**Resultado:** `{50: 1, 20: 2, 2: 1, 1: 1}` → `50 + 20 + 20 + 2 + 1 = 93`,
con 5 monedas. Que 5 sea el mínimo posible lo garantiza la canonicidad del
sistema: con denominaciones canónicas, tomar siempre la moneda más grande
que entra nunca da peor resultado que ninguna otra combinación.

Casos borde:

- `monto = 0` → mapa vacío (ninguna moneda hace falta).
- Denominaciones mayores al monto se ignoran en el paso 1 (100 en el
  ejemplo).
- El mapa sólo debe contener las denominaciones usadas, con cantidad > 0.
