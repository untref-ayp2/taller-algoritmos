package main

import "fmt"

func busquedaBinaria(array []int, inicio, fin, x int) int {
	if inicio > fin {
		return -1
	}

	medio := (inicio + fin) / 2

	if array[medio] > x {
		return busquedaBinaria(array, inicio, medio-1, x)
	}

	if array[medio] < x {
		return busquedaBinaria(array, medio+1, fin, x)
	}

	return medio
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	pos := busquedaBinaria(arr, 0, len(arr)-1, 7)
	fmt.Printf("Posición de 7: %d\n", pos)

	pos = busquedaBinaria(arr, 0, len(arr)-1, 10)
	fmt.Printf("Posición de 10: %d\n", pos)
}
