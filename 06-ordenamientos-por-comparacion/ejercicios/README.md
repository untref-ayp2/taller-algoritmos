# Ejercicios: Ordenamientos por Comparación de Claves

---

## 1. Mergesort Genérico

Implementar el algoritmo **MergeSort** sobre un slice `[]T` usando una función de comparación `less`.

La función `MergeSort[T any](s []T, less func(a, b T) bool)` ordena el slice recibido
(dividiendo recursivamente, ordenando cada mitad y fusionando) sin crear uno nuevo.

**Requisitos:**
- Usar un slice auxiliar para la fusión (merge)
- Ser estable (elementos iguales mantienen orden relativo)
- Funcionar con cualquier tipo `T`

**Preguntas:**
- ¿Por qué MergeSort necesita en O(n) memoria extra?
- ¿Qué complejidad temporal tiene en el peor caso?

→ `01-mergesort/`

---

## 2. Quicksort Genérico

Implementar el algoritmo **QuickSort** sobre un slice `[]T` usando una función de comparación `less`.

La función `QuickSort[T any](s []T, less func(a, b T) bool)` ordena el slice recibido
seleccionando un pivote (último elemento), particionando y ordenando recursivamente cada subslice, sin crear uno nuevo.

**Requisitos:**
- Usar el último elemento como pivote
- Particionar in-place
- Funcionar con cualquier tipo `T`

**Preguntas:**
- ¿Qué ocurre si el slice ya está ordenado? ¿Y si está en orden inverso?
- ¿Cómo se podría elegir un mejor pivote?

→ `02-quicksort/`

---

## 3. Heapsort Genérico

Implementar el algoritmo **HeapSort** sobre un slice `[]T` usando una función de comparación `less`.

La función `HeapSort[T any](s []T, less func(a, b T) bool)` ordena el slice recibido
(construyendo un heap y extrayendo repetidamente la raíz) sin crear uno nuevo.

**Funciones auxiliares:**
- `downHeap[T any](s []T, n, i int, less func(a, b T) bool)` — hunde el nodo `i` dentro de un heap de tamaño `n`

**Preguntas:**
- ¿HeapSort es estable? Justificar.
- ¿Qué complejidad temporal tiene?

→ `03-heapsort/`

---

## 4. Ordenar Estructuras

Ordenar un slice de `Persona{nombre string, edad int}` usando diferentes criterios.

Elegir **uno** de los algoritmos implementados en los ejercicios 1-3 (MergeSort, QuickSort o
HeapSort) y usarlo para implementar tres funciones:
- `OrdenarPorEdadAsc` — de menor a mayor edad
- `OrdenarPorNombreDesc` — alfabéticamente en orden descendente
- `OrdenarPorEdadAscNombreDesc` — por edad ascendente y, ante empate, por nombre descendente

Cada función debe pasar la función `less` adecuada al algoritmo elegido.

**Preguntas:**
- ¿Cómo harías para componer comparadores genéricamente?
- ¿Qué ventaja tiene separar el criterio de orden del algoritmo?
- ¿Cambia la elección del algoritmo el resultado final?

→ `04-ordenar-estructuras/`

---

## 5. k-ésimo Elemento Más Pequeño (QuickSelect)

Implementar **QuickSelect** para encontrar el k-ésimo elemento más pequeño sin ordenar todo el slice.

La función `QuickSelect[T any](s []T, k int, less func(a, b T) bool) T` devuelve el elemento
en la posición `k` (0-indexed) si el slice estuviera ordenado.

**Nota:** La función lee el slice pero no lo modifica permanentemente (usa una copia o particiona
sobre una copia local).

**Preguntas:**
- ¿Qué complejidad esperada tiene QuickSelect? ¿Y en el peor caso?
- ¿Cómo encontrarías la mediana con QuickSelect?

→ `05-quickselect/`

---

## 6. Fusionar Dos Slices Ordenados

Implementar una función que reciba dos slices ya ordenados y devuelva uno nuevo con todos los elementos ordenados (la fase de fusión de MergeSort).

La función `MergeSorted[T any](a, b []T, less func(a, b T) bool) []T` recorre ambos slices
simultáneamente insertando el menor elemento en cada paso.

**Preguntas:**
- ¿Qué complejidad temporal y espacial tiene?
- ¿Es estable?

→ `06-fusionar-ordenados/`

---

## 7. Contar Inversiones

Implementar una función que cuente la cantidad de pares `(i, j)` con `i < j`
y `arr[i] > arr[j]` en un arreglo de enteros. Debe hacerlo en tiempo
$O(n \log n)$ usando una modificación del algoritmo MergeSort.

**Preguntas:**
- ¿Qué complejidad temporal tiene?
- ¿Cómo se compara con un enfoque ingenuo de $O(n^2)$?

→ `07-contar-inversiones/`

---

## 8. Top-K Elementos Más Frecuentes

Dado un slice de enteros y un entero `k`, devolver los `k` elementos que más
se repiten, ordenados por frecuencia descendente. Usar un mapa de frecuencias
y ordenamiento parcial.

**Preguntas:**
- ¿Qué complejidad temporal tiene?
- ¿Se puede hacer sin ordenar todos los elementos?

→ `08-k-frecuente/`

---

## 9. MergeSort sobre Lista Enlazada

Implementar MergeSort in-place sobre una lista enlazada simple de enteros.
La función recibe la cabeza de la lista y devuelve la nueva cabeza ordenada.

**Preguntas:**
- ¿Por qué MergeSort es preferible a QuickSort en listas enlazadas?
- ¿Qué complejidad espacial tiene esta implementación?

→ `09-ordenar-lista/`

---

**Nota para el alumno**: las respuestas a las preguntas teóricas deben
incluirse como comentarios al final del archivo `.go` de implementación, en un
bloque encabezado con `// === PREGUNTAS TEÓRICAS ===`.
