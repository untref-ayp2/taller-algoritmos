package hanoi

import (
	"reflect"
	"testing"
)

func TestHanoi(t *testing.T) {
	tests := []struct {
		n      int
		want   []string
	}{
		{1, []string{"Mover disco 1 de A a C"}},
		{2, []string{
			"Mover disco 1 de A a B",
			"Mover disco 2 de A a C",
			"Mover disco 1 de B a C",
		}},
	}
	for _, tt := range tests {
		var got []string
		Hanoi(tt.n, "A", "C", "B", &got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Hanoi(%d) = %v; esperado %v", tt.n, got, tt.want)
		}
	}
}

func TestHanoiCantidadPasos(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 3},
		{3, 7},
		{4, 15},
	}
	for _, tt := range tests {
		var pasos []string
		Hanoi(tt.n, "A", "C", "B", &pasos)
		if len(pasos) != tt.want {
			t.Errorf("Hanoi(%d) generó %d pasos; esperado %d", tt.n, len(pasos), tt.want)
		}
	}
}
