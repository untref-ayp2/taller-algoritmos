package main

import "fmt"

// OrdenarBurbujeo ordena el slice en orden ascendente.
// Complejidad: O(n²) peor caso, O(n) mejor caso. Estable, In Place.
func OrdenarBurbujeo(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		intercambiado := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				intercambiado = true
			}
		}
		if !intercambiado {
			break
		}
	}
}

func main() {
	arr := []int{5, 1, 4, 2, 8}
	fmt.Println("Antes: ", arr)
	OrdenarBurbujeo(arr)
	fmt.Println("Después:", arr)
	// Salida:
	// Antes:  [5 1 4 2 8]
	// Después: [1 2 4 5 8]
}
