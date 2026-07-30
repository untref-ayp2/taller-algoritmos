package main

import "fmt"

// OrdenarInsercion ordena el slice en orden ascendente.
// Complejidad: O(n²) peor caso, O(n) mejor caso. Estable, In Place, Online.
func OrdenarInsercion(arr []int) {
	for i := 1; i < len(arr); i++ {
		clave := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > clave {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = clave
	}
}

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Antes: ", arr)
	OrdenarInsercion(arr)
	fmt.Println("Después:", arr)
	// Salida:
	// Antes:  [5 2 4 6 1 3]
	// Después: [1 2 3 4 5 6]
}
