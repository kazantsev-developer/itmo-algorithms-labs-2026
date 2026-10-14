// Package algo содержит реализации алгоритмов динамического программирования
package algo

type Item struct {
	Weight int
	Value  int
	Time   int
}

type Result struct {
	MaxValue    int
	ChosenItems []Item
}

func SolveKnapsack(items []Item, maxW int, maxT int) Result {
	dp := make([][]int, maxW+1)
	for i := range dp {
		dp[i] = make([]int, maxT+1)
	}

	for _, item := range items {
		for w := maxW; w >= item.Weight; w-- {
			for t := maxT; t >= item.Time; t-- {
				if dp[w-item.Weight][t-item.Time]+item.Value > dp[w][t] {
					dp[w][t] = dp[w-item.Weight][t-item.Time] + item.Value
				}
			}
		}
	}

	maxValue := dp[maxW][maxT]

	return Result{
		MaxValue:    maxValue,
		ChosenItems: nil,
	}
}
