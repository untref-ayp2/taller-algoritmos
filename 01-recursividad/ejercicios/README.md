# Ejercicios: Recursividad y División y Conquista

1. **Palíndromo.** Escribir una función recursiva que tome una cadena
   y devuelva `true` si la cadena es un palíndromo o `false` en caso contrario.
   → `01-palindromo/`

2. **Contar caracter.** Escribir una función recursiva que reciba una cadena
   y un caracter y devuelva la cantidad de veces que aparece el caracter
   en la cadena.
   → `02-contar-caracter/`

3. **Torres de Hanoi.** Escribir una función recursiva que resuelva el
   juego de las [Torres de Hanoi](https://es.wikipedia.org/wiki/Torres_de_Han%C3%B3i).
   Debe imprimir los pasos necesarios para mover `n` discos del poste origen
   al poste destino usando un poste auxiliar.
   → `03-hanoi/`

4. **Búsqueda ternaria.** Programar la búsqueda ternaria recursiva, donde
   en lugar de dividir el arreglo en dos partes iguales, se divide en tres
   partes iguales. Calcular el orden aplicando el Teorema maestro y
   comparar el orden con la búsqueda binaria.
   → `04-busqueda-ternaria/`

5. **Pico.** Se tiene un arreglo de `len(arr) >= 3` elementos en forma de
   pico, esto es: estrictamente creciente hasta una determinada posición `p`,
   y estrictamente decreciente a partir de ella (con `0 < p < n-1`).
   Por ejemplo, en el arreglo `[1, 2, 3, 1, 0, -2]` la posición del pico
   es `p = 2`. Implementar un algoritmo de división y conquista de orden
   `O(log n)` que encuentre la posición `p` del pico.
   → `05-pico/`
