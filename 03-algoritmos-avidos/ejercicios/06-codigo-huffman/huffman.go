package huffman

type Nodo struct {
	Car   rune
	Freq  int
	Izq,
	Der *Nodo
}

type ArbolHuffman struct {
	Raiz *Nodo
}

func Construir(frecuencias map[rune]int) *ArbolHuffman {
	// TODO: implementar usando cola de prioridad
	return nil
}

func (a *ArbolHuffman) Codigos() map[rune]string {
	// TODO: recorrer el árbol y generar códigos binarios
	return nil
}

func (a *ArbolHuffman) Comprimir(texto string) string {
	// TODO: reemplazar cada caracter por su código Huffman
	return ""
}

func (a *ArbolHuffman) Descomprimir(codigo string) string {
	// TODO: recorrer el árbol según bits para reconstruir texto
	return ""
}
