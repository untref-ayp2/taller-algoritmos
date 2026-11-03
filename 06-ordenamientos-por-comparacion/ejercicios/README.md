# Ejercicios: Ordenamientos por Comparación de Claves

En la carpeta `../ejemplos/` hay implementaciones de referencia de MergeSort,
QuickSort y HeapSort usando una función `cmp` (comparador de 3 vías). Los tres
primeros ejercicios consisten en **adaptar esos algoritmos** para que usen, en
lugar de `cmp`, una función `less` (booleana).

## La función `less` vs `cmp`

Los algoritmos de `../ejemplos/` reciben `cmp func(a, b T) int`, que devuelve
un entero negativo, cero o positivo. En estos ejercicios se pide reemplazarla
por `less func(a, b T) bool`, donde `less(a, b)` indica si `a` va antes que `b`.

**Ventajas de `less`:**
- Es el criterio mínimo que necesita un ordenamiento: responder "¿`a` va antes
  que `b`?" con un booleano.
- Se escribe directo para tipos ordenados: `a.Edad < b.Edad`.
- Menos propenso a errores de signo que un comparador de 3 vías (no hay que
  acordarse de devolver 0).

**Desventajas de `less`:**
- No distingue igualdad en una sola llamada: para saber si `a == b` hay que
  evaluar `less(a, b)` **y** `less(b, a)`, duplicando comparaciones.
- Aporta menos información por comparación que `cmp` (que devuelve `<0`, `0` o
  `>0`).
- La biblioteca estándar idiomática (`slices.SortFunc`, `cmp.Compare`) usa
  comparadores de 3 vías; con `less` hay que adaptar.

---

## 1. Mergesort con `less`

En `../ejemplos/mergesort/` está implementado `MergeSort[T any](s []T, cmp func(a, b T) int)`.
Modificalo para que reciba una función `less` en lugar de `cmp`:

```go
func MergeSort[T any](s []T, less func(a, b T) bool)
```

`less(a, b)` debe indicar si `a` va antes que `b` en el orden final (por
ejemplo, `a < b` para orden ascendente). El algoritmo ordena el slice recibido
(dividiendo recursivamente, ordenando cada mitad y fusionando) sin crear uno
nuevo.

**Requisitos:**
- Usar un slice auxiliar para la fusión (merge)
- Ser estable (elementos iguales mantienen orden relativo)
- Funcionar con cualquier tipo `T`

**Preguntas:**
- ¿Por qué MergeSort necesita O(n) memoria extra?
- ¿Qué complejidad temporal tiene en el peor caso?

→ `01-mergesort/`

---

## 2. Quicksort con `less`

En `../ejemplos/quicksort/` está implementado `QuickSort[T any](s []T, cmp func(a, b T) int)`.
Modificalo para que reciba una función `less` en lugar de `cmp`:

```go
func QuickSort[T any](s []T, less func(a, b T) bool)
```

`less(a, b)` debe indicar si `a` va antes que `b` en el orden final. El
algoritmo ordena el slice recibido seleccionando un pivote (elemento del medio),
particionando y ordenando recursivamente cada subslice, sin crear uno nuevo.

**Requisitos:**
- Elegir el elemento del medio como pivote
- Particionar in-place (partición de 2 vías)
- Funcionar con cualquier tipo `T`

**Preguntas:**
- ¿Qué ocurre si el slice ya está ordenado? ¿Y si está en orden inverso?
- ¿Cómo se podría elegir un mejor pivote?

→ `02-quicksort/`

---

## 3. Heapsort con `less`

En `../ejemplos/heapsort/` está implementado `HeapSort[T any](s []T, cmp func(a, b T) int)`.
Modificalo para que reciba una función `less` en lugar de `cmp`:

```go
func HeapSort[T any](s []T, less func(a, b T) bool)
```

`less(a, b)` debe indicar si `a` va antes que `b` en el orden final. El
algoritmo ordena el slice recibido (construyendo un heap y extrayendo
repetidamente la raíz) sin crear uno nuevo.

**Funciones auxiliares:**
- `downHeap[T any](s []T, n, i int, less func(a, b T) bool)` — hunde el nodo `i` dentro de un heap de tamaño `n`

**Preguntas:**
- ¿HeapSort es estable? Justificar.
- ¿Qué complejidad temporal tiene?

→ `03-heapsort/`

---

## 4. k-ésimo Elemento Más Pequeño (QuickSelect)

Implementar **QuickSelect** para encontrar el k-ésimo elemento más pequeño sin ordenar todo el slice.

La función `QuickSelect[T any](s []T, k int, less func(a, b T) bool) T` devuelve el elemento
en la posición `k` (0-indexed) si el slice estuviera ordenado.

**Nota:** La función lee el slice pero no lo modifica permanentemente (usa una copia o particiona
sobre una copia local).

**Preguntas:**
- ¿Qué complejidad esperada tiene QuickSelect? ¿Y en el peor caso?
- ¿Cómo encontrarías la mediana con QuickSelect?

→ `04-quickselect/`

---

## 5. Fusionar Dos Slices Ordenados

Implementar una función que reciba dos slices ya ordenados y devuelva uno nuevo con todos los elementos ordenados (la fase de fusión de MergeSort).

La función `MergeSorted[T any](a, b []T, less func(a, b T) bool) []T` recorre ambos slices
simultáneamente insertando el menor elemento en cada paso.

**Preguntas:**
- ¿Qué complejidad temporal y espacial tiene?
- ¿Es estable?

→ `05-fusionar-ordenados/`

---

## 6. Contar Inversiones

Implementar una función que cuente la cantidad de pares `(i, j)` con `i < j`
y `arr[i] > arr[j]` en un arreglo de enteros. Debe hacerlo en tiempo
$O(n \log n)$ usando una modificación del algoritmo MergeSort.

**Preguntas:**
- ¿Qué complejidad temporal tiene?
- ¿Cómo se compara con un enfoque ingenuo de $O(n^2)$?

→ `06-contar-inversiones/`

---

## 7. Top-K Elementos Más Frecuentes

Dado un slice de enteros y un entero `k`, devolver los `k` elementos que más
se repiten, ordenados por frecuencia descendente. Usar un mapa de frecuencias
y ordenamiento parcial.

**Preguntas:**
- ¿Qué complejidad temporal tiene?
- ¿Se puede hacer sin ordenar todos los elementos?

→ `07-k-frecuente/`

---

## 8. MergeSort sobre Lista Enlazada

Implementar MergeSort in-place sobre una lista enlazada simple de enteros.
La función recibe la cabeza de la lista y devuelve la nueva cabeza ordenada.

**Preguntas:**
- ¿Por qué MergeSort es preferible a QuickSort en listas enlazadas?
- ¿Qué complejidad espacial tiene esta implementación?

→ `08-ordenar-lista/`

---

**Nota para el alumno**: las respuestas a las preguntas teóricas deben
incluirse como comentarios al final del archivo `.go` de implementación, en un
bloque encabezado con `// === PREGUNTAS TEÓRICAS ===`.
