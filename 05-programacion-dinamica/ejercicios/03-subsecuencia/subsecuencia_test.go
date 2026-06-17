package subsecuencia

import "testing"

func TestLCS(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"", "", 0},
		{"abc", "", 0},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"AGGTAB", "GXTXAYB", 4}, // "GTAB"
		{"ABCDGH", "AEDFHR", 3},  // "ADH"
		{"abcdef", "acf", 3},     // "acf"
	}
	for _, tt := range tests {
		got := LCS(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("LCS(%q, %q) = %d, esperado %d", tt.a, tt.b, got, tt.want)
		}
	}
}
