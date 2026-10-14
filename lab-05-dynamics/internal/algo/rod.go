// Package algo содержит реализации алгоритмов динамического программирования
package algo

type Order struct {
	Length int
	Price  int
}

type RodResult struct {
	MaxProfit int
	Chosen    []Order
	TotalLen  int
}

func SolveRodCutting(orders []Order, L, C int) RodResult {
	type candidate struct {
		order Order
		gain  int
	}

	var cands []candidate
	for _, o := range orders {
		g := o.Price - C
		if g > 0 && o.Length <= L {
			cands = append(cands, candidate{o, g})
		}
	}

	n := len(cands)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, L+1)
	}

	for i := 1; i <= n; i++ {
		d := cands[i-1].order.Length
		g := cands[i-1].gain
		for l := 0; l <= L; l++ {
			dp[i][l] = dp[i-1][l]
			if l >= d && dp[i-1][l-d]+g > dp[i][l] {
				dp[i][l] = dp[i-1][l-d] + g
			}
		}
	}

	var chosen []Order
	totalLen := 0
	l := L
	for i := n; i >= 1; i-- {
		if dp[i][l] != dp[i-1][l] {
			o := cands[i-1].order
			chosen = append(chosen, o)
			totalLen += o.Length
			l -= o.Length
		}
	}
	for i, j := 0, len(chosen)-1; i < j; i, j = i+1, j-1 {
		chosen[i], chosen[j] = chosen[j], chosen[i]
	}

	return RodResult{
		MaxProfit: dp[n][L],
		Chosen:    chosen,
		TotalLen:  totalLen,
	}
}
