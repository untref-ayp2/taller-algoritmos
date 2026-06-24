package mochila

// Item representa un objeto con peso y valor.
type Item struct {
	Peso  int
	Valor int
}

// MochilaTab resuelve el problema de la mochila 0/1 usando tabulación (bottom-up).
// Retorna el valor máximo alcanzable sin superar la capacidad.
// Tiempo O(n*W), espacio O(n*W) donde n = len(items), W = capacidad.
func MochilaTab(items []Item, capacidad int) int {
	n := len(items)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacidad+1)
	}
	for i := 1; i <= n; i++ {
		for w := 1; w <= capacidad; w++ {
			if items[i-1].Peso > w {
				dp[i][w] = dp[i-1][w]
			} else {
				sin := dp[i-1][w]
				con := items[i-1].Valor + dp[i-1][w-items[i-1].Peso]
				if con > sin {
					dp[i][w] = con
				} else {
					dp[i][w] = sin
				}
			}
		}
	}
	return dp[n][capacidad]
}

// MochilaMemo resuelve el problema de la mochila 0/1 usando memoización (top-down).
// La matriz memo se inicializa con -1 para indicar «no calculado».
// Tiempo O(n*W), espacio O(n*W).
func MochilaMemo(items []Item, capacidad int) int {
	n := len(items)
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, capacidad+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return mochilaMemo(items, n, capacidad, memo)
}

func mochilaMemo(items []Item, i int, w int, memo [][]int) int {
	if i == 0 || w == 0 {
		return 0
	}
	if memo[i][w] != -1 {
		return memo[i][w]
	}
	if items[i-1].Peso > w {
		memo[i][w] = mochilaMemo(items, i-1, w, memo)
	} else {
		sin := mochilaMemo(items, i-1, w, memo)
		con := items[i-1].Valor + mochilaMemo(items, i-1, w-items[i-1].Peso, memo)
		if con > sin {
			memo[i][w] = con
		} else {
			memo[i][w] = sin
		}
	}
	return memo[i][w]
}

// MochilaTabConItems resuelve el problema y además devuelve los items seleccionados.
func MochilaTabConItems(items []Item, capacidad int) (int, []Item) {
	n := len(items)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacidad+1)
	}
	for i := 1; i <= n; i++ {
		for w := 1; w <= capacidad; w++ {
			if items[i-1].Peso > w {
				dp[i][w] = dp[i-1][w]
			} else {
				sin := dp[i-1][w]
				con := items[i-1].Valor + dp[i-1][w-items[i-1].Peso]
				if con > sin {
					dp[i][w] = con
				} else {
					dp[i][w] = sin
				}
			}
		}
	}
	// Reconstruir la selección
	w := capacidad
	seleccion := make([]Item, 0)
	for i := n; i > 0 && w > 0; i-- {
		if dp[i][w] != dp[i-1][w] {
			seleccion = append(seleccion, items[i-1])
			w -= items[i-1].Peso
		}
	}
	return dp[n][capacidad], seleccion
}
