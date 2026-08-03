# Ejercicios: Ordenamientos Lineales

---

## 1. Counting Sort

Implementar **Counting Sort** para ordenar un slice de enteros no negativos recibiendo el rango máximo como parámetro.

La función `CountingSort(arr []int, maxVal int) []int` recibe un slice de enteros en el rango `[0, maxVal]` y devuelve un nuevo slice ordenado.

**Requisitos:**
- Versión no estable (reconstrucción secuencial del arreglo de conteo)
- Devolver un nuevo slice, no modificar el original

**Preguntas:**
- ¿Qué complejidad temporal tiene? ¿De qué depende el espacio extra?
- ¿Qué ocurre si maxVal es mucho mayor que n?

→ `01-counting-sort/`

---

## 2. Radix Sort

Implementar **Radix Sort** (LSD) para ordenar un slice de enteros no negativos usando Counting Sort como subrutina estable para cada dígito.

La función `RadixSort(arr []int) []int` procesa dígito por dígito desde el menos significativo.

**Requisitos:**
- Usar Counting Sort como subrutina estable
- Devolver un nuevo slice ordenado

**Preguntas:**
- ¿Por qué la subrutina debe ser estable?
- ¿Qué complejidad temporal tiene? ¿Y si los números tienen distinta cantidad de dígitos?

→ `02-radix-sort/`

---

## 3. Bucket Sort (básico)

Implementar **Bucket Sort** para ordenar un slice de números reales uniformemente distribuidos en el rango `[0, 1)`.

La función `BucketSort(arr []float64) []float64` distribuye los elementos en buckets,
ordena cada bucket individualmente (por ejemplo con Insertion Sort) y concatena.

**Preguntas:**
- ¿Qué complejidad tiene en promedio? ¿Y en el peor caso?
- ¿El algoritmo es estable? ¿De qué depende?

→ `03-bucket-sort/`

---

## 4. Counting Sort Estable

Implementar la versión **estable** de Counting Sort, que preserva el orden relativo de elementos iguales.

La función `CountingSortStable(arr []int, maxVal int) []int` debe:
1. Contar frecuencias en un arreglo de conteo
2. Convertir las frecuencias en sumas prefijo (posiciones finales)
3. Recorrer el arreglo original **de derecha a izquierda** colocando cada elemento en su posición según la suma prefijo

**Preguntas:**
- ¿Por qué es necesario recorrer de derecha a izquierda para mantener la estabilidad?
- ¿Qué ventaja tiene esta versión sobre la no estable?
- ¿Se puede usar Counting Sort estable como subrutina de Radix Sort?

→ `04-counting-sort-estable/`

---

## 5. Radix Sort sobre Cadenas

Implementar **Radix Sort** para ordenar un slice de cadenas alfanuméricas mayúsculas (0-9, A-Z) de longitud fija.

La función `RadixSortStrings(arr []string, length int) []string` procesa desde el último carácter
(derecha) hasta el primero (izquierda) usando Counting Sort como subrutina estable.
El charset tiene 36 caracteres posibles: 10 dígitos (0-9) seguidos de 26 letras (A-Z).

**Requisitos:**
- Las cadenas contienen solo caracteres 0-9 y A-Z
- Todas tienen exactamente `length` caracteres
- El orden debe ser lexicográfico estándar (0-9 < A-Z)

**Preguntas:**
- ¿Cómo se modifica el algoritmo si las cadenas tuvieran longitud variable?
- ¿Cómo adaptarías el algoritmo para ordenar cadenas que incluyan minúsculas?

→ `05-radix-sort-strings/`

---

## 6. Bucket Sort Generalizado

Implementar una versión generalizada de **Bucket Sort** que funcione para cualquier rango de valores.

La función `BucketSort(arr []float64, k int) []float64` debe:
1. Calcular el mínimo y máximo del slice
2. Dividir el rango `[min, max]` en `k` intervalos del mismo tamaño
3. Distribuir cada elemento en el bucket correspondiente
4. Ordenar cada bucket individualmente
5. Concatenar los buckets en orden

Si `k <= 0`, usar `k = 10`.

**Preguntas:**
- ¿Cómo cambia la complejidad si los datos no están uniformemente distribuidos?
- ¿Qué valor de k conviene elegir? ¿Cómo se relaciona con n?
- ¿La versión básica (ejercicio 3) es un caso particular de esta?

→ `06-bucket-sort-generalizado/`

---

## 7. Aplicación: Ordenar Personas por Edad

Usar **Counting Sort estable** para ordenar un slice de estructuras `Persona{nombre string, edad int}`.

Funciones a implementar:
- `OrdenarPorEdadAsc(personas []Persona) []Persona` — orden ascendente
- `OrdenarPorEdadDesc(personas []Persona) []Persona` — orden descendente

Ambas deben ser estables (preservar orden relativo ante edades iguales).
La edad es un entero en rango acotado, lo que hace a Counting Sort ideal para esta tarea.

**Preguntas:**
- ¿Por qué Counting Sort es adecuado para ordenar por edad?
- ¿Qué complejidad temporal tiene ordenar n personas? ¿Y si usáramos MergeSort en su lugar?
- ¿Cómo harías para ordenar primero por edad y luego por nombre dentro de cada edad?

→ `07-aplicacion-estructuras/`

---

**Nota para el alumno**: las respuestas a las preguntas teóricas deben
incluirse como comentarios al final del archivo `.go` de implementación, en un
bloque encabezado con `// === PREGUNTAS TEÓRICAS ===`.
