# Ejercicios: Programación Dinámica

1. **Escaleras.** Una persona puede subir escalones saltando de a 1 o
   2 escalones a la vez. Implementar una función que calcule cuántas
   formas distintas tiene de llegar al escalón `n`. Usar programación
   dinámica (bottom-up o top-down con memoización).
   → `01-escaleras/`

2. **Cambio de monedas.** Dado un monto y un conjunto de denominaciones,
   encontrar la cantidad mínima de monedas necesarias para formar el
   monto. Implementar con **tabulación** (`CambioTab`) y con
   **memoización** (`CambioMemo`).

   A diferencia del algoritmo ávido visto en el capítulo de algoritmos
   ávidos, la programación dinámica **no requiere que las denominaciones
   cumplan condiciones especiales** (como ser submúltiplos). PD siempre
   encuentra la solución óptima para cualquier conjunto de denominaciones.

   Por ejemplo, para monto=6 y denominaciones {1, 3, 4}:
   - Ávido: 4 + 1 + 1 = 3 monedas
   - PD (óptimo): 3 + 3 = 2 monedas
   → `02-cambio/`

3. **Subsecuencia Común Más Larga (LCS).** Dadas dos cadenas, encontrar
   la longitud de la subsecuencia común más larga. Implementar con PD
   usando una tabla `(m+1) × (n+1)`.
   → `03-subsecuencia/`
