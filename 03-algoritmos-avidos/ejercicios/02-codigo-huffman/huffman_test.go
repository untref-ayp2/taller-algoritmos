package huffman

import "testing"

func TestHuffmanCodigos(t *testing.T) {
	freqs := map[rune]int{'a': 5, 'b': 9, 'c': 12, 'd': 13, 'e': 16, 'f': 45}
	arbol := Construir(freqs)
	if arbol == nil || arbol.Raiz == nil {
		t.Fatal("Construir devolvió árbol nil")
	}
	codigos := arbol.Codigos()
	if len(codigos) != len(freqs) {
		t.Errorf("se esperaban %d códigos, se obtuvieron %d", len(freqs), len(codigos))
	}
	for r, c := range codigos {
		if c == "" {
			t.Errorf("código vacío para %q", r)
		}
	}
}

func TestHuffmanCompresion(t *testing.T) {
	freqs := map[rune]int{'a': 5, 'b': 9, 'c': 12, 'd': 13, 'e': 16, 'f': 45}
	arbol := Construir(freqs)
	original := "abcdef"
	comprimido := arbol.Comprimir(original)
	if comprimido == "" {
		t.Error("Comprimir devolvió cadena vacía")
	}
	descomprimido := arbol.Descomprimir(comprimido)
	if descomprimido != original {
		t.Errorf("ciclo compresión/descompresión: %q → %q", original, descomprimido)
	}
}
