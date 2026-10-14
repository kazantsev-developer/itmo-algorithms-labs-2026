// Package algo содержит реализации алгоритмов динамического программирования
package algo

type BSTResult struct {
	Cost int
	N    int
	Root [][]int
}

func SolveOptimalBST(p []int) BSTResult {
	n := len(p)

	pref := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pref[i] = pref[i-1] + p[i-1]
	}
	sum := func(i, j int) int {
		if i > j {
			return 0
		}
		return pref[j] - pref[i-1]
	}

	cost := make([][]int, n+2)
	root := make([][]int, n+2)
	for i := range cost {
		cost[i] = make([]int, n+2)
		root[i] = make([]int, n+2)
	}

	for length := 1; length <= n; length++ {
		for i := 1; i+length-1 <= n; i++ {
			j := i + length - 1
			best := 1 << 30
			bestK := i
			for k := i; k <= j; k++ {
				v := cost[i][k-1] + cost[k+1][j]
				if v < best {
					best = v
					bestK = k
				}
			}
			cost[i][j] = best + sum(i, j)
			root[i][j] = bestK
		}
	}

	return BSTResult{
		Cost: cost[1][n],
		N:    n,
		Root: root,
	}
}
