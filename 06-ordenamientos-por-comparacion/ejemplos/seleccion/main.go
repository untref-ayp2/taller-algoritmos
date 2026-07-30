package main

import "fmt"

// OrdenarSeleccion ordena el slice en orden ascendente.
// Complejidad: O(n²) en todos los casos. No es estable. In Place.
func OrdenarSeleccion(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Antes: ", arr)
	OrdenarSeleccion(arr)
	fmt.Println("Después:", arr)
	// Salida:
	// Antes:  [64 25 12 22 11]
	// Después: [11 12 22 25 64]
}
