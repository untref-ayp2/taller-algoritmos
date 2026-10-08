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

4. **Coeficiente Binomial.** Calcular C(n, k) con programación
   dinámica (tabulación), construyendo el triángulo de Pascal fila
   por fila.
   → `04-coeficiente-binomial/`

5. **Corte de Varilla.** Dada una varilla de longitud `n` y una lista
   de precios por longitud, maximizar la ganancia al cortarla
   (Rod Cutting).
   → `05-varilla/`

6. **Camino de Costo Mínimo.** Dada una grilla con costos por celda,
   encontrar el camino de costo mínimo de `(0,0)` a la esquina
   inferior derecha y reconstruir la ruta.
   → `06-grilla/`
