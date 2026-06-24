package cambio

// CambioTab calcula la cantidad mínima de monedas necesarias para dar un monto
// usando las denominaciones dadas. Usa programación dinámica con tabulación
// (bottom-up).
//
// A diferencia del enfoque ávido, PD no requiere que las denominaciones
// cumplan condiciones especiales (como ser submúltiplos unas de otras).
// Encuentra la solución óptima para cualquier conjunto de denominaciones.
//
// Retorna -1 si no es posible formar el monto con las denominaciones dadas.
//
// Ejemplo: CambioTab(6, []int{1,3,4}) = 2 (usando 3+3)
// Mientras que un algoritmo ávido tomaría 4+1+1 = 3 monedas.
func CambioTab(monto int, denominaciones []int) int {
	// TODO: implementar con tabulación (bottom-up)
	return -1
}

// CambioMemo calcula la cantidad mínima de monedas usando memoización (top-down).
//
// Retorna -1 si no es posible formar el monto con las denominaciones dadas.
func CambioMemo(monto int, denominaciones []int) int {
	// TODO: implementar con memoización (top-down)
	return -1
}
