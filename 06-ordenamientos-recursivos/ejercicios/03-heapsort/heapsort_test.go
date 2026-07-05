package heapsort

import (
	"slices"
	"testing"
)

func isSorted[T any](s []T, less func(a, b T) bool) bool {
	for i := 1; i < len(s); i++ {
		if less(s[i], s[i-1]) {
			return false
		}
	}
	return true
}

func TestHeapSortEmpty(t *testing.T) {
	s := []int{}
	HeapSort(s, func(a, b int) bool { return a < b })
}

func TestHeapSortSingle(t *testing.T) {
	s := []int{42}
	HeapSort(s, func(a, b int) bool { return a < b })
	if s[0] != 42 {
		t.Errorf("expected 42, got %d", s[0])
	}
}

func TestHeapSortAsc(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	HeapSort(s, func(a, b int) bool { return a < b })
	if !isSorted(s, func(a, b int) bool { return a < b }) {
		t.Errorf("not sorted ascending: %v", s)
	}
}

func TestHeapSortDesc(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	HeapSort(s, func(a, b int) bool { return a > b })
	if !isSorted(s, func(a, b int) bool { return a > b }) {
		t.Errorf("not sorted descending: %v", s)
	}
}

func TestHeapSortAlreadySorted(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	HeapSort(s, func(a, b int) bool { return a < b })
	if !isSorted(s, func(a, b int) bool { return a < b }) {
		t.Errorf("already sorted should remain sorted")
	}
}

func TestHeapSortReverseSorted(t *testing.T) {
	s := []int{5, 4, 3, 2, 1}
	HeapSort(s, func(a, b int) bool { return a < b })
	if !isSorted(s, func(a, b int) bool { return a < b }) {
		t.Errorf("reverse sorted should become sorted: %v", s)
	}
}

func TestHeapSortDuplicates(t *testing.T) {
	s := []int{7, 7, 7, 7, 7}
	HeapSort(s, func(a, b int) bool { return a < b })
	if !isSorted(s, func(a, b int) bool { return a < b }) {
		t.Errorf("duplicates: %v", s)
	}
}

func TestHeapSortStrings(t *testing.T) {
	s := []string{"zorro", "abeja", "casa", "bosque"}
	HeapSort(s, func(a, b string) bool { return a < b })
	if !isSorted(s, func(a, b string) bool { return a < b }) {
		t.Errorf("strings not sorted: %v", s)
	}
}

func TestHeapSortLarge(t *testing.T) {
	s := make([]int, 1000)
	for i := range s {
		s[i] = 1000 - i
	}
	HeapSort(s, func(a, b int) bool { return a < b })
	if !slices.IsSorted(s) {
		t.Errorf("large slice not sorted")
	}
}
