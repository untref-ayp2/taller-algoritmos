package fibonacci

// FibonacciTab calcula el n-ésimo número de Fibonacci usando tabulación (bottom-up).
// Tiempo O(n), espacio O(n).
func FibonacciTab(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// FibonacciMemo calcula el n-ésimo número de Fibonacci usando memoización (top-down).
// Tiempo O(n), espacio O(n) por la pila de recursión y el mapa de caché.
func FibonacciMemo(n int) int {
	memo := make(map[int]int)
	return fibMemo(n, memo)
}

func fibMemo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
	return memo[n]
}
