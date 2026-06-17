package palindromo

import "testing"

func TestEsPalindromo(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"radar", true},
		{"reconocer", true},
		{"golang", false},
		{"anita lava la tina", true},
	}
	for _, tt := range tests {
		got := EsPalindromo(tt.input)
		if got != tt.want {
			t.Errorf("EsPalindromo(%q) = %v; esperado %v", tt.input, got, tt.want)
		}
	}
}
