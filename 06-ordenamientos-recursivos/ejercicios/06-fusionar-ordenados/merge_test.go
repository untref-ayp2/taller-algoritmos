package merge

import (
	"slices"
	"testing"
)

func TestMergeSortedBothEmpty(t *testing.T) {
	got := MergeSorted([]int{}, []int{}, func(a, b int) bool { return a < b })
	if len(got) != 0 {
		t.Errorf("esperado vacío, got %v", got)
	}
}

func TestMergeSortedOneEmpty(t *testing.T) {
	a := []int{1, 2, 3}
	got := MergeSorted(a, []int{}, func(a, b int) bool { return a < b })
	if !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("esperado [1 2 3], got %v", got)
	}
}

func TestMergeSortedOtherEmpty(t *testing.T) {
	b := []int{4, 5, 6}
	got := MergeSorted([]int{}, b, func(a, b int) bool { return a < b })
	if !slices.Equal(got, []int{4, 5, 6}) {
		t.Errorf("esperado [4 5 6], got %v", got)
	}
}

func TestMergeSortedInterleaved(t *testing.T) {
	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8}
	got := MergeSorted(a, b, func(a, b int) bool { return a < b })
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !slices.Equal(got, want) {
		t.Errorf("esperado %v, got %v", want, got)
	}
}

func TestMergeSortedDifferentLengths(t *testing.T) {
	a := []int{2, 5, 9}
	b := []int{3}
	got := MergeSorted(a, b, func(a, b int) bool { return a < b })
	want := []int{2, 3, 5, 9}
	if !slices.Equal(got, want) {
		t.Errorf("esperado %v, got %v", want, got)
	}
}

func TestMergeSortedAllFromFirst(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	got := MergeSorted(a, b, func(a, b int) bool { return a < b })
	want := []int{1, 2, 3, 4, 5, 6}
	if !slices.Equal(got, want) {
		t.Errorf("esperado %v, got %v", want, got)
	}
}

func TestMergeSortedAllFromSecond(t *testing.T) {
	a := []int{4, 5, 6}
	b := []int{1, 2, 3}
	got := MergeSorted(a, b, func(a, b int) bool { return a < b })
	want := []int{1, 2, 3, 4, 5, 6}
	if !slices.Equal(got, want) {
		t.Errorf("esperado %v, got %v", want, got)
	}
}

func TestMergeSortedDuplicates(t *testing.T) {
	a := []int{1, 2, 4, 4}
	b := []int{2, 4, 5}
	got := MergeSorted(a, b, func(a, b int) bool { return a < b })
	want := []int{1, 2, 2, 4, 4, 4, 5}
	if !slices.Equal(got, want) {
		t.Errorf("esperado %v, got %v", want, got)
	}
}

func TestMergeSortedDesc(t *testing.T) {
	a := []int{7, 5, 3}
	b := []int{6, 4, 2}
	got := MergeSorted(a, b, func(a, b int) bool { return a > b })
	want := []int{7, 6, 5, 4, 3, 2}
	if !slices.Equal(got, want) {
		t.Errorf("esperado %v, got %v", want, got)
	}
}

func TestMergeSortedStrings(t *testing.T) {
	a := []string{"abeja", "casa"}
	b := []string{"bosque", "zorro"}
	got := MergeSorted(a, b, func(a, b string) bool { return a < b })
	want := []string{"abeja", "bosque", "casa", "zorro"}
	if !slices.Equal(got, want) {
		t.Errorf("esperado %v, got %v", want, got)
	}
}

func TestMergeSortedStable(t *testing.T) {
	a := []int{1, 3}
	b := []int{1, 4}
	got := MergeSorted(a, b, func(a, b int) bool { return a < b })
	// con <, elementos iguales preservan orden: a[0]=1 antes que b[0]=1
	if got[0] != 1 || got[1] != 1 || got[2] != 3 || got[3] != 4 {
		t.Errorf("no estable: %v", got)
	}
}
