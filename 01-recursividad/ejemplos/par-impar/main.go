package main

import "fmt"

func esPar(n int) bool {
	if n == 0 {
		return true
	}
	return esImpar(n - 1)
}

func esImpar(n int) bool {
	if n == 0 {
		return false
	}
	return esPar(n - 1)
}

func main() {
	fmt.Println("esPar(4):", esPar(4))
	fmt.Println("esPar(7):", esPar(7))
	fmt.Println("esImpar(3):", esImpar(3))
	fmt.Println("esImpar(6):", esImpar(6))
}
