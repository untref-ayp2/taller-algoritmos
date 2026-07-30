# Taller de Algoritmos

Repositorio complementario de la sección **Diseño de Algoritmos** de los apuntes de Algoritmos y Programación II.

## Estructura

Cada capítulo tiene ejercicios (esqueletos con tests) y los primeros dos también incluyen ejemplos funcionales.

### 01 — Recursividad y División y Conquista (capítulo 4-1)

**Ejemplos:**

| Directorio | Descripción |
|---|---|
| `ejemplos/factorial/` | Cálculo recursivo del factorial de un número |
| `ejemplos/par-impar/` | Recursión indirecta: `esPar` y `esImpar` se llaman mutuamente |
| `ejemplos/busqueda-binaria/` | Búsqueda binaria recursiva (división y conquista) |

**Ejercicios:**

| Directorio | Descripción |
|---|---|
| `ejercicios/01-palindromo/` | Determinar si una cadena es palíndromo (recursivo) |
| `ejercicios/02-contar-caracter/` | Contar apariciones de un caracter en una cadena |
| `ejercicios/03-hanoi/` | Resolver Torres de Hanoi, imprimir pasos |
| `ejercicios/04-busqueda-ternaria/` | Búsqueda ternaria recursiva (divide en 3 partes) |
| `ejercicios/05-pico/` | Encontrar el pico en un arreglo bitónico en O(log n) |
| `ejercicios/06-mcd/` | Máximo común divisor con algoritmo de Euclides |
| `ejercicios/07-potencia/` | Exponenciación binaria recursiva |
| `ejercicios/08-raiz-digital/` | Raíz digital: suma de dígitos hasta un solo dígito |

### 02 — Patrones de Diseño (capítulo 4-2)

**Ejemplos:**

| Directorio | Descripción |
|---|---|
| `ejemplos/adapter/` | Adapter: Robot que mide en cm → cliente espera pulgadas |
| `ejemplos/composite/` | Composite: sistema de archivos con `Archivo` y `Carpeta` |
| `ejemplos/iterator/` | Iterator con API de 2 métodos (`Siguiente` / `Valor`) |

**Ejercicios:**

| Directorio | Descripción |
|---|---|
| `ejercicios/01-sistema-figuras/` | Composite: jerarquía de figuras con área y perímetro |
| `ejercicios/02-notificador/` | Adapter: unificar interfaces de email, SMS y push |
| `ejercicios/03-iterador-lista-doble/` | Iterator para lista doblemente enlazada (avance y retroceso) |

### 03 — Algoritmos Ávidos (capítulo 4-3)

**Ejercicios:**

| Directorio | Descripción |
|---|---|
| `ejercicios/01-sesion/` | Seleccionar máxima cantidad de sesiones sin superponer |
| `ejercicios/02-codigo-huffman/` | Comprimir texto con código Huffman |
| `ejercicios/03-maquina-expendedora/` | Cambio con la menor cantidad de monedas (versión ávida) |
| `ejercicios/04-mochila-fraccionaria/` | Maximizar valor en mochila con ítems fraccionables |
| `ejercicios/05-minimizar-espera/` | Minimizar tiempo total de espera (Shortest Job First) |
| `ejercicios/06-planificar-tareas/` | Maximizar ganancia con tareas con deadline |

### 04 — Backtracking (capítulo 4-4)

**Ejercicios:**

| Directorio | Descripción |
|---|---|
| `ejercicios/01-subset-sum/` | Determinar si un subconjunto suma exactamente k |
| `ejercicios/02-cambio/` | Encontrar todas las combinaciones de monedas para un monto |
| `ejercicios/03-laberinto/` | Encontrar camino desde (0,0) hasta (N-1,N-1) |
| `ejercicios/04-parentesis/` | Generar todas las combinaciones válidas de n pares de paréntesis |
| `ejercicios/05-sudoku/` | Resolver sudoku 9×9 con backtracking y poda |
| `ejercicios/06-particion-igual/` | Dividir arreglo en dos subconjuntos de igual suma |

### 05 — Programación Dinámica (capítulo 4-5)

**Ejercicios:**

| Directorio | Descripción |
|---|---|
| `ejercicios/01-escaleras/` | Contar formas de subir n escalones saltando 1 o 2 |
| `ejercicios/02-cambio/` | Cambio de monedas con PD: cantidad mínima de monedas (tabulación y memoización) |
| `ejercicios/03-subsecuencia/` | Longitud de la subsecuencia común más larga (LCS) |
| `ejercicios/04-coeficiente-binomial/` | Coeficiente binomial C(n,k) con PD (tabulación) |
| `ejercicios/05-varilla/` | Maximizar ganancia al cortar una varilla (Rod Cutting) |
| `ejercicios/06-grilla/` | Camino de costo mínimo en grilla (reconstrucción de ruta) |

### 06 — Ordenamientos por Comparación de Claves (capítulo 4-6)

**Ejercicios:**

| Directorio | Descripción |
|---|---|
| `ejercicios/01-mergesort/` | MergeSort genérico con función de comparación |
| `ejercicios/02-quicksort/` | QuickSort genérico con función de comparación |
| `ejercicios/03-heapsort/` | HeapSort genérico con función de comparación |
| `ejercicios/04-ordenar-estructuras/` | Ordenar struct `Persona` por múltiples criterios |
| `ejercicios/05-quickselect/` | k-ésimo elemento más pequeño (QuickSelect) |
| `ejercicios/06-fusionar-ordenados/` | Fusionar dos slices ordenados |
| `ejercicios/07-contar-inversiones/` | Contar inversiones con MergeSort modificado |
| `ejercicios/08-k-frecuente/` | Top-K elementos más frecuentes |
| `ejercicios/09-ordenar-lista/` | MergeSort in-place sobre lista enlazada simple |

### 07 — Ordenamientos Lineales (capítulo 4-7)

**Ejercicios:**

| Directorio | Descripción |
|---|---|
| `ejercicios/01-counting-sort/` | Counting Sort para enteros no negativos |
| `ejercicios/02-radix-sort/` | Radix Sort usando Counting Sort como subrutina estable |
| `ejercicios/03-bucket-sort/` | Bucket Sort para reales en [0, 1) |
| `ejercicios/04-counting-sort-estable/` | Versión estable de Counting Sort con suma prefijo |
| `ejercicios/05-radix-sort-strings/` | Radix Sort para cadenas alfanuméricas de longitud fija |
| `ejercicios/06-bucket-sort-generalizado/` | Bucket Sort con k buckets configurables |
| `ejercicios/07-aplicacion-estructuras/` | Ordenar personas por edad usando Counting Sort estable |
| `ejercicios/08-bandera-holandesa/` | Dutch National Flag: ordenar arreglo de 0, 1, 2 |
| `ejercicios/09-ordenar-estudiantes/` | Bucket Sort con 101 buckets para promedios [0,10] |
| `ejercicios/10-rango-etario/` | Agrupar personas en 5 rangos etarios con Counting Sort |

## Cómo usar

```bash
# Ejecutar todos los tests
make test

# Con detalle de cada caso
make test-v

# Tests de un ejercicio específico
go test -v ./01-recursividad/ejercicios/01-palindromo/...

# Compilar todo
make build
```

Para ejecutar un ejemplo:

```bash
go run ./01-recursividad/ejemplos/factorial
go run ./02-patrones-de-diseno/ejemplos/adapter
```

Para más información, ver [CONTRIBUTING.md](CONTRIBUTING.md).

## Requisitos

- Go 1.22 o superior

## Licencia

MIT
