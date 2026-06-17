package hanoi

import "fmt"

// Hanoi imprime los pasos para mover n discos del poste origen al poste
// destino usando el poste auxiliar.
func Hanoi(n int, origen, destino, auxiliar string, pasos *[]string) {
	// TODO: implementar
}

func main() {
	var pasos []string
	Hanoi(3, "A", "C", "B", &pasos)
	for _, p := range pasos {
		fmt.Println(p)
	}
}
