package quicksort

import (
	"cmp"
	"slices"
	"testing"
)

func isSorted[T any](s []T, cmp func(a, b T) int) bool {
	for i := 1; i < len(s); i++ {
		if cmp(s[i], s[i-1]) < 0 {
			return false
		}
	}
	return true
}

func TestQuickSortEmpty(t *testing.T) {
	s := []int{}
	QuickSort(s, cmp.Compare)
}

func TestQuickSortSingle(t *testing.T) {
	s := []int{42}
	QuickSort(s, cmp.Compare)
	if s[0] != 42 {
		t.Errorf("expected 42, got %d", s[0])
	}
}

func TestQuickSortAsc(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	QuickSort(s, cmp.Compare)
	if !isSorted(s, cmp.Compare) {
		t.Errorf("not sorted ascending: %v", s)
	}
}

func TestQuickSortDesc(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	desc := func(a, b int) int { return cmp.Compare(b, a) }
	QuickSort(s, desc)
	if !isSorted(s, desc) {
		t.Errorf("not sorted descending: %v", s)
	}
}

func TestQuickSortAlreadySorted(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	QuickSort(s, cmp.Compare)
	if !isSorted(s, cmp.Compare) {
		t.Errorf("already sorted slice should remain sorted")
	}
}

func TestQuickSortReverseSorted(t *testing.T) {
	s := []int{5, 4, 3, 2, 1}
	QuickSort(s, cmp.Compare)
	if !isSorted(s, cmp.Compare) {
		t.Errorf("reverse sorted should become sorted: %v", s)
	}
}

func TestQuickSortDuplicates(t *testing.T) {
	s := []int{7, 7, 7, 7, 7}
	QuickSort(s, cmp.Compare)
	if !isSorted(s, cmp.Compare) {
		t.Errorf("duplicates: %v", s)
	}
}

func TestQuickSortStrings(t *testing.T) {
	s := []string{"zorro", "abeja", "casa", "bosque"}
	QuickSort(s, cmp.Compare)
	if !isSorted(s, cmp.Compare) {
		t.Errorf("strings not sorted: %v", s)
	}
}

func TestQuickSortLarge(t *testing.T) {
	s := make([]int, 1000)
	for i := range s {
		s[i] = 1000 - i
	}
	QuickSort(s, cmp.Compare)
	if !slices.IsSorted(s) {
		t.Errorf("large slice not sorted")
	}
}
