package quickselect

import "testing"

func TestQuickSelectMin(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6}
	got := QuickSelect(s, 0, func(a, b int) bool { return a < b })
	if got != 1 {
		t.Errorf("min esperado 1, got %d", got)
	}
}

func TestQuickSelectMax(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6}
	got := QuickSelect(s, len(s)-1, func(a, b int) bool { return a < b })
	if got != 9 {
		t.Errorf("max esperado 9, got %d", got)
	}
}

func TestQuickSelectMedian(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6}
	// sorted: [1, 1, 2, 3, 4, 5, 6, 9], median pos 3 = 3, pos 4 = 4
	got := QuickSelect(s, 3, func(a, b int) bool { return a < b })
	if got != 3 {
		t.Errorf("k=3 esperado 3, got %d", got)
	}
	got = QuickSelect(s, 4, func(a, b int) bool { return a < b })
	if got != 4 {
		t.Errorf("k=4 esperado 4, got %d", got)
	}
}

func TestQuickSelectDesc(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6}
	got := QuickSelect(s, 0, func(a, b int) bool { return a > b })
	if got != 9 {
		t.Errorf("mayor con less invertido esperado 9, got %d", got)
	}
}

func TestQuickSelectSingle(t *testing.T) {
	s := []int{42}
	got := QuickSelect(s, 0, func(a, b int) bool { return a < b })
	if got != 42 {
		t.Errorf("unico elemento esperado 42, got %d", got)
	}
}

func TestQuickSelectDuplicates(t *testing.T) {
	s := []int{5, 5, 5, 5}
	for k := 0; k < 4; k++ {
		got := QuickSelect(s, k, func(a, b int) bool { return a < b })
		if got != 5 {
			t.Errorf("k=%d esperado 5, got %d", k, got)
		}
	}
}

func TestQuickSelectStrings(t *testing.T) {
	s := []string{"zorro", "abeja", "casa"}
	got := QuickSelect(s, 1, func(a, b string) bool { return a < b })
	// sorted: [abeja, casa, zorro]
	if got != "casa" {
		t.Errorf("mediana esperada 'casa', got %s", got)
	}
}

func TestQuickSelectPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("se esperaba panic por k fuera de rango")
		}
	}()
	QuickSelect([]int{1, 2, 3}, 5, func(a, b int) bool { return a < b })
}

func TestQuickSelectLarge(t *testing.T) {
	s := make([]int, 1000)
	for i := range s {
		s[i] = 1000 - i
	}
	got := QuickSelect(s, 499, func(a, b int) bool { return a < b })
	if got != 500 {
		t.Errorf("elemento en pos 499 esperado 500, got %d", got)
	}
}
