package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacciTab(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3},
		{5, 5}, {6, 8}, {7, 13}, {10, 55},
	}
	for _, tt := range tests {
		got := FibonacciTab(tt.n)
		if got != tt.want {
			t.Errorf("FibonacciTab(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestFibonacciMemo(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3},
		{5, 5}, {10, 55}, {20, 6765},
	}
	for _, tt := range tests {
		got := FibonacciMemo(tt.n)
		if got != tt.want {
			t.Errorf("FibonacciMemo(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func Example() {
	fmt.Println("Tab(10):", FibonacciTab(10))
	fmt.Println("Memo(10):", FibonacciMemo(10))
	// Output:
	// Tab(10): 55
	// Memo(10): 55
}
