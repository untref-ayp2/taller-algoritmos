package contar

import "testing"

func TestContarCaracter(t *testing.T) {
	tests := []struct {
		s     string
		c     rune
		want  int
	}{
		{"", 'a', 0},
		{"hola", 'a', 1},
		{"banana", 'a', 3},
		{"banana", 'n', 2},
		{"abcdef", 'z', 0},
		{"aaaaaa", 'a', 6},
	}
	for _, tt := range tests {
		got := ContarCaracter(tt.s, tt.c)
		if got != tt.want {
			t.Errorf("ContarCaracter(%q, %q) = %d; esperado %d", tt.s, tt.c, got, tt.want)
		}
	}
}
