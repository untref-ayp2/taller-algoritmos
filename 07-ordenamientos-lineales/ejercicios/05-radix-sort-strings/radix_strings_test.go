package radix

import (
	"reflect"
	"testing"
)

func TestRadixSortStrings(t *testing.T) {
	tests := []struct {
		input  []string
		length int
	}{
		{[]string{}, 6},
		{[]string{"AAA202"}, 6},
		{[]string{"OPV220", "AAA202", "ZZZ999", "BCD123"}, 6},
		{[]string{"ZZZ000", "AAA999", "BBB111", "AAA000"}, 6},
		{[]string{"000000", "999999", "AAAAAA", "ZZZZZZ", "123ABC"}, 6},
	}
	for _, tt := range tests {
		got := RadixSortStrings(tt.input, tt.length)
		expected := make([]string, len(tt.input))
		copy(expected, tt.input)
		// Orden lexicográfico estándar (el que usa Radix LSD con 0-9 < A-Z)
		for i := 1; i < len(expected); i++ {
			for j := i; j > 0 && expected[j-1] > expected[j]; j-- {
				expected[j-1], expected[j] = expected[j], expected[j-1]
			}
		}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("RadixSortStrings(%v, %d) = %v, esperado %v", tt.input, tt.length, got, expected)
		}
	}
}

func TestRadixSortStringsEstable(t *testing.T) {
	input := []string{"ZAB", "BCA", "ACD", "DBE", "ZEF", "CBA"}
	got := RadixSortStrings(input, 3)
	expected := []string{"ACD", "BCA", "CBA", "DBE", "ZAB", "ZEF"}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("RadixSortStrings inestable o incorrecto.\n  got:      %v\n  esperado: %v", got, expected)
	}
}
