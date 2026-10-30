# Mochila Fraccionaria

**Función a implementar:** `MochilaFraccionaria(items []Item, capacidad float64) float64`

## El problema

Hay una mochila con una `capacidad` dada (un peso máximo soportable) y un
conjunto de ítems, cada uno con su peso y su valor. A diferencia de la
mochila clásica, acá se puede llevar un ítem **entero o una fracción** de
él: si algo no entra completo, se corta y se lleva la parte que quepa. El
objetivo es **maximizar el valor total** llevado. La función devuelve sólo
el valor total (un `float64`), no el detalle de qué se llevó.

## Conceptos clave

- **Ítem:** registro `Item{Peso, Valor}` con ambos campos en `float64`,
  porque los pesos y valores pueden no ser enteros y porque se permite
  fraccionar.

- **Densidad (valor por peso):** `ValorPorPeso() = Valor / Peso`. Es el
  valor que aporta cada unidad de peso del ítem. Un ítem `{Peso: 10,
  Valor: 60}` aporta 6 de valor por unidad de peso; `{Peso: 30, Valor: 120}`
  aporta 4. Es la medida de "cuánto rinde" cada ítem. (Si `Peso == 0`, el
  método ya está implementado y devuelve 0.)

- **Fraccionamiento:** tomar una fracción `f` (entre 0 y 1) de un ítem
  aporta `f × Valor` de valor y ocupa `f × Peso` de capacidad, en proporción
  exacta. Este es el detalle que hace que el greedy funcione: nunca queda un
  hueco que no se puede aprovechar, porque el último ítem se ajusta al
  espacio restante.

- **Solución óptima:** combinación (con fracciones permitidas) de mayor
  valor total que respeta la capacidad.

A diferencia de otros problemas de esta serie, la complejidad no importa
mucho: con ordenar y recorrer una vez alcanza (`O(n log n)`).

## Estrategia greedy

1. Ordenar los ítems por densidad (valor/peso) de **mayor a menor**.
2. Recorrer ese orden mientras quede capacidad:
   - si el ítem entra **entero**, llevarlo entero y restarlo de la
     capacidad;
   - si **no entra entero**, llevar sólo la fracción que llene el resto de
     la capacidad (`valorPorPeso × capacidadRestante`) y terminar.
3. Si la capacidad llega a 0 o se acaban los ítems, parar. Si
   `capacidad <= 0` o no hay ítems, el resultado es 0.

**¿Por qué funciona?** Siempre conviene llenar la mochila con lo que más
valor por unidad de peso da, porque cada kilo disponible es un recurso
escaso y el mismo kilo rinde más en el ítem de mayor densidad. Se puede
probar por intercambio: si la mochila contiene un ítem de densidad menor
que algún ítem que quedó afuera, reemplazar una parte del primero por el
segundo aumenta el valor sin pasar de la capacidad; es decir, cualquier
solución que no respete el orden por densidad se puede mejorar. La solución
óptima, entonces, es llenar respetando ese orden.

**Ojo:** este argumento **depende** de poder fraccionar. En la variante en
la que sólo se pueden llevar ítems enteros (mochila 0/1), el greedy por
densidad no da el óptimo; esa variante se ve más adelante y requiere otra
técnica.

## Ejemplo a mano

Ítems de los tests (y sus densidades):

| Ítem | Peso | Valor | Densidad (valor/peso) |
|------|------|-------|-----------------------|
| A    | 10   | 60    | 60/10 = **6**         |
| B    | 20   | 100   | 100/20 = **5**        |
| C    | 30   | 120   | 120/30 = **4**        |

`capacidad = 50`. Orden por densidad: A, B, C (ya viene ordenado).

| Paso | Ítem | ¿Entra entero? | Acción | Valor acumulado | Capacidad restante |
|------|------|----------------|--------|-----------------|--------------------|
| 1    | A    | sí (10 ≤ 50)   | lo llevo entero | 60        | 40                 |
| 2    | B    | sí (20 ≤ 40)   | lo llevo entero | 160       | 20                 |
| 3    | C    | no (30 > 20)   | llevo 20/30 = 2/3 de C | 160 + 80 = **240** | 0 |

El ítem C se fracciona: aporta `120 × (20/30) = 80` de valor (equivalente a
`4 × 20 = 80` por su densidad). **Resultado: 240**, que es el óptimo.

Otro caso del test: un solo ítem `{Peso: 10, Valor: 60}` con `capacidad =
5`. No entra entero, así que llevo la mitad: `60 × (5/10) = 30`.

Casos borde: `capacidad = 0` (o negativa) → 0; sin ítems → 0. Los tests
comparan con tolerancia `0.001`, así que basta con no acumular errores
grandes de redondeo.
